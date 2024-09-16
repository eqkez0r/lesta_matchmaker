package middleware

import (
	"github.com/eqkez0r/lesta_matchmaker/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(
	l logger.ILogger,
) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		data := &responseData{
			status: 0,
			size:   0,
		}
		lw := loggingResponseWriter{
			ResponseWriter: c.Writer,
			resData:        data,
		}

		c.Writer = &lw

		c.Next()

		data.status = c.Writer.Status()
		data.size = lw.Size()

		duration := time.Since(start)

		l.Info(
			"URL ", c.Request.URL,
			" METHOD ", c.Request.Method,
			" STATUS ", data.status,
			" SIZE ", data.size,
			" DURATION ", duration,
		)
	}
}

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		gin.ResponseWriter
		resData *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := r.ResponseWriter.Write(b)
	r.resData.size += size // захватываем размер
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	// записываем код статуса, используя оригинальный http.ResponseWriter
	r.ResponseWriter.WriteHeader(statusCode)
	r.resData.status = statusCode // захватываем код статуса
}
