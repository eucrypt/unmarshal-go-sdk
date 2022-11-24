package constants

type Environment string

const (
	Prod  Environment = "prod"
	Stage             = "stage"
)

var endpointMap = map[Environment]string{
	Prod:  "https://api.prod.unmarshal.com",
	Stage: "https://stg-api.unmarshal.com",
}

func (e Environment) GetEndpoint() string {
	return endpointMap[e]
}
