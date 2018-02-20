package getChainCode

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Query the chain code
func Query(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	fmt.Printf("inside getChainCode.Query")

	fmt.Println("value of params[value]")
	fmt.Println(params["value"])

	// "args": ` + `[` + params["value"] + `]` + `

	//WORKS!
	// curl -d '{"channel":"platypusorgan1orderer", "chaincode":"file-trace", "chaincodeVer": "1.0", "method":"newDocument", "args":[0,1,2,3,4,5,6,7,8]}' -H "Content-Type: application/json" -X POST http://129.146.109.15:4000/bcsgw/rest/v1/transaction/invocation

	// curl -d '{"channel":"platypusorgan1orderer", "chaincode":"file-trace", "chaincodeVer": "1.0", "method":"getDocumentHistory", "args":["1"]}' -H "Content-Type: application/json" -X POST http://129.146.109.15:4000/bcsgw/rest/v1/transaction/invocation

	// curl -d '{"channel":"platypusorgan1orderer", "chaincode":"file-trace", "chaincodeVer": "1.0", "method":"getDocumentHistory", "args":["1"]}' -H "Content-Type: application/json" -X POST http://129.146.109.15:4000/bcsgw/rest/v1/transaction/query

	// curl -d '{"channel":"platypusorgan1orderer", "chaincode":"file-trace", "chaincodeVer": "1.0", "method":"viewDocument", "args":["0"]}' -H "Content-Type: application/json" -X POST http://129.146.109.15:4000/bcsgw/rest/v1/transaction/invocation

	//curl -d '{"channel":"platypusorgan1orderer", "chaincode":"file-trace", "chaincodeVer": "1.0", "method":"newDocument", "args":["0", "1", "2", "3", "4", "5", "6", "7", "8"]}' -H "Content-Type: application/json" -X POST http://129.146.109.15:4000/bcsgw/rest/v1/transaction/invocation

	var url = "http://129.146.109.15:4000/bcsgw/rest/v1/transaction/invocation"

	var jsonStr = []byte(`{
                "channel":"platypusorgan1orderer",
                "chaincode":"file-trace",
                "chaincodeVer":"1.0",
                "method":"modifyDocument",
                "args":["1", "2", "3", "7", "5", "6", "7", "8", "9"]
                }`)
	response, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	response.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	//   axios.post('http://129.146.106.151:4002/bcsgw/rest/v1/transaction/invocation', {
	//           "channel": "doctorpharmacist",
	//           "chaincode": "file-trace",
	//           "chaincodeVer": "v1",
	//           "method": "newDocument",
	//           "args": [RXID, patientID, FirstName, LastName, Timestamp, Doctor, Prescription, Refills, Status]
	// })

	// url := "http://restapi3.apiary.io/notes"
	// fmt.Println("URL:>", url)

	// var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	// req.Header.Set("Content-Type", "application/json")
	//
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	//
	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	// var apikey = "OAO5GXY4IMRLFLII"
	//
	// var url = "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=" + params["symbol"] + "&interval=1min&apikey=" + apikey
	//
	// fmt.Printf("value of url: %s", url)
	//
	// response, err := http.Get(url)
	//
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }
	//
	// responseData, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println(string(json.Unmarshal(responseData)))
	// // json.NewEncoder(w).Encode()
	// w.Write(responseData)
}
