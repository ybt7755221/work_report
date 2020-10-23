package jaegerMid

import (
	"context"
	"work_report/libraries/jaeger"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Listen() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer, closer := jaeger.InitJaeger("work_report")
		defer closer.Close()
		path := c.Request.URL.Path
		span := tracer.StartSpan(path, ext.SpanKindRPCServer)
		ext.HTTPUrl.Set(span, c.Request.Method)
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		c.Set("ctx", ctx)
		c.Next()
		ext.HTTPStatusCode.Set(span, uint16(c.Writer.Status()))
		span.Finish()
	}
}
