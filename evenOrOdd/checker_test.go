package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestChecker(t *testing.T) {

	numList := createList()

	//We should make sure the list is not empty and maybe of int8 type

	if len(numList) == 0 {
		fmt.Println("List is empty")
	}

	numList.TypeToString()

	if reflect.TypeOf(numList).Kind() != reflect.Int8 {
		fmt.Println("numList contains integers as expected")

	} else {
		fmt.Println("Bad character(s) in list")

	}

}

// need to convert type to string
func (k numberList) TypeToString() string {
	numT := reflect.TypeOf(k).Kind()
	newString := fmt.Sprintf("%v", numT)
	return newString
}
