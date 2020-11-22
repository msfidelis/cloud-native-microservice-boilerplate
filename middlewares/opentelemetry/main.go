package opentracing

import (
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Load() gin.HandlerFunc {

	return func(c *gin.Context) {

		operationPrefix := os.Getenv("JAEGER_SERVICE_NAME")

		// all before request is handled
		var span opentracing.Span
		if cspan, ok := c.Get("tracing-context"); ok {
			span = StartSpanWithParent(cspan.(opentracing.Span).Context(), string(operationPrefix)+c.Request.Method, c.Request.Method, c.Request.URL.Path)

		} else {
			span = StartSpanWithHeader(&c.Request.Header, string(operationPrefix)+c.Request.Method, c.Request.Method, c.Request.URL.Path)
		}
		defer span.Finish()            // after all the other defers are completed.. finish the span
		c.Set("tracing-context", span) // add the span to the context so it can be used for the duration of the request.
		c.Next()

		span.SetTag(string(ext.HTTPStatusCode), c.Writer.Status())

	}

}

func StartSpanWithParent(parent opentracing.SpanContext, operationName, method, path string) opentracing.Span {
	options := []opentracing.StartSpanOption{
		opentracing.Tag{Key: ext.SpanKindRPCServer.Key, Value: ext.SpanKindRPCServer.Value},
		opentracing.Tag{Key: string(ext.HTTPMethod), Value: method},
		opentracing.Tag{Key: string(ext.HTTPUrl), Value: path},
		opentracing.Tag{Key: "current-goroutines", Value: runtime.NumGoroutine()},
	}

	if parent != nil {
		options = append(options, opentracing.ChildOf(parent))
	}

	return opentracing.StartSpan(operationName, options...)
}

func StartSpanWithHeader(header *http.Header, operationName, method, path string) opentracing.Span {
	var wireContext opentracing.SpanContext
	if header != nil {
		wireContext, _ = opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(*header))
	}
	return StartSpanWithParent(wireContext, operationName, method, path)
}
