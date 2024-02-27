package types

type PayUPostOrderBody struct {
	continueUrl   string    `json:"continueUrl"`
	notifyUrl     string    `json:"notifyUrl"`
	customerIp    string    `json:"customerIp"`
	merchantPosId string    `json:"merchantPosId"`
	description   string    `json:"description"`
	currencyCode  string    `json:"currencyCode"`
	totalAmount   string    `json:"totalAmount"`
	products      []Product `json:"products"`
}

type PayUAuth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	GrantType   string `json:"grant_type"`
}

type Product struct {
	Name      string
	UnitPrice string
	quantity  string
}
