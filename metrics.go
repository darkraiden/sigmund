package sigmund

import (
	"fmt"
)

func identifyMetric(m string) (string, error) {
	var k string
	var err error
	switch {
	case m == "memory":
		k = "isLowMemory"
	case m == "cpu":
		k = "isLowCPU"
	default:
		err = fmt.Errorf("%v is an invalid parameter. Please provide either 'memory' or 'cpu' keys", m)
	}
	return k, err
}
