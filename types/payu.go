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

// POS ID (pos_id):	475696
// Second key (MD5):	02884264a9501e9c7daae8bb27a58bf0
// OAuth protocol - client_id :	475696
// OAuth protocol - client_secret:	075968a77327f6b62a3b03b3d1c78529
