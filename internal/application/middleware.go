package application

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/adolsalamanca/go-clean-boilerplate/pkg/logger"
)

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

		next.ServeHTTP(w, r)
		elapsed := time.Since(start)

		s.logger.Info("request finished",
			logger.NewFieldInt("duration_ms", int(elapsed.Milliseconds())),
			logger.NewFieldString("host", r.Host),
			logger.NewFieldString("headers", headers),
			logger.NewFieldString("method", r.Method),
		)
	})
}
