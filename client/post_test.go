package main

import (
	"testing"
	//"fmt"
	"io"
	
)
func TestMakeRandObj(t *testing.T) {
	JsonCh := MakeRandObj(2) //channel of objects
	if JsonCh == nil{
		//Test of creation
		t.Error("Expected new channel object after calling MakeRandObj, not nil")
	}

		for object := range JsonCh {
			var ob interface{} = object 
			_, ok := ob.(MyObject)
			if !ok {
				t.Error("Expected channel of MyObjects")
		}
	}			
}

func TestObjToJson(t *testing.T) {
	JsonCh := ObjToJson(MakeRandObj(2)) //Channel of io.Reader

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
	ReaderCh := JsonToReader(ObjToJson(MakeRandObj(2)))
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

/*func benchmarkPost(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Post(ObjToJson(MakeRandObj(i)))
	}
}

func BenchmarkPost10(b *testing.B) { benchmarkPost(10, b) } 
func BenchmarkPost100(b *testing.B) { benchmarkPost(100, b) }
func BenchmarkPost1000(b *testing.B) { benchmarkPost(1000, b) }
func BenchmarkPost10000(b *testing.B) { benchmarkPost(10000, b) }
func BenchmarkPost100000(b *testing.B) { benchmarkPost(100000, b) }
func BenchmarkPost1000000(b *testing.B) { benchmarkPost(1000000, b) }
		
*/