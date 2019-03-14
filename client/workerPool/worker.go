package workerPool

import (
	"fmt"
	"encoding/json"
	"bytes"
	"io"
	"net/http"
	
	
	
)



type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

//develop pipeline
//a type to implement WorkerLauncher type:
type PostWorker struct {}

func (w *PostWorker) LaunchWorker(in chan Request) {
	Post(JsonToReader(ObjToJson(in)))
}

func ObjToJson(in <-chan Request) <-chan []byte {
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


func Post(in <-chan io.Reader) {
	go func(){
		for object := range in {
			_, err := http.Post("http://localhost:9000/", "application/json", object)
			
			if err != nil {
				fmt.Println(err)
				
			}
		}
	}()
}
