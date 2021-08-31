package _interface

type MetricsCollector interface {
	AddSample(key []string, val float32)
	IncrCounter(key []string, val float32)
	SetGauge(key []string, val float32)
}
