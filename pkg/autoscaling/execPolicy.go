package autoscaling

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

// TriggerPolicy is a method called by the Autoscaling Client
func (client *Client) TriggerPolicy(asg *Autoscaling) error {
	input := &autoscaling.ExecutePolicyInput{
		AutoScalingGroupName: aws.String(asg.AutoScalingGroupName),
		HonorCooldown:        aws.Bool(true),
		PolicyName:           aws.String(asg.PolicyName),
	}

	result, err := client.ExecutePolicy(input)
	awsErr := checkAutoscalingError(err)

	fmt.Println(result)
	return awsErr
}
