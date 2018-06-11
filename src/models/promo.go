package models

type Buy struct {
	Category string `json:"category"`
	Count    int `json:"count"`
}

type Discount struct {
	Percentage float32 `json:"percentage"`
}

type Fixed struct {
	Price float32 `json:"price"`
}

// Either discount or fixed
type Off struct {
	Discount *Discount `json:"discount"`
	Fixed    *Fixed `json:"fixed"`
}

// Either "all" or specify a count
type Get struct {
	Category string `json:"category"`
	All      bool `json:"all"`
	Count    int `json:"count"`
	Off      Off `json:"off"`
}

type Promo struct {
	Id   string `json:"id"`
  Description string `json:"description"`
	Buys []Buy `json:"buys"`
	Gets []Get `json:"gets"`
}
