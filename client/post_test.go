package main

import (
	"testing"
	//"fmt"
	"io"
	
)
func TestMakeRandJSON(t *testing.T) {
	JsonCh := MakeRandJSON(2) //channel of objects
	if JsonCh == nil{
		//Test of creation
		t.Error("Expected new channel object after calling MakeRandJSON, not nil")
	}

		for object := range JsonCh {
			var ob interface{} = object 
			_, ok := ob.(MyObject)
			if !ok {
				t.Error("Expected channel of MyObjects")
		}
	}			
}

func TestJsonToReader(t *testing.T) {
	ReaderCh := JsonToReader(MakeRandJSON(2)) //Channel of io.Reader

	if ReaderCh == nil {
		//Test of creation
		t.Error("Expected new channel object after calling MakeRandJSON, not nil")
	}
	for object := range ReaderCh {
		_, ok := object.(io.Reader)
		if !ok {
			t.Error("Expected channel of io.Reader objects")
		}
	}
	
} 
		
