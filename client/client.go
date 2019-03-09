package main

import (
	"fmt"

	"net/http"
	"math/rand"
	"os"
	"strconv"
)

func MakeRandJSON(max int) <-chan MyObject{
	outChJson := make(chan MyObject, max)
	
	go func() {
		for i := 1; i <= max; i++ {
			var object MyObject
			var arr []byte
			for j := 0; j < rand.Intn(50); j++ { 
				arr = append(arr, byte(rand.Intn(122-65)+65)) //generating random words of random chars and length
			}

			object.Word = string(arr)
			object.Number = rand.Intn(100000)

			outChJson <- object
		}

		close(outChJson)
	}()

	return outChJson
}
	
	//oJson, _ := json.Marshal(object)
	//r := bytes.NewReader(oJson)
	//return r


type MyObject struct {
	Word string		`json:"string"`
	Number int		`json:"number"`
}

func Post() {
	r := MakeRandJSON()
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
