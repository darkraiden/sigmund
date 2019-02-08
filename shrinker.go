package sigmund

import (
	"github.com/darkraiden/sigmund/pkg/autoscaling"
	"github.com/darkraiden/sigmund/pkg/storage/dynamo"
)

// Shrink is the core function of the package
// which Executes an Autoscaling Group Policy
// when requirements are met
func (s *Sigmund) Shrink() error {
	var item *DBItem
	dbClient, err := s.newDBClient()
	if err != nil {
		return err
	}

	// Run a Select query
	item, err = s.readDynamo(dbClient)
	if err != nil {
		return err
	}

	switch s.Dynamo.Key {
	case "isLowCPU":
		if !item.IsLowCPU {
			err = dbClient.WriteToTable(s.Dynamo.TableName, s.Dynamo.Key, true)
			if err != nil {
				return err
			}
			item.IsLowCPU = true
		}
	case "isLowMemory":
		if !item.IsLowMemory {
			err = dbClient.WriteToTable(s.Dynamo.TableName, s.Dynamo.Key, true)
			if err != nil {
				return err
			}
			item.IsLowMemory = true
		}
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

func (s *Sigmund) newDBClient() (*dynamo.Client, error) {
	dynamo, err := dynamo.New(s.Dynamo.TableName, s.Dynamo.Region, s.Dynamo.Key)
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
