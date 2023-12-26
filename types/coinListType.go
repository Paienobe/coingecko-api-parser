package types

type Coin struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	Image            string    `json:"image"`
	CurrentPrice     float64   `json:"current_price"`
	Rank             int       `json:"market_cap_rank"`
	Volume           int       `json:"total_volume"`
	MarketCap        int       `json:"market_cap"`
	PriceChange1h    float64   `json:"price_change_percentage_1h_in_currency"`
	PriceChange24h   float64   `json:"price_change_percentage_24h_in_currency"`
	PriceChange7d    float64   `json:"price_change_percentage_7d_in_currency"`
	SparkLine7d      SparkLine `json:"sparkline_in_7d"`
	CiculatingSupply float64   `json:"circulating_supply"`
	TotalSupply      float64   `json:"total_supply"`
}

type SparkLine struct {
	Price []float64 `json:"price"`
}
