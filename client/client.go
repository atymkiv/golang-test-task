package main

import (
	
	"math/rand"
	"encoding/json"
	"bytes"
	"io"
	"net/http"
	"fmt"
	"os"
	"strconv"
	
)

func MakeRandObj(max int) <-chan MyObject{
	outChJson := make(chan MyObject, max)
	
	go func() {
		for i := 1; i <= max; i++ {
			var object MyObject
			var arr []byte

			//generating random words of random chars and length
			for j := 0; j < rand.Intn(50); j++ {  //rand length
				arr = append(arr, byte(rand.Intn(122-65)+65)) //rand chars 
			}

			object.Word = string(arr)
			object.Number = rand.Intn(100000)

			outChJson <- object
		}

		close(outChJson)
	}()

	return outChJson
}
	
func ObjToJson(in <-chan MyObject) <-chan []byte {
	out := make(chan []byte, 100)
	go func(){
		for object := range in{
			oJson, _ := json.Marshal(object)
			out <- oJson
		}
		close(out)	
	}()
	return out
}

func JsonToReader(in <-chan []byte) <-chan io.Reader {
	out := make(chan io.Reader, 100)
	go func(){
		for oJson := range in{
			r := bytes.NewReader(oJson)
			out <- r
		}
		close(out)	
	}()
	return out		
}

type MyObject struct {
	Word string		`json:"string"`
	Number int		`json:"number"`
}

func Post(in <-chan io.Reader) error{
	go func(){
		for object := range in {
			_, err := http.Post("http://localhost:9000/", "application/json", object)
			
			if err != nil{
				fmt.Println(err)
				
			}
		}
	}()
	return nil 
}

func main() {
	length, _ := strconv.Atoi(os.Args[1])
	Post(JsonToReader(ObjToJson(MakeRandObj(length))))
    var input string
    fmt.Scanln(&input)

}
