package autoscaling

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

// Autoscaling represents the structure of a series
// of AWS autoscaling group parameters to trigger
// a dimension change
type Autoscaling struct {
	AutoScalingGroupName string
	PolicyName           string
	Region               string
}

// Client type is an Autoscaling group client
type Client struct {
	*autoscaling.AutoScaling
}

// New is the Package constructor that initialises
// the Autoscaling config
func New(asgName, policyName, region string) (*Autoscaling, error) {
	config, err := checkConfig(asgName, policyName, region)
	if err != nil {
		return nil, fmt.Errorf("Initialisation error: %v", err)
	}
	return config, nil
}

// NewClient creates an Autoscaling client
func (asg *Autoscaling) NewClient() (*Client, error) {
	var svc Client

	sess, err := newSession(asg.Region)
	if err != nil {
		return nil, fmt.Errorf("Cannot open a new AWS session: %v", err)
	}

	svc.AutoScaling = autoscaling.New(sess)

	return &svc, nil
}

func checkConfig(asgName, policyName, region string) (*Autoscaling, error) {
	switch {
	case asgName == "":
		return nil, errors.New("The autoscaling group name cannot be empty")
	case policyName == "":
		return nil, errors.New("The autoscaling policy cannot be empty")
	case region == "":
		return nil, errors.New("Region cannot be empty")
	}
	return &Autoscaling{
		AutoScalingGroupName: asgName,
		PolicyName:           policyName,
		Region:               region,
	}, nil
}

func newSession(region string) (*session.Session, error) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}
