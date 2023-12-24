package types

type BitcoinChartResponse struct {
	Prices      [][]float64 `json:"prices"`
	TotalVolume [][]float64 `json:"total_volumes"`
}
