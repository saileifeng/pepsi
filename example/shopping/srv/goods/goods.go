package main

import (
	"context"
	"github.com/saileifeng/pepsi/example/shopping/name"
	"github.com/saileifeng/pepsi/example/shopping/pb"
	"github.com/saileifeng/pepsi/registry/consul"
	"log"
)

var consulAddr = "127.0.0.1:8500"
var port = 0
var serviceName = name.SrvGoods

//GoodsService ...
type GoodsService struct {
}

//GetGoods 处理购买物品请求
func (gs *GoodsService) GetGoods(ctx context.Context, req *pb.GetGoodsRequest) (*pb.GetGoodsResponse, error) {
	log.Println("GetGoods :", req)
	total := int64(0)
	for _, v := range req.GoodsInfos {
		v.UnitPrice = v.GoodsID * 10
		total += v.UnitPrice * v.Count
	}
	return &pb.GetGoodsResponse{GoodsInfos: req.GoodsInfos, TotalPrice: total}, nil
}

func main() {
	r := consul.NewRegister(consulAddr, serviceName, port)
	//consul.NewClietnConn(consulAddr,serviceName)
	pb.RegisterGoodsServiceServer(r.Server, &GoodsService{})
	r.Run()
}
