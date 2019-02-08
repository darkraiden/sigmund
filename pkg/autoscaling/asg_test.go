package autoscaling

import "testing"

type config struct {
	autoscalingGroupName string
	policyName           string
	region               string
	expected             error
}

var configs = []config{
	{
		autoscalingGroupName: "test1",
		policyName:           "policy1",
		region:               "us-east-1",
		expected:             nil,
	},
	{
		autoscalingGroupName: "test2",
		policyName:           "policy2",
		region:               "us-east-1",
		expected:             nil,
	},
	{
		autoscalingGroupName: "test3",
		policyName:           "policy3",
		region:               "us-east-1",
		expected:             nil,
	},
}

func TestNew(t *testing.T) {
	for _, conf := range configs {
		_, err := checkConfig(conf.autoscalingGroupName, conf.policyName, conf.region)
		if err != nil {
			t.Errorf("Error message was incorrect, expected: %v, got: %v", conf.expected, err)
		}
	}
}
