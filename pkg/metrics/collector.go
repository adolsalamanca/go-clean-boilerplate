package metrics

import (
	"runtime"
	"time"

	"github.com/armon/go-metrics"
)

type Collector struct {
	statsdUDPAgent *metrics.StatsdSink
	serviceName    string
}

func NewMetricsCollector(address string, serviceName string) (*Collector, error) {
	statsdSink, err := metrics.NewStatsdSink(address)
	if err != nil {
		return nil, err
	}
	metrics.NewGlobal(metrics.DefaultConfig(serviceName), statsdSink)

	collector := &Collector{
		statsdUDPAgent: statsdSink,
		serviceName:    serviceName,
	}
	go readMetrics(collector)

	return collector, nil
}

func (c Collector) AddSample(key []string, val float32) {
	usedKey := append([]string{c.serviceName}, key...)
	c.statsdUDPAgent.AddSample(usedKey, val)
}

func (c Collector) IncrCounter(key []string, val float32) {
	usedKey := append([]string{c.serviceName}, key...)
	c.statsdUDPAgent.IncrCounter(usedKey, val)
}

func (c Collector) SetGauge(key []string, val float32) {
	usedKey := append([]string{c.serviceName}, key...)
	c.statsdUDPAgent.SetGauge(usedKey, val)
}

func readMetrics(collector *Collector) {
	runtimeMemoryStats := runtime.MemStats{}
	for range time.Tick(time.Second) {
		runtime.ReadMemStats(&runtimeMemoryStats)
		collector.SetGauge([]string{"mem_alloc"}, float32(runtimeMemoryStats.Alloc))
		collector.SetGauge([]string{"goroutine_count"}, float32(runtime.NumGoroutine()))
		collector.SetGauge([]string{"mem_heap_objects"}, float32(runtimeMemoryStats.HeapObjects))
		collector.SetGauge([]string{"gc_pause_ns"}, float32(runtimeMemoryStats.PauseNs[(runtimeMemoryStats.NumGC+255)%256]))
	}
}
