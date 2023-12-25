package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Paienobe/coingecko-api/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := ":" + os.Getenv("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/bitcoin-charts", controllers.BitcoinChartHandler).Methods("GET")
	router.HandleFunc("/coins-list/{pageNumber}", controllers.CoinsListHandler).Methods("GET")
	router.HandleFunc("/coin/{coinID}", controllers.SingleCoinHandler).Methods("GET")
	router.HandleFunc("/global", controllers.GlobalMarketHandler).Methods("GET")

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(portString, router))
}
