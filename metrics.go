package sigmund

import "fmt"

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

func (s *Sigmund) identifyDBKey() (*dbKey, error) {
	switch metricsToStrings[s.Dynamo.Key] {
	case "LowMemory":
		return &dbKey{
			"isLowMemory",
			true,
		}, nil

	case "OkMemory":
		return &dbKey{
			"isLowMemory",
			false,
		}, nil

	case "LowCPU":
		return &dbKey{
			"isLowCPU",
			true,
		}, nil

	case "OkCPU":
		return &dbKey{
			"isLowCPU",
			false,
		}, nil
	default:
		return nil, fmt.Errorf("I can't identify the Key to be used for this DB")
	}
}
