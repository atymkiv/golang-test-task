package main 

import (
	"net/http" 
	"encoding/json"
	"fmt"
	//"io"
	"log"
	)

type MyObject struct {
	Word string		`json:"string"`
	Number int		`json:"number"`
}

func Receive(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var object MyObject
	//err := json.NewDecoder(io.LimitReader(r.Body, 10)).Decode(&object)
	err := decoder.Decode(&object)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Word is %s, Number is %d\n", object.Word, object.Number)
    
}
func main() {
	http.HandleFunc("/", Receive)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}

}