package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// prometheus metrics setup
var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "gateway_trace_service",
			Subsystem: "gateway_trace_service",
			Name:      "http_requests_total",
			Help:      "A counter for requests.",
		},
		[]string{"code", "method"},
	)

	ResponseDurationHist = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "gateway_trace_service",
			Subsystem: "gateway_trace_service",
			Name:      "http_request_duration_seconds",
			Help:      "A histogram of request latencies.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"code", "method"},
	)

	WriteHeaderDurationHist = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "gateway_trace_service",
			Subsystem: "gateway_trace_service",
			Name:      "http_write_header_duration_seconds",
			Help:      "A histogram of time to first write latencies.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"code", "method"},
	)

	PrometheusPostRequestDurations = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "gateway_trace_service",
		Subsystem: "gateway_trace_service",
		Name:      "post_request_durations_seconds",
		Help:      "The duration of each post request",
		Buckets:   prometheus.LinearBuckets(0.01, 0.05, 10),
	})

	PrometheusGetRequestDurations = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "gateway_trace_service",
		Subsystem: "gateway_trace_service",
		Name:      "get_request_durations_seconds",
		Help:      "The duration of each get request",
		Buckets:   prometheus.LinearBuckets(0.01, 0.05, 10),
	})

	PrometheusPostRequestErrorCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "gateway_trace_service",
		Subsystem: "gateway_trace_service",
		Name:      "post_request_error_counter",
		Help:      "The number of accumalative post errors",
	})

	PrometheusGetRequestErrorCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "gateway_trace_service",
		Subsystem: "gateway_trace_service",
		Name:      "get_request_error_counter",
		Help:      "The number of accumalative get errors",
	})

	PrometheusGetRequestElasticSearchFailureCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "gateway_trace_service",
		Subsystem: "gateway_trace_service",
		Name:      "get_request_elastic_search_failure_counter",
		Help:      "The number of accumalative elastic search failure of get request",
	})

	PrometheusPostRequestElasticSearchFailureCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "gateway_trace_service",
		Subsystem: "gateway_trace_service",
		Name:      "post_request_elastic_search_failure_counter",
		Help:      "The number of accumalative elastic search failure of post request",
	})

	PrometheusPostTraceIndicator = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "gateway_trace_service",
		Subsystem: "gateway_trace_service",
		Name:      "post_logs_number_counter",
		Help:      "The number of accumulative number of posted trace logs",
	})
)

func init() {
	prometheus.MustRegister(RequestCounter, ResponseDurationHist, WriteHeaderDurationHist, PrometheusGetRequestDurations, PrometheusPostRequestDurations, PrometheusPostRequestErrorCounter, PrometheusGetRequestErrorCounter, PrometheusGetRequestElasticSearchFailureCounter, PrometheusPostRequestElasticSearchFailureCounter, PrometheusPostTraceIndicator)
}
