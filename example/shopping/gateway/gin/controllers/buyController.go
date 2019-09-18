package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/saileifeng/pepsi/example/shopping/gateway/gin/view_models"
	"github.com/saileifeng/pepsi/example/shopping/pb"
	"google.golang.org/grpc"
	"net/http"
)
//BuyGoodsControllers ...
type BuyGoodsControllers struct {
	CC *grpc.ClientConn
}
//BuyGoods 购买商品服务
func (bgc *BuyGoodsControllers)BuyGoods(ctx *gin.Context)  {
	obj := &pb.BuyGoodsRequest{}
	//解析为json结构体
	//TODO json结构体需要重新定义，不能用proto生成的,字段都有为空选项，数据完成性会出问题
	ctx.BindJSON(obj)
	bc := pb.NewBuyServiceClient(bgc.CC)

	//context, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//请求购买服务生成订单信息
	resp,err := bc.BuyGoods(context.Background(),obj)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,&viewmodels.ResultInfo{Code:1,Data:&err})
		return
	}
	ctx.JSON(http.StatusOK,&viewmodels.ResultInfo{Code:0,Data:resp})
}
