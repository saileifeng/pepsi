package resolver


import (
	"context"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"log"
	"sync"
	"time"
)

type consulBuilder struct {
	address     string
	client      *consulapi.Client
	serviceName string
	waitIndex uint64
}
// NewConsulBuilder 创建consulBuilder
func NewConsulBuilder(address string) resolver.Builder {
	config := consulapi.DefaultConfig()
	config.Address = address
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("create consul client error", err.Error())
		return nil
	}
	return &consulBuilder{address: address, client: client}
}

func (cb *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	cb.serviceName = target.Endpoint
	//first need get all
	adds, _, err := cb.resolve()
	if err != nil {
		return nil, err
	}

	cc.UpdateState(resolver.State{Addresses:adds})

	consulResolver := newConsulResolver(&cc, cb, opts)
	consulResolver.wg.Add(1)
	go consulResolver.watcher()

	return consulResolver, nil
}

//func (cb consulBuilder) syncResolve() ([]resolver.Address, string, error) {
//	serviceEntries, _, err := cb.client.Health().Service(cb.serviceName, "", true, &consulapi.QueryOptions{})
//	if err != nil {
//		return nil, "", err
//	}
//	//cb.waitIndex = queryMeta.LastIndex
//	//log.Println("queryMeta syncResolve",queryMeta.LastIndex)
//	adds := make([]resolver.Address, 0)
//	for _, serviceEntry := range serviceEntries {
//		address := resolver.Address{Addr: fmt.Sprintf("%s:%d", serviceEntry.Service.Address, serviceEntry.Service.Port)}
//		adds = append(adds, address)
//	}
//	return adds, "", nil
//}


func (cb consulBuilder) resolve() ([]resolver.Address, string, error) {
	//WaitIndex 会阻塞等待服务状态有变化之后才获取服务状态
	//但是我在现在的版本中不会阻塞，可能是我代码写的有问题
	serviceEntries, queryMeta, err := cb.client.Health().Service(cb.serviceName, "", true, &consulapi.QueryOptions{WaitIndex:cb.waitIndex})
	if err != nil {
		return nil, "", err
	}
	cb.waitIndex = queryMeta.LastIndex
	//log.Println("queryMeta asyncResolve",queryMeta.LastIndex,cb.waitIndex)
	adds := make([]resolver.Address, 0)
	for _, serviceEntry := range serviceEntries {
		address := resolver.Address{Addr: fmt.Sprintf("%s:%d", serviceEntry.Service.Address, serviceEntry.Service.Port)}
		adds = append(adds, address)
	}
	return adds, "", nil
}


func (cb *consulBuilder) Scheme() string {
	return "consul"
}

type consulResolver struct {
	clientConn           *resolver.ClientConn
	consulBuilder        *consulBuilder
	t                    *time.Ticker
	wg                   sync.WaitGroup
	rn                   chan struct{}
	ctx                  context.Context
	cancel               context.CancelFunc
	disableServiceConfig bool
}
// newConsulResolver 创建服务发现
func newConsulResolver(cc *resolver.ClientConn, cb *consulBuilder, opts resolver.BuildOption) *consulResolver {
	ctx, cancel := context.WithCancel(context.Background())
	return &consulResolver{
		clientConn:           cc,
		consulBuilder:        cb,
		t:                    time.NewTicker(time.Second),
		ctx:                  ctx,
		cancel:               cancel,
		disableServiceConfig: opts.DisableServiceConfig}
}

func (cr *consulResolver) watcher() {
	cr.wg.Done()
	for {
		select {
		case <-cr.ctx.Done():
			return
		case <-cr.rn:
		case <-cr.t.C:
		}
		adds, _, err := cr.consulBuilder.resolve()
		if err != nil {
			log.Fatal("query service entries error:", err.Error())
		}
		//(*cr.clientConn).NewAddress(adds)
		//(*cr.clientConn).NewServiceConfig(serviceConfig)
		(*cr.clientConn).UpdateState(resolver.State{Addresses:adds})
	}
}

func (cr *consulResolver) Scheme() string {
	return cr.consulBuilder.Scheme()
}

func (cr *consulResolver) ResolveNow(rno resolver.ResolveNowOption) {
	select {
	case cr.rn <- struct{}{}:
	default:
	}
}

func (cr *consulResolver) Close() {
	cr.cancel()
	cr.wg.Wait()
	cr.t.Stop()
}

type consulClientConn struct {
	adds []resolver.Address
	sc   string
	state resolver.State
}

// NewConsulClientConn 创建服务发现客户端连接
func NewConsulClientConn() resolver.ClientConn {
	return &consulClientConn{}
}

// NewAddress 已经废弃
func (cc *consulClientConn) NewAddress(addresses []resolver.Address) {
	//log.Println("NewAddress",addresses)
	cc.adds = addresses
}
// NewServiceConfig 已经废弃
func (cc *consulClientConn) NewServiceConfig(serviceConfig string) {
	//log.Println("NewServiceConfig",serviceConfig)
	cc.sc = serviceConfig
}

// UpdateState 新版本中依靠它来更新服务状态
func (cc *consulClientConn) UpdateState(state resolver.State){
	//log.Println("UpdateState",state)
	cc.state = state
}

// StartConsulResolver 启动服务发现
func StartConsulResolver(consulAddr string, serviceName string) (schema string, err error) {
	builder := NewConsulBuilder(consulAddr)
	target := resolver.Target{Scheme: builder.Scheme(), Endpoint: serviceName}
	_, err = builder.Build(target, NewConsulClientConn(), resolver.BuildOption{})
	if err != nil {
		return builder.Scheme(), err
	}
	resolver.Register(builder)
	schema = builder.Scheme()
	return
}