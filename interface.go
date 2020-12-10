package trace

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// Trace interface def.
type Trace interface {
	SetGlobalTracer()
	StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span
	Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error
	Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error)
	Close() error
}

var _ Trace = &TraceImpl{}

// MakeTrace return the point of TraceImpl
func MakeTrace(cfg config.Configuration) (*TraceImpl, error) {
	tracer, closer, err := cfg.NewTracer(
		config.Logger(log.NullLogger),
		config.Metrics(metrics.NullFactory),
	)
	if err != nil {
		return nil, err
	}
	return &TraceImpl{
		tracer: tracer,
		closer: closer,
	}, nil
}
