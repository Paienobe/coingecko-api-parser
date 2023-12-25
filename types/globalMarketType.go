package types

type GlobalMarketData struct {
	Data AllMarketData `json:"data"`
}

type AllMarketData struct {
	ActiveCurrencies       int                  `json:"active_cryptocurrencies"`
	Exchanges              int                  `json:"markets"`
	TotalMarketCap         SupportedCurrencies  `json:"total_market_cap"`
	MarketCapPercentages   MarketCapPercentages `json:"market_cap_percentage"`
	MarketCapPercentage24h float64              `json:"market_cap_change_percentage_24h_usd"`
}

type MarketCapPercentages struct {
	Btc float64 `json:"btc"`
	Eth float64 `json:"eth"`
}
