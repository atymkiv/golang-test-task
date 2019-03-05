package main

import (
	"testing"
	"fmt"
	"encoding/json"
	//"bytes"
	//"net/http"
)
func TestMakeRand(t *testing.T) {
	obj := MakeRand()
	if obj == nil {
		//Test of creation
		t.Error("Expected new io.Reader object after calling MakeRand, not nil")
	}
		var object MyObject
		err := obj.Decode(&object)

		fmt.Printf("Word is %s, Number is %d\n", object.Word, object.Number)
}
