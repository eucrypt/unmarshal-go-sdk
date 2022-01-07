package types

type NFTAsset struct {
	AssetContract      string `json:"asset_contract"`
	TokenId            string `json:"token_id"`
	Owner              string `json:"owner"`
	ExternalLink       string `json:"external_link"`
	Type               string `json:"type"`
	Balance            int    `json:"balance"`
	IssuerSpecificData struct {
		EntireResponse string `json:"entire_response"`
		ImageUrl       string `json:"image_url"`
		Name           string `json:"name"`
	} `json:"issuer_specific_data"`
	Price        string `json:"price"`
	AnimationUrl string `json:"animation_url"`
	Description  string `json:"description"`
	NftMetadata  []struct {
		TraitType   string      `json:"trait_type"`
		Value       interface{} `json:"value"`
		DisplayType string      `json:"display_type,omitempty"`
	} `json:"nft_metadata,omitempty"`
}

type NFTAssetsResp []NFTAsset

type NFTTxns struct {
	ContractAddress string `json:"contract_address"`
	TokenId         string `json:"token_id"`
	TransactionHash string `json:"transaction_hash"`
	BlockNumber     int    `json:"block_number"`
	BlockHash       string `json:"block_hash"`
	TransactionIdx  int    `json:"transaction_idx"`
	LogIdx          int    `json:"log_idx"`
	Sender          string `json:"sender"`
	To              string `json:"to"`
}

type NFTTxnsResp []NFTTxns

type NFTByTokenIDResp struct {
	ContractAddress string `json:"contract_address"`
	TokenId         string `json:"token_id"`
	TokenUri        string `json:"token_uri"`
	ActualOwner     string `json:"actual_owner"`
	Properties      []struct {
		TraitType   string      `json:"trait_type"`
		Value       interface{} `json:"value"`
		DisplayType string      `json:"display_type"`
	} `json:"properties,omitempty"`
	ImageUrl       string `json:"image_url,omitempty"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Price          string `json:"price"`
	AnimationUrl   string `json:"animation_url"`
	EntireResponse string `json:"entire_response"`
	Description    string `json:"description"`
}

//NFTHolderResponse the response is an array of addresses
type NFTHolderResponse []string
