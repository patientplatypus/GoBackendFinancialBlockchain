package main

import (
	"fmt"
	"net/http"

	"backendBLOK/routes/getChainCode"
	"backendBLOK/routes/getStockData"
	"backendBLOK/routes/postStockData"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"backendBLOK/routes/testPost"
)

// our main function
func main() {
	router := mux.NewRouter()
	//   mux :=
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     w.Header().Set("Content-Type", "application/json")
	//     w.Write([]byte("{\"hello\": \"world\"}"))
	// })

	// router.HandleFunc("/testPost", testPost.Test1).Methods("POST")
	router.HandleFunc("/testPost", testPost.Test1).Methods("GET")
	router.HandleFunc("/stock/{symbol}", getStockData.Symbol).Methods("GET")
	router.HandleFunc("/order", postStockData.Order).Methods("POST")
	router.HandleFunc("/querychain/{value}", getChainCode.Query).Methods("GET")
	router.HandleFunc("/exchangetochain", postStockData.ExchangeToChain).Methods("POST")
	router.HandleFunc("/moneyinit", postStockData.MoneyINIT).Methods("GET")

	// router.HandleFunc("/people", GetPeople).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	// func GetPeople(w http.ResponseWriter, r *http.Request) {}
	// func GetPerson(w http.ResponseWriter, r *http.Request) {}
	// func CreatePerson(w http.ResponseWriter, r *http.Request) {}
	// func DeletePerson(w http.ResponseWriter, r *http.Request) {}
	fmt.Printf("Now listening on port 8000")
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8000", handler)
}
