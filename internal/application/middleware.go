package application

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/adolsalamanca/go-clean-boilerplate/pkg/logger"
)

type ResponseTracker struct {
	responseCode int
	responseBody []byte
	w            http.ResponseWriter
}

func (rt *ResponseTracker) Header() http.Header {
	return rt.w.Header()
}

func (rt *ResponseTracker) Write(b []byte) (int, error) {
	rt.responseBody = b
	return rt.w.Write(b)
}

func (rt *ResponseTracker) WriteHeader(statusCode int) {
	rt.responseCode = statusCode
	rt.w.WriteHeader(statusCode)
}

func (s *Server) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var headersArr []string
		for key, val := range r.Header {
			headersArr = append(headersArr, fmt.Sprintf("%s: %s", key, val))
		}
		headers := strings.Join(headersArr, ",")

		s.logger.Info("request arrived",
			logger.NewFieldString("host", r.Host),
			logger.NewFieldString("headers", headers),
			logger.NewFieldString("method", r.Method),
		)
		s.collector.IncrCounter([]string{fmt.Sprintf("http_requests_%s%s", r.Method, r.URL.Path)}, 1)

		rt := &ResponseTracker{
			responseCode: 0,
			responseBody: make([]byte, 0),
			w:            w,
		}
		next.ServeHTTP(rt, r)
		elapsed := time.Since(start)

		s.logger.Info("request finished",
			logger.NewFieldInt("response_code", rt.responseCode),
			logger.NewFieldFloat("duration_µs", float32(elapsed.Microseconds())),
			logger.NewFieldString("host", r.Host),
			logger.NewFieldString("headers", headers),
			logger.NewFieldString("method", r.Method),
		)

		// TODO: Print JSON fields directly, implement Object/Array marshallers or use any.
		// Info: https://github.com/uber-go/zap/issues/405
		/*
			s.logger.Info("request response",
				logger.NewFieldAny("body", rt.responseBody),
			)
		*/

		s.collector.SetGauge([]string{"http_request_latency"}, float32(elapsed.Microseconds()))
		s.collector.SetGauge([]string{"http_request_response_code"}, float32(rt.responseCode))
	})
}
