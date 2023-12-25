package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Paienobe/coingecko-api/constants"
	"github.com/Paienobe/coingecko-api/types"
	"github.com/Paienobe/coingecko-api/utils"
	"github.com/gorilla/mux"
)

var CG_KEY = os.Getenv("COIN_GECKO_KEY")

func BitcoinChartHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	url := fmt.Sprintf(constants.BASE_URL+"/coins/bitcoin/market_chart?vs_currency=usd&days=180&interval=daily&x_cg_demo_api_key=%s", CG_KEY)
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusBadRequest)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Sprintln("Error reading data", err)
	}

	var responseObject types.BitcoinChartResponse
	json.Unmarshal(responseBody, &responseObject)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	v, _ := json.Marshal(responseObject)
	w.Write(v)
}

func CoinsListHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	vars := mux.Vars(r)
	pageNumber := vars["pageNumber"]

	url := fmt.Sprintf(constants.BASE_URL+"/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=50&page=%s&sparkline=true&price_change_percentage=1h%%2C24h%%2C7d&x_cg_demo_api_key=%s", pageNumber, CG_KEY)
	// %% fixes invalid interpolation verb - it shows up as just one % sign

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusBadRequest)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
	}

	var responseObject []types.Coin
	json.Unmarshal(responseBody, &responseObject)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res, _ := json.Marshal(responseObject)
	w.Write(res)
}

func SingleCoinHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	vars := mux.Vars(r)
	coinID := vars["coinID"]

	url := fmt.Sprintf(constants.BASE_URL+"/coins/%s?localization=false&tickers=false&market_data=true&community_data=true&developer_data=false&sparkline=false&x_cg_demo_api_key=%s", coinID, CG_KEY)

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusBadRequest)
	}

	responseBody, err := io.ReadAll(response.Body)

	var responseObject types.SingleCoin
	json.Unmarshal(responseBody, &responseObject)

	res, err := json.Marshal(responseObject)
	if err != nil {
		http.Error(w, "Error Marshalling data", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GlobalMarketHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	url := fmt.Sprintf(constants.BASE_URL+"/global?x_cg_demo_api_key=%s", CG_KEY)
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusBadRequest)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body", err)
	}

	var responseObject types.GlobalMarketData
	json.Unmarshal(responseBody, &responseObject)

	res, err := json.Marshal(responseObject)
	if err != nil {
		http.Error(w, "Error marshalling data", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
