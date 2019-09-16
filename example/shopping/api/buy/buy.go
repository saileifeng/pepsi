package main

import (
	"context"
	"github.com/saileifeng/pepsi/example/shopping/name"
	"github.com/saileifeng/pepsi/example/shopping/pb"
	"github.com/saileifeng/pepsi/registry/consul"
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

}
//BuyGoods 购买商品生成订单
func (bgs *BuyGoodsService)BuyGoods(ctx context.Context, req *pb.BuyGoodsRequest) (*pb.BuyGoodsResponse, error)  {
	log.Println("BuyGoods :",req)
	//从货架上拿到物品
	goodsService := pb.NewGoodsServiceClient(bgs.GoodsCC)
	context1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	goodsResp,err :=goodsService.GetGoods(context1,&pb.GetGoodsRequest{GoodsInfos:req.GoodsInfos})
	if err != nil{
		return nil,err
	}

	//生成交易订单--超时未支付的订单需要将商品放回货架上
	orderService := pb.NewOrderServiceClient(bgs.OderCC)
	context2, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	orderResp,err :=orderService.CreateOrder(context2,&pb.CreateOrderInfoRequest{GoodsInfos:goodsResp.GoodsInfos,TotalPrice:goodsResp.TotalPrice,UserID:req.UserID})
	if err != nil{
		return nil,err
	}
	return &pb.BuyGoodsResponse{OrderID:orderResp.OrderID,UserID:req.UserID,GoodsInfos:goodsResp.GoodsInfos,TotalPrice:goodsResp.TotalPrice},nil
}


func main() {
	r := consul.NewRegister(consulAddr,serviceName,port)
	goodsCC := consul.NewClietnConn(consulAddr,name.SRV_GOODS)
	oderCC := consul.NewClietnConn(consulAddr,name.SRV_ORDER)
	//注册购物服务
	pb.RegisterBuyServiceServer(r.Server,&BuyGoodsService{GoodsCC:goodsCC,OderCC:oderCC})

	//支付订单
	//支付扣款--账户余额扣款
	//确定交易订单完成
	r.Run()
}
