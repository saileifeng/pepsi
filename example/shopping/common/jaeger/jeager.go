package jaeger

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"log"
)

var Closer io.Closer

func InitJaeger() (tracer opentracing.Tracer,closer io.Closer,err error){
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		// parsing errors might happen here, such as when we get a string where we expect a number
		log.Printf("Could not parse Jaeger env vars: %s", err.Error())
		return
	}

	tracer, closer, err = cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	opentracing.SetGlobalTracer(tracer)
	Closer = closer
	return
}

