# pepsi

gin+grpc+consul实现的微服务demo:购物模块

### 业务流程
```mermaid
    graph LR
		client[用户] -- 提交订单 --> gateway[网关服务]
		gateway --> client
		gateway -- 提交请求 --> buy[购物api服务]
		buy -- 返回订单信息 --> gateway
		buy -- 查询物品库存并减少 --> goods[商品服务]
		goods --> buy
		buy -- 生成订单返回订单信息 --> order[订单服务]
		order --> buy
```
### 启动方式

consul:

    consul agent -dev

gateway:

    go run example/gateway/gin/gin.go

api:
   
    go run example/api/buy/buy.go 

srv:

    go run example/srv/goods/goods.go 
    go run example/srv/order/order.go 
    
test:

    curl http://127.0.0.1:8080/shopping/v1/buyGoods -X POST -H "Content-Type:application/json" -d '{"userID":89757,"goodsInfos":[{"goodsID":1,"count":2},{"goodsID":2,"count":3}]}'
 