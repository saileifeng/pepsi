package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saileifeng/pepsi/example/shopping/gateway/gin/controllers"
	"github.com/saileifeng/pepsi/example/shopping/name"
	"github.com/saileifeng/pepsi/example/shopping/utils"
	"github.com/saileifeng/pepsi/registry/consul"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	gin.SetMode("debug")
	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	//cc := consul.NewClietnConn("127.0.0.1:8500","buy")
	buy := &controllers.BuyGoodsControllers{CC: consul.NewClietnConn("127.0.0.1:8500", name.APIBuy)}

	//处理购物
	engine.POST("/shopping/v1/buyGoods", buy.BuyGoods)

	go func() {
		if err := engine.Run(fmt.Sprintf("%s:%s", "0.0.0.0", "8080")); err != nil {
			panic(err)
		}
	}()

	utils.ShutDownHook(func() {
		log.Println("server stop")
	})
}
