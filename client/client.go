package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"log"
	"runtime/pprof"
	"flag"
	"time"
	pool "github.com/atymkiv/golang-test-task/client/workerPool"
)

func NewRequest() pool.Request {
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


var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	singletonCounter := pool.GetInstance() // Concurrent singleton to count post requests
	defer singletonCounter.Stop() 
	timer := time.NewTimer(60 *time.Second) // timer to count requests per interval
	go func() {
		select{
		case <- timer.C: 
		fmt.Println(singletonCounter.GetCount()) //print requests per time interval
		
		}
	}()
	
 	// make pprof 
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	
	bufferSize := 1000
	var dispatcher pool.Dispatcher = pool.NewDispatcher(bufferSize) 

	workers := 4
	for i := 0; i < workers; i++ {
		var w pool.WorkerLauncher = &pool.PostWorker{
		}
		dispatcher.LaunchWorker(w)
	}

	requests, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < requests; i++ {
		req := NewRequest()
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()
	
	var input string
	fmt.Scanln(&input)
}
