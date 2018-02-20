package testPost

import (
	"fmt"
	"net/http"
)

//Test1 yo dawg
func Test1(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("hello there sailor")
}
