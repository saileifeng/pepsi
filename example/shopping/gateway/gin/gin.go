package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saileifeng/pepsi/example/shopping/gateway/gin/controllers"
	"github.com/saileifeng/pepsi/example/shopping/name"
	"github.com/saileifeng/pepsi/example/shopping/utils"
	"github.com/saileifeng/pepsi/registry/consul"
	"log"
	"net/http"
)

var consulAddr  = "127.0.0.1:8500"
var port  = 8080

func main() {
	flag.StringVar(&consulAddr, "registry_address", "127.0.0.1:8500", "registry address")
	flag.IntVar(&port,"server_port",8080,"server port")
	flag.Parse()

	engine := gin.Default()
	gin.SetMode("debug")
	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	//cc := consul.NewClietnConn("127.0.0.1:8500","buy")
	buy := &controllers.BuyGoodsControllers{CC: consul.NewClietnConn(consulAddr, name.APIBuy)}

	//处理购物
	engine.POST("/shopping/v1/buyGoods", buy.BuyGoods)

	go func() {
		if err := engine.Run(fmt.Sprintf("%s:%d", "0.0.0.0", port)); err != nil {
			panic(err)
		}
	}()

	utils.ShutDownHook(func() {
		log.Println("server stop")
	})
}
