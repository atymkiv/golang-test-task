package main

import (
	"testing"
	//"fmt"
	"io"
	//"sync"
	pool "github.com/atymkiv/golang-test-task/client/workerPool"
)

//var wg sync.WaitGroup
func TestNewRequest(t *testing.T) {
	
	request := NewRequest() //channel of objects
	
	//test of creation
	var ob interface{} = request
	_, ok := ob.(pool.Request)
	if !ok {
		t.Error("Expected Request")
	}
			
}

func TestObjToJson(t *testing.T) {
	ReqCh := make(chan pool.Request, 2) 
	for i := 0; i<2; i++{
		ReqCh <- NewRequest()
		
	}

	JsonCh := pool.ObjToJson(ReqCh) //Channel of jsons
	close(ReqCh)
	if JsonCh == nil {
		//Test of creation
		t.Error("Expected new channel object after calling ObjToJson, not nil")
	}

	for object := range JsonCh {
		var ob interface{} = object
		_, ok := ob.([]byte)
		if !ok {
			t.Error("Expected channel of []byte objects")
		}
	}
}

func TestJsonToReader(t *testing.T) {
	ReqCh := make(chan pool.Request, 2) 
	for i := 0; i<2; i++{
		ReqCh <- NewRequest()
		
	}

	

	ReaderCh := pool.JsonToReader(pool.ObjToJson(ReqCh))
	close(ReqCh)

	if ReaderCh == nil {
		//Test of creation
		t.Error("Expected new channel object after calling JsonToReader, not nil")
	}
	for object := range ReaderCh {
		_, ok := object.(io.Reader)
		if !ok {
			t.Error("Expected channel of io.Reader objects")
		}
	}
}

func benchmarkPost(requests int, workers int, b *testing.B) {
	bufferSize := 100
	var dispatcher pool.Dispatcher = pool.NewDispatcher(bufferSize)

	for i := 0; i < workers; i++ {
		var w pool.WorkerLauncher = &pool.PostWorker{
		}
		dispatcher.LaunchWorker(w)
	}

	for i := 0; i < requests; i++ {
		req := NewRequest()
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	
}

//func BenchmarkPost1000000_1worker(b *testing.B) { benchmarkPost(1000000, 1, b) }
func BenchmarkPost1000000_2workers(b *testing.B) { benchmarkPost(1000000, 2, b) }		
//func BenchmarkPost1000000_3workers(b *testing.B) { benchmarkPost(1000000, 3, b) }
//func BenchmarkPost1000000_4workers(b *testing.B) { benchmarkPost(1000000, 4, b) }
