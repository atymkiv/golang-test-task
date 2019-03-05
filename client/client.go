package main

import (
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"strconv"
)

func MakeRand() io.Reader{
	var object MyObject

	var arr []byte
	for i := 0; i < rand.Intn(50); i++ {
		arr = append(arr, byte(rand.Intn(122-65)+65))
	}
	object.Word = string(arr)
	object.Number = rand.Intn(100000)
	oJson, _ := json.Marshal(object)
	r := bytes.NewReader(oJson)
	return r
}

type MyObject struct {
	Word string		`json:"string"`
	Number int		`json:"number"`
}

func Post() {
	r := MakeRand()
	resp, err := http.Post("http://localhost:9000/", "application/json", r)
	fmt.Printf("%v %v\n", err, resp)
}

func main() {
	length, _ := strconv.Atoi(os.Args[1])
	for i := 0; i<length; i++{
        go Post()
    }
    var input string
    fmt.Scanln(&input)

}
