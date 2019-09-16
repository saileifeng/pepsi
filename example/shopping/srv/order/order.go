package main

import (
	"context"
	"github.com/saileifeng/pepsi/example/shopping/name"
	"github.com/saileifeng/pepsi/example/shopping/pb"
	"github.com/saileifeng/pepsi/registry/consul"
	"log"
	"time"
)


var consulAddr = "127.0.0.1:8500"
var port = 0
var serviceName = name.SRV_ORDER

type OrderService struct {

}

func (os *OrderService)CreateOrder(ctx context.Context, req *pb.CreateOrderInfoRequest) (*pb.CreateOrderInfoResponse, error)  {
	log.Println("CreateOrder :",req)
	return &pb.CreateOrderInfoResponse{OrderID:time.Now().Unix()},nil
}

func main() {
	r := consul.NewRegister(consulAddr,serviceName,port)
	pb.RegisterOrderServiceServer(r.Server,&OrderService{})
	r.Run()
}
