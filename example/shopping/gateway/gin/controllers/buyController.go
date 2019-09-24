package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/saileifeng/pepsi/example/shopping/gateway/gin/view_models"
	"github.com/saileifeng/pepsi/example/shopping/pb"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net/http"
)
//BuyGoodsControllers ...
type BuyGoodsControllers struct {
	CC *grpc.ClientConn
	Limit * rate.Limiter
}
//BuyGoods 购买商品服务
func (bgc *BuyGoodsControllers)BuyGoods(ctx *gin.Context)  {
	//jaeger
	buygoodsSpan,buygoodsCtx:= opentracing.StartSpanFromContext(context.Background(),"BuyGoods")
	defer buygoodsSpan.Finish()
	//buygoodsSpan.SetTag("ctx",ctx)
	if bgc.Limit.Allow() {
		buygoodsSpan.SetTag("bgc.Limit.Allow",true)
		//jaeger
		decodeJsonSpan,_:= opentracing.StartSpanFromContext(buygoodsCtx,"decodeJson")
		defer decodeJsonSpan.Finish()
		obj := &pb.BuyGoodsRequest{}
		//解析为json结构体
		//TODO json结构体需要重新定义，不能用proto生成的,字段都有为空选项，数据完成性会出问题
		err := ctx.BindJSON(obj)
		if err != nil{
			ctx.JSON(http.StatusBadRequest,&viewmodels.ResultInfo{Code:1,Data:&err})
			return
		}
		decodeJsonSpan.SetTag("req_json",obj)
		//jaeger
		serviceSpan,serviceCtx:=opentracing.StartSpanFromContext(buygoodsCtx,"request buygoods service")
		defer serviceSpan.Finish()
		bc := pb.NewBuyServiceClient(bgc.CC)

		//context, cancel := context.WithTimeout(context.Background(), time.Second)
		//defer cancel()
		//请求购买服务生成订单信息
		resp,err := bc.BuyGoods(serviceCtx,obj)
		if err != nil {
			ctx.JSON(http.StatusBadRequest,&viewmodels.ResultInfo{Code:1,Data:&err})
			return
		}
		ctx.JSON(http.StatusOK,&viewmodels.ResultInfo{Code:0,Data:resp})
	}else {
		buygoodsSpan.SetTag("bgc.Limit.Allow",false)
		buygoodsSpan.SetTag("http.StatusBadRequest",http.StatusBadRequest)
		buygoodsSpan.SetTag("viewmodels.ResultInfo","request too fast")
		ctx.JSON(http.StatusBadRequest,&viewmodels.ResultInfo{Code:1,Data:"request too fast"})
	}

}
