package types

type GiftCard struct {
	Email      string `json:"email" form:"email"`
	Name       string `json:"name" form:"name"`
	Background string `json:"background" form:"background"`
	Service    string `json:"service" form:"service"`
}

type ServiceTime struct {
	Name  string
	Price float64
}
