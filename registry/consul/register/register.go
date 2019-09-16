package register

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"time"
)

var (
	//ErrNoClientAgent no client agent
	ErrNoClientAgent = errors.New("no client agent")
)

// HealthImpl 健康检查实现
type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	//log.Println("check health")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch 健康检测监听
func (h *HealthImpl) Watch(hcr *grpc_health_v1.HealthCheckRequest, ws grpc_health_v1.Health_WatchServer) error{
	//log.Println("check Watch")
	return nil
}


// NewConsulRegister create a new consul register
func NewConsulRegister(addr , service string,port int) *ConsulRegister {
	return &ConsulRegister{
		Address: addr,
		Service: service,
		ServiceID:fmt.Sprintf("%v-%v-%v", service, localIP(), port),
		Tag:     []string{},
		Port:    port,
		DeregisterCriticalServiceAfter: time.Duration(10) * time.Second,
		Interval:                       time.Duration(3) * time.Second,
	}
}

// ConsulRegister consul service register
type ConsulRegister struct {
	Address                        string
	Service                        string
	ServiceID						string
	Tag                            []string
	Port                           int
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
	ClientAgent						*api.Client
}
// Deregister 注销服务
func (r *ConsulRegister) Deregister() error {
	log.Println("ConsulRegister Deregister")
	if r.ClientAgent == nil {
		return ErrNoClientAgent
	}
	if err := r.ClientAgent.Agent().ServiceDeregister(r.ServiceID); err != nil {
		log.Fatal(err)
	}
	log.Println("ConsulRegister Deregister nil")
	return nil
}

// Register register service
func (r *ConsulRegister) Register() error {
	config := api.DefaultConfig()
	config.Address = r.Address
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}
	r.ClientAgent = client
	agent := client.Agent()

	IP := localIP()
	reg := &api.AgentServiceRegistration{
		ID:      r.ServiceID, // 服务节点的名称
		Name:    r.Service,    // 服务名称
		Tags:    r.Tag,                                          // tag，可以为空
		Port:    r.Port,                                         // 服务端口
		Address: IP,                                             // 服务 IP
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval: r.Interval.String(),                            // 健康检查间隔
			GRPC:     fmt.Sprintf("%v:%v/%v", IP, r.Port, r.Service), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(), // 注销时间，相当于过期时间
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}
	return nil
}

//func (r *ConsulRegister)ShutDownHook()  {
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
//	<-quit
//	r.Deregister()
//}


func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {

			if ipnet.IP.To4() != nil {
				//log.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
