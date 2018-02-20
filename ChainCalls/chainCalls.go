package chaincalls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type MakeaNewObj struct {
	Returncode    string `json:"returnCode"`
	Info          string `json:"info"`
	Transactionid string `json:"transactionID"`
	Arg0          string `json:"arg0"`
	Arg1          string `json:"arg1"`
	Arg2          string `json:"arg2"`
	Arg3          string `json:"arg3"`
	Arg4          string `json:"arg4"`
	Arg5          string `json:"arg5"`
	Arg6          string `json:"arg6"`
	Arg7          string `json:"arg7"`
	Arg8          string `json:"arg8"`
}

type ModifyReturn struct {
	Returncode    string `json:"returnCode"`
	Info          string `json:"info"`
	Transactionid string `json:"transactionID"`
}

type FindReturnCheck struct {
	Returncode   string `json:"returnCode"`
	ResultString string `json:"result"`
	Info         string `json:"info"`
}

type FindReturn struct {
	Returncode string          `json:"returnCode"`
	Result     json.RawMessage `json:"result"`
	Info       string          `json:"info"`
}

type Result struct {
	Txid  string `json:"TxId"`
	Value struct {
		Doctype      string `json:"docType"`
		Documentid   string `json:"documentID"`
		Folderid     string `json:"folderID"`
		Action       string `json:"action"`
		User         string `json:"user"`
		Orgname      string `json:"orgName"`
		Datetime     int64  `json:"dateTime"`
		Message      string `json:"message"`
		Documenthash string `json:"documentHash"`
		Product      string `json:"product"`
		Docextension string `json:"docExtension"`
	} `json:"Value"`
	Isdelete string `json:"IsDelete"`
}

//NewOrMod looks for the entry and makes it if it doesnt exist
func NewOrMod(arg []string) []byte {
	entryexists, _ := QueryExistance(arg[0])
	if entryexists == true {
		return ModifyEntry(arg)
	} else if entryexists == false {
		return NewEntry(arg)
	}
	var dummyreturn []byte
	return dummyreturn
}

//ParserMarshaller ~ credit where credit is due - came from #go-nutsUSERd9k
func ParserMarshaller(httpreturn []byte) (string, []byte) {
	// fmt.Println("13")
	r1 := new(FindReturn)
	// fmt.Println("14")
	r2 := make([]Result, 0)
	// fmt.Println("15")
	if err := json.Unmarshal(httpreturn, &r1); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("16")
	x, err := strconv.Unquote(string(r1.Result))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("17")
	if err := json.Unmarshal([]byte(x), &r2); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("18")
	oneR, _ := json.Marshal(r2[len(r2)-1])
	return x, oneR
}

//QueryExistance checks to see if the entry entry exists
func QueryExistance(firstArg string) (bool, []byte) {
	// fmt.Println("1")
	// fmt.Println("inside QueryExistance")
	// fmt.Println("2")
	// fmt.Println("inside QueryExistance and value of firstArg")
	// fmt.Println("3")
	// fmt.Println(firstArg)
	// fmt.Println("4")
	var urlFind = "http://129.146.109.15:4000/bcsgw/rest/v1/transaction/query"

	var jsonStr = []byte(`{
                "channel":"platypusorgan1orderer",
                "chaincode":"file-trace",
                "chaincodeVer":"1.0",
                "method":"getDocumentHistory",
                "args":["` + firstArg + `"]
                }`)

	fmt.Printf("value of json before send")
	fmt.Printf(string(jsonStr))

	responseFind, _ := http.NewRequest("POST", urlFind, bytes.NewBuffer(jsonStr))
	responseFind.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	respFind, err := client.Do(responseFind)
	if err != nil {
		panic(err)
	}
	defer respFind.Body.Close()
	bodyFind, _ := ioutil.ReadAll(respFind.Body)
	// fmt.Println("response from trying to find url:", string(bodyFind))

	// fmt.Println("5")
	var findobjcheck FindReturnCheck
	var findobj FindReturn
	// fmt.Println("6")
	json.Unmarshal(bodyFind, &findobjcheck)
	json.Unmarshal(bodyFind, &findobj)
	// fmt.Println("7")
	var dummybyte []byte
	// fmt.Println("8")
	// fmt.Println("value of findobj")
	// fmt.Println(findobj)
	// fmt.Println(findobjcheck)

	if findobjcheck.ResultString == "[]" {
		fmt.Println("9")
		return false, dummybyte
	} else {
		fmt.Println("11")
		_, parsebyte := ParserMarshaller(bodyFind)
		fmt.Println("12")
		return true, parsebyte
	}

	// parsestring, parsebyte := ParserMarshaller(bodyFind)
	//
	// if parsestring == "[]" {
	// 	return false, parsebyte
	// }
	// return true, parsebyte
}

