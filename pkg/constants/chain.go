package constants

type Chain string

const (
	ETH   Chain = "eth"
	BSC         = "bsc"
	MATIC       = "matic"
	XDC         = "xinfin"
)

func (c Chain) String() string {
	return string(c)
}
