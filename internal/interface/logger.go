package _interface

import (
	"os"

	"github.com/adolsalamanca/go-clean-boilerplate/internal/application"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type StandardLogger struct {
	logger *zap.Logger
}

func (s StandardLogger) Debug(msg string, fields ...application.Field) {
	zapFields := toZapFields(fields)

	s.logger.Debug(msg, zapFields[:]...)
}

func (s StandardLogger) Info(msg string, fields ...application.Field) {
	zapFields := toZapFields(fields)

	s.logger.Info(msg, zapFields[:]...)
}

func (s StandardLogger) Warn(msg string, fields ...application.Field) {
	zapFields := toZapFields(fields)

	s.logger.Warn(msg, zapFields[:]...)
}

func (s StandardLogger) Error(msg string, fields ...application.Field) {
	zapFields := toZapFields(fields)

	s.logger.Error(msg, zapFields[:]...)
}

func NewLogger() *StandardLogger {

	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))

	defer logger.Sync()
	atom.SetLevel(zap.DebugLevel)

	return &StandardLogger{
		logger: logger,
	}
}

func toZapFields(fields []application.Field) []zap.Field {
	var zapFields []zap.Field
	if len(fields) > 0 {
		for _, f := range fields {
			if f.String != "" {
				zapFields = append(zapFields, zap.String(f.Key, f.String))
				continue
			}
			if f.Int != 0 {
				zapFields = append(zapFields, zap.Int(f.Key, f.Int))
				continue
			}

		}
	}
	return zapFields
}
