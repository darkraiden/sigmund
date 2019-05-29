package sigmund

import "errors"

// Autoscaling represents the structure of a series
// of AWS autoscaling group parameters to trigger
// a dimension change
type Autoscaling struct {
	AutoScalingGroupName string
	PolicyName           string
	Region               string
}

// Dynamo contains the DyanamoDB basic information
// needed to connect to the instance
type Dynamo struct {
	TableName string
	Region    string
	Key       Metric
}

// Sigmund is a struct coontaining info regarding
// the autoscaling group and
// the Datastore used by the application
type Sigmund struct {
	Autoscaling
	Dynamo
}

// DBItem represents the structure of the JSON body coming back from DynamoDB after
// a Select query is executed
type DBItem struct {
	ID          int  `json:"ID"`
	IsLowMemory bool `json:"isLowMemory"`
	IsLowCPU    bool `json:"isLowCPU"`
}

// New is the Package constructor that initialises
// the Sigmund config
func New(region, asgName, policyName, tableName, metric string) (*Sigmund, error) {
	metricDBKey, err := checkConfig(region, asgName, policyName, tableName, metric)
	if err != nil {
		return nil, err
	}

	return &Sigmund{
		Autoscaling{
			AutoScalingGroupName: asgName,
			PolicyName:           policyName,
			Region:               region,
		},
		Dynamo{
			TableName: tableName,
			Region:    region,
			Key:       metricDBKey,
		},
	}, nil
}

func checkConfig(region, asgName, policyName, tableName, metric string) (Metric, error) {
	val, okMetric := stringsToMetrics[metric]
	switch {
	case region == "":
		return val, errors.New("Region cannot be empty")
	case asgName == "":
		return val, errors.New("The autoscaling group name cannot be empty")
	case policyName == "":
		return val, errors.New("The autoscaling policy cannot be empty")
	case tableName == "":
		return val, errors.New("Table name cannot be empty")
	case metric == "":
		return val, errors.New("Key cannot be empty")
	case !okMetric:
		return val, errors.New("Invalid metric Key")
	default:
		return val, nil
	}
}
