package main

import (
	"context"
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/saileifeng/pepsi/example/shopping/name"
	"github.com/saileifeng/pepsi/example/shopping/pb"
	"github.com/saileifeng/pepsi/registry/consul"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"log"
	"time"
)

var consulAddr = "127.0.0.1:8500"
var port = 0
var serviceName = name.APIBuy

//BuyGoodsService 购物服务
type BuyGoodsService struct {
	//商品服务
	GoodsCC *grpc.ClientConn
	//订单服务
	OderCC *grpc.ClientConn
	Limit  *rate.Limiter
}

//BuyGoods 购买商品生成订单
func (bgs *BuyGoodsService) BuyGoods(ctx context.Context, req *pb.BuyGoodsRequest) (*pb.BuyGoodsResponse, error) {
	//limitCtx,cancel := context.WithCancel(context.Background())
	//log.Println("BuyGoods")
	//defer cancel()

	if allow := bgs.Limit.Allow(); !allow {
		//log.Println("limit cancel")
		return nil, errors.New("limit cancel")
	}
	resp := &pb.BuyGoodsResponse{}
	err := hystrix.Do(serviceName, func() error {
		//从货架上拿到物品
		goodsService := pb.NewGoodsServiceClient(bgs.GoodsCC)
		context1, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		goodsResp, err := goodsService.GetGoods(context1, &pb.GetGoodsRequest{GoodsInfos: req.GoodsInfos})
		if err != nil {
			return err
		}

		//生成交易订单--超时未支付的订单需要将商品放回货架上
		orderService := pb.NewOrderServiceClient(bgs.OderCC)
		context2, cancel1 := context.WithTimeout(context.Background(), time.Second)
		defer cancel1()
		orderResp, err := orderService.CreateOrder(context2, &pb.CreateOrderInfoRequest{GoodsInfos: goodsResp.GoodsInfos, TotalPrice: goodsResp.TotalPrice, UserID: req.UserID})
		if err != nil {
			return err
		}
		resp = &pb.BuyGoodsResponse{OrderID: orderResp.OrderID, UserID: req.UserID, GoodsInfos: goodsResp.GoodsInfos, TotalPrice: goodsResp.TotalPrice}
		log.Println("BuyGoods :", resp)
		return err
	}, func(e error) error {
		log.Println("BuyGoods err :", e)
		//服务降级处理
		switch e {
		case hystrix.ErrMaxConcurrency:
		case hystrix.ErrCircuitOpen:
		case hystrix.ErrTimeout:
		}
		return e
	})
	//log.Println(err)
	return resp, err

}

func main() {
	r := consul.NewRegister(consulAddr, serviceName, port)

	//创建限流器 初始容量为10，每秒产生一个令牌
	limit := rate.NewLimiter(rate.Every(time.Second), 15)

	//创建熔断器
	hystrix.ConfigureCommand(serviceName, hystrix.CommandConfig{
		Timeout: 2000, //超时时间设置  单位毫秒

		MaxConcurrentRequests: 8, //最大请求数

		SleepWindow: 1, //过多长时间，熔断器再次检测是否开启。单位毫秒

		ErrorPercentThreshold: 30, //错误率

		RequestVolumeThreshold: 5, //请求阈值  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
	})

	goodsCC := consul.NewClietnConn(consulAddr, name.SrvGoods)
	oderCC := consul.NewClietnConn(consulAddr, name.SrvOrder)
	//注册购物服务
	pb.RegisterBuyServiceServer(r.Server, &BuyGoodsService{GoodsCC: goodsCC, OderCC: oderCC, Limit: limit})

	//支付订单
	//支付扣款--账户余额扣款
	//确定交易订单完成
	r.Run()
}
