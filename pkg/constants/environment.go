package constants

type Environment string

const (
	Prod     Environment = "prod"
	Stage                = "stage"
	Internal             = "internal" //Internal only for use within Unmarshal. Will not work otherwise.
)

var endpointMap = map[Environment]string{
	Prod:     "https://api.unmarshal.com",
	Stage:    "https://stg-api.unmarshal.com",
	Internal: "https://api.prod.unmarshal.com",
}

func (e Environment) GetEndpoint() string {
	return endpointMap[e]
}
