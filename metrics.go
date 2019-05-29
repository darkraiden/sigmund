package sigmund

import (
	log "github.com/sirupsen/logrus"
)

// Metric is a type that will be used to serialise the metric types received from SNS
type Metric int

type dbKey struct {
	metric string
	value  bool
}

const (
	// LowMemory is a metric that indicates low memory
	LowMemory Metric = iota
	// LowCPU is a metric that indicates low CPU
	LowCPU
	// OkMemory is a metric that indicates the absence of memory issues
	OkMemory
	// OkCPU is a metric that indicates the absence of CPU issues
	OkCPU
)

var stringsToMetrics = map[string]Metric{
	"LowMemory": LowMemory,
	"LowCPU":    LowCPU,
	"OkMemory":  OkMemory,
	"OkCPU":     OkCPU,
}

func (m Metric) String() string {
	metricNames := [...]string{
		"LowMemory",
		"LowCPU",
		"OkMemory",
		"OkCPU",
	}

	if m < 0 || m > 3 {
		log.WithField("integer", m).Panic("Value cannot be interpreted as a metric")
	}

	return metricNames[m]
}

var metricsTodbKey = map[Metric]dbKey{
	LowMemory: {
		metric: "isLowMemory",
		value:  true,
	},
	LowCPU: {
		metric: "isLowCPU",
		value:  true,
	},
	OkMemory: {
		metric: "isLowMemory",
		value:  false,
	},
	OkCPU: {
		metric: "isLowCPU",
		value:  false,
	},
}
