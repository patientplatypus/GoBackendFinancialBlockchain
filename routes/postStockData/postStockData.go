package postStockData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"backendBLOK/chaincalls"
	"backendBLOK/globals"
)

// {"payload":{"payload":{"ordertype":"sell","sharenumber":"4","shareprice":"4","stocksymbol":"S"}}}

//Payload is payload
type Payload struct {
	Ordertype   string  `json:"ordertype"`
	Sharenumber int     `json:"sharenumber"`
	Shareprice  float64 `json:"shareprice"`
	Stocksymbol string  `json:"stocksymbol"`
}

//ColorGroup is the color group
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

//LastPrice is blah
type LastPrice struct {
	Symbol    string  `json:"symbol"`
	Timestamp string  `json:"timestamp"`
	Lastprice float64 `json:"lastprice"`
}

//FindReturn is find return
type FindReturn struct {
	Returncode string `json:"returnCode"`
	Result     string `json:"result"`
	Info       string `json:"info"`
}

type JSONString string

func (j JSONString) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}

//MoneyINIT initializes the traders money
func MoneyINIT(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("inside MoneyINIT in postStockData")
	arg := []string{"TRADER_$$$", "1000", "placeholder2", "placeholder3", "placeholder4", "placeholder5", "placeholder6", "placeholder7", "placeholder8"}
	// chaincalls.EntryPoint(arg)
	bytereturn := chaincalls.NewOrMod(arg)
	w.Write(bytereturn)
}

//ExchangeToChain takes the exchange values and writes them to the chain
func ExchangeToChain(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	var price LastPrice
	json.Unmarshal(b, &price)
	var firstArg = "EXCHANGE_" + price.Symbol

	//assign price and symbol to global temp variables
	globals.GlobalStockPrice = price.Lastprice
	globals.GlobalStockSymbol = string(price.Symbol)

	arg := []string{firstArg, fmt.Sprint(price.Lastprice), "placeholder2", "placeholder3", "placeholder4", "placeholder5", "placeholder6", "placeholder7", "placeholder8"}
	bytereturn := chaincalls.NewOrMod(arg)
	w.Write(bytereturn)
}

//Order posting the order to buy some stock.
func Order(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("inside postStockData.Order b, %s", b)
	// var order PayloadStruct
	var order Payload
	err := json.Unmarshal(b, &order)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	fmt.Println("from the payload we have")
	fmt.Printf("value of the global variable: ")
	printglobals := fmt.Sprint(globals.GlobalStockPrice)
	fmt.Printf(printglobals)
	fmt.Println(order.Sharenumber)
	json.NewEncoder(w).Encode(order)

	// value, err := strconv.ParseFloat(globals.GlobalStockPrice, 32)
	// fmt.Println("value of value from GlobalStockPrice")
	// fmt.Println(value)
	// if err != nil {
	// 	fmt.Println("oh noes!")
	// }
	// GlobalStockPrice := float32(value)

	fmt.Println("value of GlobalStockPrice")
	fmt.Println(printglobals)

	arg := []string{"TRADER_ORDER", order.Ordertype, fmt.Sprint(order.Sharenumber), fmt.Sprint(order.Shareprice), fmt.Sprint(order.Stocksymbol), "placeholder5", "placeholder6", "placeholder7", "placeholder8"}

	if order.Ordertype == "sell" {
		if globals.GlobalStockPrice >= order.Shareprice {
			bytereturn := chaincalls.NewOrMod(arg)
			w.Write(bytereturn)
		} else {
			s := `{"error":"SELL PRICE not LOW ENOUGH, not written to chain", "Value": "nada"}`
			content, _ := json.Marshal(JSONString(s))
			w.Write(content)
		}
	} else if order.Ordertype == "buy" {
		if globals.GlobalStockPrice <= order.Shareprice {
			bytereturn := chaincalls.NewOrMod(arg)
			w.Write(bytereturn)
		} else {
			s := `{"error":"BUY PRICE not HIGH ENOUGH, not written to chain", "Value": "nada"}`
			content, _ := json.Marshal(JSONString(s))
			w.Write(content)
		}
	}
}

// {"returnCode":"Success","result":"[]","info":null}

// FIX THIS TOMORROW BY CALLING CHAINCALLS!!!

//ExchangeToChain sends the auto updating exchange data to save on the block chain

