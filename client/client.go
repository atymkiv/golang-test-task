package main

import (
	
	"sync"
	"os"
	"strconv"
	"math/rand"
	
	pool "github.com/atymkiv/golang-test-task/client/workerPool"
)

func NewRequest(wg *sync.WaitGroup) pool.Request {
			var myRequest pool.Request
			var arr []byte

			//generating random words of random chars and length
			for j := 0; j < rand.Intn(50); j++ {  //rand length
				arr = append(arr, byte(rand.Intn(122-65)+65)) //rand chars 
			}

			myRequest.Word = string(arr)
			myRequest.Number = rand.Intn(100000)

	return myRequest
}	


func main() {
	
	bufferSize := 100
	var dispatcher pool.Dispatcher = pool.NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		var w pool.WorkerLauncher = &pool.PostWorker{
		}
		dispatcher.LaunchWorker(w)
	}

	requests, _ := strconv.Atoi(os.Args[1])

	var wg sync.WaitGroup
	wg.Add(requests-1)

	for i := 0; i < requests; i++ {
		req := NewRequest(&wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	wg.Wait()
}
