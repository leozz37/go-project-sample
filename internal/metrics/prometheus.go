package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	SuccessCacheGet = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sample_cache_get_sucessfully_total",
		Help: "Total cache usage",
	})
	CacheSet = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sample_cache_set_total",
		Help: "Total persistence to cache",
	})
)
