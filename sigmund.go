package sigmund

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

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
	Key       string
}

// Logger wraps a logger, so the correct logging context
// can be passed from the main package into Sigmund
type Logger struct {
	Log *logrus.Logger
}

// Sigmund is a struct containing info regarding
// the autoscaling group and
// the Datastore used by the application
type Sigmund struct {
	Autoscaling
	Dynamo
	Logger
}

// DBItem represents the structure of the JSON body coming back from DynamoDB after
// a Select query is executed
type DBItem struct {
	ID          int  `json:"ID"`
	IsLowMemory bool `json:"isLowMemory"`
	IsLowCPU    bool `json:"isLowCPU"`
}

// Config is a struct containing all the info required by the constructor function
type Config struct {
	Region     string
	AsgName    string
	PolicyName string
	TableName  string
	Metric     Metric
	Log        *logrus.Logger
}

// New is the Package constructor that takes in
// a Config and returns a pointer to Sigmund
func New(conf *Config) *Sigmund {
	checkConfig(conf)

	metricDBKey := conf.Metric

	return &Sigmund{
		Autoscaling{
			AutoScalingGroupName: conf.AsgName,
			PolicyName:           conf.PolicyName,
			Region:               conf.Region,
		},
		Dynamo{
			TableName: conf.TableName,
			Region:    conf.Region,
			Key:       metricDBKey,
		},
		Logger{
			Log: conf.Log,
		},
	}
}

func checkConfig(conf *Config) {

	fields := make(map[string]interface{})

	if conf.Region == "" {
		fields["region"] = "empty"
		//return errors.New("Region cannot be empty")
	}

	if conf.AsgName == "" {
		fields["autoscaling group name"] = "empty"
		//return errors.New("The autoscaling group name cannot be empty")
	}

	if conf.PolicyName == "" {
		fields["autoscaling policy"] = "empty"
		//return errors.New("The autoscaling policy cannot be empty")
	}

	if conf.TableName == "" {
		fields["table name"] = "empty"
		//return errors.New("Table name cannot be empty")
	}

	if len(fields) != 0 {
		conf.Log.WithFields(log.Fields(fields)).Panic("One or more environment variables are undefined")
	}
}
