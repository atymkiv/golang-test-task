package main

import (
	"testing"
	"fmt"
)
func TestMakeRandJSON(t *testing.T) {
	JsonCh := MakeRandJSON(2)
	if JsonCh == nil {
		//Test of creation
		t.Error("Expected new channel object after calling MakeRandJSON, not nil")
	}

		for object := range JsonCh {
			fmt.Printf("Word is %s, Number is %d\n", object.Word, object.Number)	
		} 

		
}
