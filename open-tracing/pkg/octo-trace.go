package pkg

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	"io"
)

func InitJager(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 100,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},

		Gen128Bit: true,
	}

	trace, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("cfg.NewTracer err:%v\n", err))
	}
	return trace, closer
}


func printTag(span opentracing.Span){
	span.Tracer().StartSpan("")
}