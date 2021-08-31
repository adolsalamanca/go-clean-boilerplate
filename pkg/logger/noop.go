package logger

// NoOpLogger fulfills the required Logger interface without logging at all
type NoOpLogger struct{}

func (s NoOpLogger) Debug(msg string, fields ...Field) {}

func (s NoOpLogger) Info(msg string, fields ...Field) {}

func (s NoOpLogger) Warn(msg string, fields ...Field) {}

func (s NoOpLogger) Error(msg string, fields ...Field) {}

func NewNoOpLogger() *NoOpLogger {
	return &NoOpLogger{}
}
