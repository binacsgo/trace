package trace

import (
	"io"

	"github.com/opentracing/opentracing-go"
)

// TraceImpl implement the Trace
type TraceImpl struct {
	tracer opentracing.Tracer
	closer io.Closer
}

// SetGlobalTracer set the global tracer
func (t *TraceImpl) SetGlobalTracer() {
	opentracing.SetGlobalTracer(t.tracer)
}

// StartSpan return the tracer's impl
func (t *TraceImpl) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	return t.tracer.StartSpan(operationName, opts...)
}

// Inject return the tracer's impl
func (t *TraceImpl) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	return t.tracer.Inject(sm, format, carrier)
}

// Extract return the tracer's impl
func (t *TraceImpl) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	return t.tracer.Extract(format, carrier)
}

// Close the closer
func (t *TraceImpl) Close() error {
	return t.closer.Close()
}