//NewEntry creates a new entry and returns it
func NewEntry(arg []string) []byte {
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	var urlNew = "http://129.146.109.15:4000/bcsgw/rest/v1/transaction/invocation"
	var jsonStrNew = []byte(`{
                "channel":"platypusorgan1orderer",
                "chaincode":"file-trace",
                "chaincodeVer":"1.0",
                "method":"newDocument",
                "args":["` + arg[0] + `","` + arg[1] + `","` + arg[2] + `","` + arg[3] + `","` + strconv.FormatInt(millis, 10) + `", "` + arg[5] + `", "` + arg[6] + `", "` + arg[7] + `", "` + arg[8] + `"]}`)

	fmt.Printf("value of jsonStrNew before send")
	fmt.Printf(string(jsonStrNew))

	responseNew, _ := http.NewRequest("POST", urlNew, bytes.NewBuffer(jsonStrNew))
	responseNew.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	respNew, err := client.Do(responseNew)
	if err != nil {
		panic(err)
	}
	defer respNew.Body.Close()

	bodyNew, _ := ioutil.ReadAll(respNew.Body)

	var newobj MakeaNewObj
	json.Unmarshal(bodyNew, &newobj)

	if newobj.Returncode == "Success" {
		newobj.Arg0 = arg[0]
		newobj.Arg1 = arg[1]
		newobj.Arg2 = arg[2]
		newobj.Arg3 = arg[3]
		newobj.Arg4 = arg[4]
		newobj.Arg5 = arg[5]
		newobj.Arg6 = arg[6]
		newobj.Arg7 = arg[7]
		newobj.Arg8 = arg[8]
		marshallednewobj, _ := json.Marshal(newobj)
		return marshallednewobj
	}
	var dummybyte []byte
	return dummybyte
	// var newified ModifyReturn
	// marshallednewified, _ := json.Marshal(newified)
	// json.Unmarshal(bodyNew, &newified)
	// if newified.Returncode == "Success" {
	// 	// newified.Arg0 = arg[0]
	// 	// newified.Arg1 = arg[1]
	// 	// newified.Arg2 = arg[2]
	// 	// newified.Arg3 = arg[3]
	// 	// newified.Arg4 = arg[4]
	// 	// newified.Arg5 = arg[5]
	// 	// newified.Arg6 = arg[6]
	// 	// newified.Arg7 = arg[7]
	// 	// newified.Arg8 = arg[8]
	//
	// 	return marshallednewified
	// }

	// fmt.Println("right before bodyModify supposed to be shown")
	// var newified ModifyNewReturn
	// json.Unmarshal(bodyNew, &newified)
	// fmt.Println(newified.Returncode)
	// if newified.Returncode == "Success" {
	// 	fmt.Println("inside newified Returncode success")
	// 	fmt.Println("inside newified and value of arg[0]")
	// 	fmt.Println(arg[0])
	// 	_, parsebyte := QueryExistance(arg[0])
	// 	return parsebyte
	// }
	//
	// var dummybyte []byte
	// return dummybyte

	//
	//
	// var newTarget FindReturn
	//
	// json.NewDecoder(respNew.Body).Decode(newTarget)
	//
	// fmt.Println("response from making a new document:", string(bodyNew))
	// _, parsebyte := ParserMarshaller(bodyNew)
	// return parsebyte

}

//ModifyEntry modifies the entry provided it exists
func ModifyEntry(arg []string) []byte {
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	var urlModify = "http://129.146.109.15:4000/bcsgw/rest/v1/transaction/invocation"
	var jsonStrModify = []byte(`{
                "channel":"platypusorgan1orderer",
                "chaincode":"file-trace",
                "chaincodeVer":"1.0",
                "method":"modifyDocument",
                "args":["` + arg[0] + `","` + arg[1] + `","` + arg[2] + `","` + arg[3] + `","` + strconv.FormatInt(millis, 10) + `", "` + arg[5] + `", "` + arg[6] + `", "` + arg[7] + `", "` + arg[8] + `"]}`)

	fmt.Printf("value of jsonStrModify\n")
	fmt.Printf(string(jsonStrModify))

	responseModify, _ := http.NewRequest("POST", urlModify, bytes.NewBuffer(jsonStrModify))
	responseModify.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	respModify, err := client.Do(responseModify)
	if err != nil {
		panic(err)
	}
	defer respModify.Body.Close()
	bodyModify, _ := ioutil.ReadAll(respModify.Body)
	fmt.Println("right before bodyModify supposed to be shown")
	var modified ModifyReturn
	json.Unmarshal(bodyModify, &modified)
	fmt.Println(modified.Returncode)
	if modified.Returncode == "Success" {
		_, parsebyte := QueryExistance(arg[0])
		return parsebyte
	}

	var dummybyte []byte
	return dummybyte
}
