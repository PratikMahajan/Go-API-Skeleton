package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics holds all the available internal metrics
type Metrics struct {
	// APIResponseDurationsMilliseconds is the number of milliseconds it takes to
	// complete API responses.
	//
	// Labels: path (request path), method (request HTTP method),
	// status_code (response HTTP status code)
	APIResponseDurationsMilliseconds *prometheus.HistogramVec

	// APIHandlerPanicsTotal is the number of times HTTP request handlers have panicked.
	//
	// Labels: path(request path), method( request HTTP method)
	APIHandlerPanicsTotal *prometheus.CounterVec

	// JobsRunDurationsMilliseconds is the number of milliseconds jobs run for.
	//
	// Labels: job_type (jobs.JobStartRequest.Type field), successful (0 = fail, 1 = success)
	JobsRunDurationsMilliseconds *prometheus.HistogramVec
}

// NewMetrics creates a Metrics struct with all the Prometheus metrics recorders initialized
func NewMetrics() Metrics {
	metrics := Metrics{
		APIResponseDurationsMilliseconds: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "GoAppTwitter",
			Subsystem: "api",
			Name:      "response_durations_milliseconds",
			Help:      "Time, in milliseconds, it took to respond to API requests",
		}, []string{"path", "method", "status_code"}),
		APIHandlerPanicsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "GoAppTwitter",
			Subsystem: "api",
			Name:      "handler_panics_total",
			Help:      "Total number of HTTP handlers which have panicked while processing a request",
		}, []string{"path", "method"}),
		JobsRunDurationsMilliseconds: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "GoAppTwitter",
			Subsystem: "jobs",
			Name:      "run_durations_milliseconds",
			Help:      "Duration, in milliseconds, of jobs",
		}, []string{"job_type", "successful"}),
	}

	prometheus.MustRegister(metrics.APIResponseDurationsMilliseconds)
	prometheus.MustRegister(metrics.APIHandlerPanicsTotal)
	prometheus.MustRegister(metrics.JobsRunDurationsMilliseconds)

	return metrics
}

// StartTimer starts a Timer for the provided Prometheus observer. Calling .Finish()
// on the returned timer will records the time elapsed in milliseconds.
func (m Metrics) StartTimer() Timer {
	return Timer{
		startTime: time.Now(),
	}
}