package getStockData

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Symbol the stock symbol we are going to retrieve from https://www.alphavantage.co
// OAO5GXY4IMRLFLII
func Symbol(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	fmt.Printf("inside getStockData.Symbol")

	var apikey = "OAO5GXY4IMRLFLII"

	var url = "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=" + params["symbol"] + "&interval=1min&apikey=" + apikey

	fmt.Printf("value of url: %s", url)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)

	fmt.Printf(string(responseData))

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(json.Unmarshal(responseData)))
	// json.NewEncoder(w).Encode()
	w.Write(responseData)
}
