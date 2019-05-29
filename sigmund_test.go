package sigmund

import (
	"errors"
	"testing"
)

type identifyMetrics struct {
	metric   string
	expected error
}

var identifyPositiveMetrics = []identifyMetrics{
	identifyMetrics{
		metric:   "cpu",
		expected: nil,
	},
	identifyMetrics{
		metric:   "memory",
		expected: nil,
	},
}

var identifyNegativeMetrics = []identifyMetrics{
	identifyMetrics{
		metric:   "CPU",
		expected: errors.New("CPU is an invalid parameter. Please provide either 'memory' or 'cpu' keys"),
	},
	identifyMetrics{
		metric:   "SuperMan",
		expected: errors.New("CPU is an invalid parameter. Please provide either 'memory' or 'cpu' keys"),
	},
	identifyMetrics{
		metric:   "",
		expected: errors.New(" is an invalid parameter. Please provide either 'memory' or 'cpu' keys"),
	},
}

// Positive Tests for Metric Identifier function
func TestPositiveIdentifyMetric(t *testing.T) {
	for _, im := range identifyPositiveMetrics {
		result, err := identifyMetric(im.metric)
		if err != nil {
			t.Errorf("Invalid parameter! Expected: %s, got: %s", im.metric, result)
		}
	}
}

// Negative Tests for Metric Identifier function
func TestNegativeIdentifyMetric(t *testing.T) {
	for _, im := range identifyNegativeMetrics {
		_, err := identifyMetric(im.metric)
		if err == nil {
			t.Errorf("Invalid error message! Expected: \"%s is an invalid parameter. Please provide either 'memory' or 'cpu' keys\", got: %v", im.metric, err)
		}
	}
}
