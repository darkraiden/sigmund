package dynamo

// Item represents the structure of the JSON body coming back from DynamoDB after
// a Select query is executed
type Item struct {
	ID          int  `json:"ID"`
	IsLowMemory bool `json:"isLowMemory"`
	IsLowCPU    bool `json:"isLowCPU"`
}
