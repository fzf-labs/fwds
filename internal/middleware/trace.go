package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

// Trace
// @Description: 链路跟踪
// @return gin.HandlerFunc
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer()
		var newCtx context.Context
		var span opentracing.Span
		//http
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), tracer, c.Request.URL.Path)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), tracer, c.Request.URL.Path, opentracing.ChildOf(spanCtx), opentracing.Tag{Key: string(ext.Component), Value: "HTTP"})
		}

		ext.HTTPMethod.Set(span, c.Request.Method)
		ext.HTTPUrl.Set(span, c.Request.URL.String())
		// add trace id and span id
		// get trace id and span id by using log
		var traceID string
		var spanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext, ok := spanContext.(jaeger.SpanContext)
			if ok {
				traceID = jaegerContext.TraceID().String()
				spanID = jaegerContext.SpanID().String()
			}
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", spanID)

		c.Request = c.Request.WithContext(newCtx)

		c.Next()

		// record HTTP status code
		ext.HTTPStatusCode.Set(span, uint16(c.Writer.Status()))

		span.Finish()
	}
}
