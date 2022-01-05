package constants

type Environment string

const (
	Prod  Environment = "prod"
	Stage             = "stage"
)

var endpointMap = map[Environment]string{
	Prod:  "https://api.unmarshal.io/",
	Stage: "https://stg-api.unmarshal.io/",
}

func (e Environment) GetEndpoint() string {
	return endpointMap[e]
}
