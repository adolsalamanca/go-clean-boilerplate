package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel int

const (
	DebugLevel = LogLevel(zapcore.DebugLevel)
	InfoLevel  = LogLevel(zapcore.InfoLevel)
	WarnLevel  = LogLevel(zapcore.WarnLevel)
	ErrorLevel = LogLevel(zapcore.ErrorLevel)
	PanicLevel = LogLevel(zapcore.PanicLevel)
	FatalLevel = LogLevel(zapcore.FatalLevel)
)

type FieldType uint8

type Field struct {
	// Any    []byte
	Key    string
	String string
	Int    int
	Float  float32
}

/*
func NewFieldAny(key string, value []byte) Field {
	return Field{
		Key: key,
		Any: value,
	}
}
*/

func NewFieldString(key, value string) Field {
	return Field{
		Key:    key,
		String: value,
	}
}

func NewFieldInt(key string, value int) Field {
	return Field{
		Key: key,
		Int: value,
	}
}

func NewFieldFloat(key string, float float32) Field {
	return Field{
		Key:   key,
		Float: float,
	}
}

type StandardLogger struct {
	logger *zap.Logger
}

func (s StandardLogger) Debug(msg string, fields ...Field) {
	zapFields := toZapFields(fields)

	s.logger.Debug(msg, zapFields[:]...)
}

func (s StandardLogger) Info(msg string, fields ...Field) {
	zapFields := toZapFields(fields)

	s.logger.Info(msg, zapFields[:]...)
}

func (s StandardLogger) Warn(msg string, fields ...Field) {
	zapFields := toZapFields(fields)

	s.logger.Warn(msg, zapFields[:]...)
}

func (s StandardLogger) Error(msg string, fields ...Field) {
	zapFields := toZapFields(fields)

	s.logger.Error(msg, zapFields[:]...)
}

func NewLogger(level LogLevel) *StandardLogger {
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
	atom.SetLevel(zapcore.Level(level))

	return &StandardLogger{
		logger: logger,
	}
}

func toZapFields(fields []Field) []zap.Field {
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
			if f.Float != 0 {
				zapFields = append(zapFields, zap.Float32p(f.Key, &f.Float))
				continue
			}
			/*
				if len(f.Any) > 0 {
								zapFields = append(zapFields, zap.Any(f.Key, f.Any))
							}
			*/
		}
	}
	return zapFields
}

// NoOpLogger fulfills the required Logger interface without logging at all
type NoOpLogger struct{}

func (s NoOpLogger) Debug(msg string, fields ...Field) {}

func (s NoOpLogger) Info(msg string, fields ...Field) {}

func (s NoOpLogger) Warn(msg string, fields ...Field) {}

func (s NoOpLogger) Error(msg string, fields ...Field) {}

func NewNoOpLogger() *NoOpLogger {
	return &NoOpLogger{}
}
