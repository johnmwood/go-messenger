package logger

import (
	"fmt"
	"runtime"
	"time"
)

type LogReporter interface {
	Report() string
}

type MetricSampler interface {
	Sample(time.Duration)
}

type Sampler struct {
	metric     Metric
	gatherFunc func() int
}

func NewSampler(m Metric, fn func() int) *Sampler {
	return &Sampler{
		metric:     m,
		gatherFunc: fn,
	}
}

func (s *Sampler) Sample(rate time.Duration) {
	for {
		time.Sleep(rate)
		s.metric.amount = int64(s.gatherFunc())
		fmt.Println(runtime.NumGoroutine())
		fmt.Printf("updated sampler %v with new amount\n", s.metric.name)
	}
}

type MetricLogger struct {
	*Sampler
}
type Metric struct {
	name   string
	amount int64
}

func (m *Metric) report() string {
	return fmt.Sprintf("%v: %d", m.name, m.amount)
}

func NewMetricLogger() *MetricLogger {
	cpu := Metric{
		name:   "cpu",
		amount: 0,
	}
	metricsLogger := MetricLogger{
		Sampler: NewSampler(cpu, runtime.NumGoroutine),
	}
	go metricsLogger.Sample(time.Second * 1)

	return &metricsLogger
}

func (l *MetricLogger) Report() string {
	return l.metric.report()
}
