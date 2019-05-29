package sigmund

// Metric is a type that will be used to
// serialise the metric types
// received from SNS
type Metric int

const (
	lowMemory Metric = iota
	lowCPU
	okMemory
	okCPU
)

var stringsToMetrics = map[string]Metric{
	"LowMemory": lowMemory,
	"LowCPU":    lowCPU,
	"OkMemory":  okMemory,
	"OkCPU":     okCPU,
}

var metricsToStrings = map[Metric]string{
	lowMemory: "LowMemory",
	lowCPU:    "LowCPU",
	okMemory:  "OkMemory",
	okCPU:     "OkCPU",
}

type dbKey struct {
	metric string
	value  bool
}

var metricsTodbKey = map[Metric]dbKey{
	lowMemory: {
		metric: "isLowMemory",
		value:  true,
	},
	lowCPU: {
		metric: "isLowCPU",
		value:  true,
	},
	okMemory: {
		metric: "isLowMemory",
		value:  false,
	},
	okCPU: {
		metric: "isLowCPU",
		value:  false,
	},
}
