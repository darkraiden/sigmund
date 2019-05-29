package sigmund

import (
	log "github.com/sirupsen/logrus"

	"github.com/darkraiden/sigmund/pkg/autoscaling"
	"github.com/darkraiden/sigmund/pkg/storage/dynamo"
)

// Shrink is the core function of the package
// which Executes an Autoscaling Group Policy
// when requirements are met
func (s *Sigmund) Shrink() {
	dbClient, err := s.newDBClient(metricsTodbKey[s.Key].metric)
	if err != nil {
		s.Log.WithError(err).Panic("Could not connect to DynamoDB")
	}

	// Run a Select query
	item, err := s.readDynamo(dbClient)
	if err != nil {
		s.Log.WithError(err).Panic("Could not read from DynamoDB")
	}

	switch s.Dynamo.Key.String() {
	case "LowMemory":
		err := dbClient.WriteToTable(s.Dynamo.TableName, "isLowMemory", true)
		if err != nil {
			s.Log.WithError(err).WithFields(log.Fields{
				"table name": s.Dynamo.TableName,
				"metric":     "isLowMemory == true",
			}).Panic()
		}
		item.IsLowMemory = true
	case "OkMemory":
		err := dbClient.WriteToTable(s.Dynamo.TableName, "isLowMemory", false)
		if err != nil {
			s.Log.WithError(err).WithFields(log.Fields{
				"table name": s.Dynamo.TableName,
				"metric":     "isLowMemory == false",
			}).Panic()
		}
		item.IsLowMemory = false
	case "LowCPU":
		err := dbClient.WriteToTable(s.Dynamo.TableName, "isLowCPU", true)
		if err != nil {
			s.Log.WithError(err).WithFields(log.Fields{
				"table name": s.Dynamo.TableName,
				"metric":     "isLowCPU == true",
			}).Panic()
		}
		item.IsLowCPU = true
	case "OkCPU":
		err := dbClient.WriteToTable(s.Dynamo.TableName, "isLowCPU", false)
		if err != nil {
			s.Log.WithError(err).WithFields(log.Fields{
				"table name": s.Dynamo.TableName,
				"metric":     "isLowCPU == false",
			}).Panic()
		}
		item.IsLowCPU = false
	}

	if item.IsLowCPU && item.IsLowMemory {
		err = dbClient.WriteToTable(s.Dynamo.TableName, "isLowCPU", false)
		if err != nil {
			s.Log.WithError(err).WithFields(log.Fields{
				"table name": s.Dynamo.TableName,
				"metric":     "isLowCPU == false",
			}).Panic()
		}
		err = dbClient.WriteToTable(s.Dynamo.TableName, "isLowMemory", false)
		if err != nil {
			s.Log.WithError(err).WithFields(log.Fields{
				"table name": s.Dynamo.TableName,
				"metric":     "isLowMemory == false",
			}).Panic()
		}
		s.execASGPolicy()
	}
}

func (s *Sigmund) newDBClient(key string) (*dynamo.Client, error) {
	dynamo, err := dynamo.New(s.Dynamo.TableName, s.Dynamo.Region, key)
	if err != nil {
		return nil, err
	}

	client, err := dynamo.NewClient()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *Sigmund) readDynamo(cli *dynamo.Client) (*DBItem, error) {
	item, err := cli.ReadFromTable(s.Dynamo.TableName)
	if err != nil {
		return nil, err
	}

	return &DBItem{
		ID:          item.ID,
		IsLowCPU:    item.IsLowCPU,
		IsLowMemory: item.IsLowMemory,
	}, nil
}

func (s *Sigmund) execASGPolicy() {
	var client *autoscaling.Client
	asg, err := autoscaling.New(s.Autoscaling.AutoScalingGroupName, s.Autoscaling.PolicyName, s.Autoscaling.Region)
	if err != nil {
		s.Log.WithError(err).WithFields(log.Fields{
			"asg name":    s.Autoscaling.AutoScalingGroupName,
			"policy name": s.Autoscaling.PolicyName,
			"region":      s.Autoscaling.Region,
		}).Panic("Failed to create autoscaling config")
	}

	client, err = asg.NewClient()
	if err != nil {
		s.Log.WithError(err).Panic("Failed to create autoscaling client")
	}

	err = client.TriggerPolicy(asg)
	if err != nil {
		s.Log.WithError(err).Panic("Failed to trigger asg policy")
	}
}
