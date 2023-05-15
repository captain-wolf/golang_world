package main

import "fmt"

type numberList []int8

func createList() numberList {

	//numberlist type of int8
	var noList = numberList{}

	noList = []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	return noList
}

func (n numberList) evenOrOddNumbers() numberList {
	filteredList := numberList{}

	for _, listItem := range filteredList {
		if listItem%2 == 0 {
			fmt.Printf("%v is even.", listItem)
		} else {
			fmt.Printf("%v is odd.", listItem)
		}
	}
	return filteredList
}
