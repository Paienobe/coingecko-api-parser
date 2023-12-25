package types

type SingleCoin struct {
	ID          string            `json:"id"`
	Symbol      string            `json:"symbol"`
	Name        string            `json:"name"`
	Description CoinDescription   `json:"description"`
	Links       CoinLinks         `json:"links"`
	Image       map[string]string `json:"image"`
	MarketData  MarketData        `json:"market_data"`
}

type CoinDescription struct {
	En string `json:"en"`
}

type CoinLinks struct {
	Homepage []string `json:"homepage"`
}

type MarketData struct {
	CurrentPrice            SupportedCurrencies `json:"current_price"`
	Ath                     SupportedCurrencies `json:"ath"`
	AthDate                 PriceDates          `json:"ath_date"`
	Atl                     SupportedCurrencies `json:"atl"`
	AtlDate                 PriceDates          `json:"atl_date"`
	MarketCap               SupportedCurrencies `json:"market_cap"`
	TotalVolume             SupportedCurrencies `json:"total_volume"`
	CirculatingSupply       float64             `json:"circulating_supply"`
	TotalSupply             float64             `json:"total_supply"`
	PriceChangePercentage24 SupportedCurrencies `json:"price_change_percentage_24h_in_currency"`
	FullyDilutedValuation   SupportedCurrencies `json:"fully_diluted_valuation"`
}

type SupportedCurrencies struct {
	Usd float64 `json:"usd"`
	Gbp float64 `json:"gbp"`
	Eur float64 `json:"eur"`
	Ngn float64 `json:"ngn"`
}

type PriceDates struct {
	Usd string `json:"usd"`
	Gbp string `json:"gbp"`
	Eur string `json:"eur"`
	Ngn string `json:"ngn"`
}
