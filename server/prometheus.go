package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(kvCounter)
	prometheus.MustRegister(setKVDurationHistogram)
	prometheus.MustRegister(saveVersionDurationHistogram)
}

func recordQueryMetrics() {
	kvCounter.Inc()
}

var (
	kvCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Subsystem: "iavl",
		Name:      "total_kv",
		Help:      "Total number of KeyValue",
	})
)

func recordSetKVDurationMetrics(startTime time.Time) {
	duration := time.Since(startTime)
	setKVDurationHistogram.Observe(duration.Seconds())
}

var (
	setKVDurationHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Subsystem: "iavl",
		Name:      "set_kv_duration_seconds",
		Help:      "Duration of set kv in seconds",
		Buckets:   []float64{0.05, 0.1, 0.25, 0.5, 0.75, 1},
	})
)

func recordSaveVersionDurationMetrics(startTime time.Time) {
	duration := time.Since(startTime)
	saveVersionDurationHistogram.Observe(duration.Seconds())
}

var (
	saveVersionDurationHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Subsystem: "iavl",
		Name:      "save_version_duration_seconds",
		Help:      "Duration of save version iavl in seconds",
		Buckets:   []float64{0.05, 0.1, 0.25, 0.5, 0.75, 1},
	})
)
