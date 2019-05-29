package sigmund

import (
	"github.com/darkraiden/sigmund/pkg/autoscaling"
	"github.com/darkraiden/sigmund/pkg/storage/dynamo"
)

// Shrink is the core function of the package
// which Executes an Autoscaling Group Policy
// when requirements are met
func (s *Sigmund) Shrink() error {
	// Create a new DBClient
	dbClient, err := s.newDBClient(metricsTodbKey[s.Key].metric)
	if err != nil {
		return err
	}

	// Run a Select query
	item, err := s.readDynamo(dbClient)
	if err != nil {
		return err
	}

	switch metricsToStrings[s.Dynamo.Key] {
	case "LowMemory":
		err := dbClient.WriteToTable(s.Dynamo.TableName, "isLowMemory", true)
		if err != nil {
			return err
		}
		item.IsLowMemory = true
	case "OkMemory":
		err := dbClient.WriteToTable(s.Dynamo.TableName, "isLowMemory", false)
		if err != nil {
			return err
		}
		item.IsLowMemory = false
	case "LowCPU":
		err := dbClient.WriteToTable(s.Dynamo.TableName, "isLowCPU", true)
		if err != nil {
			return err
		}
		item.IsLowCPU = true
	case "OkCPU":
		err := dbClient.WriteToTable(s.Dynamo.TableName, "isLowCPU", false)
		if err != nil {
			return err
		}
		item.IsLowCPU = false
	}

	if item.IsLowCPU && item.IsLowMemory {
		err = dbClient.WriteToTable(s.Dynamo.TableName, "isLowCPU", false)
		if err != nil {
			return err
		}
		err = dbClient.WriteToTable(s.Dynamo.TableName, "isLowMemory", false)
		if err != nil {
			return err
		}
		err = s.execASGPolicy()
		if err != nil {
			return err
		}
	}

	return nil
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

func (s *Sigmund) execASGPolicy() error {
	var client *autoscaling.Client
	asg, err := autoscaling.New(s.Autoscaling.AutoScalingGroupName, s.Autoscaling.PolicyName, s.Autoscaling.Region)
	if err != nil {
		return err
	}

	client, err = asg.NewClient()
	if err != nil {
		return err
	}

	err = client.TriggerPolicy(asg)
	if err != nil {
		return err
	}

	return nil
}
