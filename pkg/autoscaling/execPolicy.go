package autoscaling

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

// TriggerPolicy is a method called by the Autoscaling Client
func (client *Client) TriggerPolicy(asg *Autoscaling) error {
	input := &autoscaling.ExecutePolicyInput{
		AutoScalingGroupName: aws.String(asg.AutoScalingGroupName),
		HonorCooldown:        aws.Bool(true),
		PolicyName:           aws.String(asg.PolicyName),
	}

	_, err := client.ExecutePolicy(input)
	awsErr := checkAutoscalingError(err)

	return awsErr
}

func checkAutoscalingError(err error) error {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case autoscaling.ErrCodeScalingActivityInProgressFault:
				return fmt.Errorf("%v: %v", autoscaling.ErrCodeScalingActivityInProgressFault, aerr.Error())
			case autoscaling.ErrCodeResourceContentionFault:
				return fmt.Errorf("%v: %v", autoscaling.ErrCodeResourceContentionFault, aerr.Error())
			default:
				return fmt.Errorf("%v", aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.//^
			//I would probably rewrite this to make the default be to return error and handle the specific cases you're interested in here.
			//Wrapping a switch in an if is tough to read.
			return err
		}
	}
	return nil
}
