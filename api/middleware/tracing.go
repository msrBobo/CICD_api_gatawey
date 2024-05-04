package middleware

import (
	"bufio"
	"dennic_api_gateway/api/response"
	"errors"
	"go.opentelemetry.io/otel"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
)

const RequestIDHeader = "X-Request-ID"

// Tracing middleware function
func GinTracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		rw := response.NewResponseWriter(c.Writer, http.StatusOK)
		// Tracing
		ctx := c.Request.Context()
		tracer := otel.Tracer("")
		ctx, span := tracer.Start(ctx, c.FullPath()) // Use Gin's FullPath
		defer span.End()

		// Add request ID to header
		c.Writer.Header().Add(RequestIDHeader, span.SpanContext().TraceID().String())

		// Serve the next middleware
		c.Next()

		// Add attributes
		span.SetAttributes(
			attribute.Key("http.method").String(c.Request.Method),
			attribute.Key("http.url").String(c.FullPath()), // Use Gin's FullPath
			attribute.Key("http.status_code").Int(rw.StatusCode()),
		)
	}
}

// Custom responseWriter to capture status code
type responseWriter struct {
	gin.ResponseWriter
	statusCode int
}

// WriteHeader method implements the http.ResponseWriter interface
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Write method implements the http.ResponseWriter interface
func (rw *responseWriter) Write(data []byte) (int, error) {
	return rw.ResponseWriter.Write(data)
}

// WriteString method implements the gin.ResponseWriter interface
func (rw *responseWriter) WriteString(s string) (int, error) {
	return rw.ResponseWriter.WriteString(s)
}

// Hijack method implements the http.Hijacker interface (optional)
func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := rw.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, errors.New("response writer does not support hijacking")
}

// CloseNotify method implements the http.CloseNotifier interface (optional)
func (rw *responseWriter) CloseNotify() <-chan bool {
	if notifier, ok := rw.ResponseWriter.(http.CloseNotifier); ok {
		return notifier.CloseNotify()
	}
	// Return a dummy channel if CloseNotify is not supported
	return make(chan bool)
}
