package consul

import (
	"fmt"
	"github.com/saileifeng/pepsi/registry/consul/register"
	"github.com/saileifeng/pepsi/registry/consul/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)


//NewClietnConn 创建客户端
func NewClietnConn(consulAddr,serviceName string) *grpc.ClientConn {
	schema, err := resolver.StartConsulResolver(consulAddr, serviceName)
	log.Println("NewClietnConn schema :",schema)
	//consul集群在未完成选举的时候需要再次重试创建,这种问题似乎出现在低版本的consul
	for i := 0;i<10;i++ {
		if err!=nil {
			time.Sleep(time.Second)
			log.Println("retry NewClietnConn")
			schema, err = resolver.StartConsulResolver(consulAddr, serviceName)
		}else {
			break
		}
	}
	if err != nil {
		log.Fatal("init consul resovler err", err.Error())
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", schema,serviceName), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

//Registry 服务注册自定义结构体
type Registry struct {
	consulAddr,service string
	port int
	listener net.Listener
	Server *grpc.Server
	register *register.ConsulRegister
}
//NewRegister 创建新的服务注册
func NewRegister(consulAddr,service string,port int) *Registry {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v",port))
	if err != nil {
		log.Fatalln(err)
	}
	addrs := strings.Split(listener.Addr().String(),":")
	port,err = strconv.Atoi(addrs[len(addrs)-1])
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("start server port :",addrs[len(addrs)-1])
	//consul service register
	nr := register.NewConsulRegister(consulAddr,service,port)
	nr.Register()
	//start grpc server
	serv :=  grpc.NewServer()
	//registe health check
	grpc_health_v1.RegisterHealthServer(serv, &register.HealthImpl{})

	return &Registry{consulAddr:consulAddr,service:service,port:port,listener:listener,Server:serv,register:nr}
}
//Run 启动
func (r *Registry)Run()  {
	//server hook
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit
		log.Println("do run hook")
		r.register.Deregister()
		r.Server.Stop()
	}()

	if err := r.Server.Serve(r.listener); err != nil {
		panic(err)
	}
}