// func ExchangeToChain(w http.ResponseWriter, r *http.Request) {
// 	b, _ := ioutil.ReadAll(r.Body)
// 	fmt.Printf("inside postStockData.Order b, %s", b)
// 	var price LastPrice
// 	err := json.Unmarshal(b, &price)
// 	if err != nil {
// 		fmt.Println("oh no there was an error!")
// 	}
// 	fmt.Printf("value of symbol ")
// 	fmt.Printf("%v\n", price.Symbol)
// 	fmt.Printf("value of timestamp ")
// 	fmt.Printf("%v\n", price.Timestamp)
// 	fmt.Printf("value of lastprice ")
// 	fmt.Printf("%v\n", fmt.Sprint(price.Lastprice))
//
// 	// floatLastPrice, _ := strconv.ParseFloat(price.Lastprice, 32)
// 	globals.GlobalStockPrice = price.Lastprice
//
// 	globals.GlobalStockSymbol = string(price.Symbol)
//
// 	var firstArg = "EXCHANGE_" + price.Symbol
//
// 	// var args string
// 	// args = `[` + firstArg + `,` + price.Symbol + `,` + string(price.Timestamp) + `,"placeholder3", "placeholder4", "placeholder5", "placeholder6", "placeholder7", "placeholder8"]`
//
// 	var urlFind = "http://129.146.109.15:4000/bcsgw/rest/v1/transaction/query"
//
// 	var jsonStr = []byte(`{
//                 "channel":"platypusorgan1orderer",
//                 "chaincode":"file-trace",
//                 "chaincodeVer":"1.0",
//                 "method":"getDocumentHistory",
//                 "args":["` + firstArg + `"]
//                 }`)
//
// 	fmt.Printf("value of json before send")
// 	fmt.Printf(string(jsonStr))
//
// 	responseFind, _ := http.NewRequest("POST", urlFind, bytes.NewBuffer(jsonStr))
// 	responseFind.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	respFind, err := client.Do(responseFind)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer respFind.Body.Close()
// 	bodyFind, _ := ioutil.ReadAll(respFind.Body)
// 	fmt.Println("response from trying to find url:", string(bodyFind))
//
// 	var findreturn FindReturn
// 	err2 := json.Unmarshal(bodyFind, &findreturn)
// 	if err2 != nil {
// 		fmt.Println("oh no there was an error!")
// 	}
//
// 	// var watch = string(time.Now().UTC().Unix())
// 	// var watchstring = string(watch)
//
// 	now := time.Now()
// 	nanos := now.UnixNano()
//
// 	millis := nanos / 1000000
//
// 	if findreturn.Result == "[]" {
// 		var urlNew = "http://129.146.109.15:4000/bcsgw/rest/v1/transaction/invocation"
// 		var jsonStrNew = []byte(`{
// 									"channel":"platypusorgan1orderer",
// 									"chaincode":"file-trace",
// 									"chaincodeVer":"1.0",
// 									"method":"newDocument",
// 									"args":["` + firstArg + `","` + price.Symbol + `","placeholder2","placeholder3","` + strconv.FormatInt(millis, 10) + `", "placeholder5", "placeholder6", "placeholder7", "placeholder8"]
// 									}`)
//
// 		fmt.Printf("value of jsonStrNew before send")
// 		fmt.Printf(string(jsonStrNew))
//
// 		responseNew, _ := http.NewRequest("POST", urlNew, bytes.NewBuffer(jsonStrNew))
// 		responseNew.Header.Set("Content-Type", "application/json")
// 		client := &http.Client{}
// 		respNew, err := client.Do(responseNew)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer respNew.Body.Close()
// 		bodyNew, _ := ioutil.ReadAll(respNew.Body)
// 		fmt.Println("response from making a new document:", string(bodyNew))
//
// 	} else {
// 		var urlModify = "http://129.146.109.15:4000/bcsgw/rest/v1/transaction/invocation"
// 		var jsonStrModify = []byte(`{
// 									"channel":"platypusorgan1orderer",
// 									"chaincode":"file-trace",
// 									"chaincodeVer":"1.0",
// 									"method":"modifyDocument",
// 									"args":["` + firstArg + `","` + price.Symbol + `","placeholder2","placeholder3","` + strconv.FormatInt(millis, 10) + `", "placeholder5", "placeholder6", "placeholder7", "placeholder8"]
// 									}`)
//
// 		fmt.Printf("value of jsonStrModify\n")
// 		fmt.Printf(string(jsonStrModify))
//
// 		responseModify, _ := http.NewRequest("POST", urlModify, bytes.NewBuffer(jsonStrModify))
// 		responseModify.Header.Set("Content-Type", "application/json")
// 		client := &http.Client{}
// 		respModify, err := client.Do(responseModify)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer respModify.Body.Close()
// 		bodyModify, _ := ioutil.ReadAll(respModify.Body)
// 		fmt.Println("response from modifying an existing document:", string(bodyModify))
// 	}
//
// }
