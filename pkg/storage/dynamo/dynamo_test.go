package dynamo

import (
	"testing"
)

type config struct {
	table    string
	region   string
	key      string
	expected error
}

var configs = []config{
	{
		table:    "test1",
		region:   "us-east-1",
		key:      "testkey1",
		expected: nil,
	},
	{
		table:    "test2",
		region:   "us-east-1",
		key:      "testkey2",
		expected: nil,
	},
	{
		table:    "test3",
		region:   "eu-west-1",
		key:      "testkey3",
		expected: nil,
	},
	{
		table:    "test4",
		region:   "us-east-1",
		key:      "testkey4",
		expected: nil,
	},
}

// Test Constructor function
func TestNew(t *testing.T) {
	for _, conf := range configs {
		_, err := checkConfig(conf.table, conf.region, conf.key)
		if err != nil {
			t.Errorf("Error message was incorrect, expected: %v, got: %v", conf.expected, err)
		}
	}
}
