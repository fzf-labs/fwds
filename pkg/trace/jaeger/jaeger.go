package jaeger

import (
	"fwds/internal/conf"
	"io"

	"fwds/pkg/log"

	"github.com/opentracing/opentracing-go"
	jaegerCli "github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	jaegerMet "github.com/uber/jaeger-lib/metrics"
)

func NewJaeger(serviceName string, cfg *conf.JaegerConfig) (opentracing.Tracer, io.Closer, error) {
	reporterConfig := &jaegerCfg.ReporterConfig{
		LogSpans:           true,
		LocalAgentHostPort: cfg.LocalAgentHostPort,
		CollectorEndpoint:  cfg.CollectorEndpoint,
		User:               cfg.CollectorUser,
		Password:           cfg.CollectorPassword,
	}
	jcfg := jaegerCfg.Configuration{
		Sampler: &jaegerCfg.SamplerConfig{
			Type:              cfg.SamplingType,
			Param:             cfg.SamplingParam,
			SamplingServerURL: cfg.SamplingServerURL,
		},
		Reporter: reporterConfig,
		Headers: &jaegerCli.HeadersConfig{
			TraceContextHeaderName: cfg.TraceContextHeaderName,
		},
	}
	jMetricsFactory := jaegerMet.NullFactory

	opts := []jaegerCfg.Option{
		jaegerCfg.Metrics(jMetricsFactory),
		jaegerCfg.Gen128Bit(cfg.Gen128Bit),
	}
	// Initialize tracer with a logger and a metrics factory
	closer, err := jcfg.InitGlobalTracer(
		serviceName,
		opts...,
	)
	if err != nil {
		log.SugaredLogger.Errorf("Could not initialize jaeger tracer: %s", err.Error())
		return nil, nil, err
	}
	return opentracing.GlobalTracer(), closer, nil
}

func Init(cfg *conf.Config) {
	_, _, err := NewJaeger(cfg.Trace.ServiceName, &cfg.Jaeger)
	if err != nil {
		log.SugaredLogger.Error("tracer init err:", err)
	}
}
