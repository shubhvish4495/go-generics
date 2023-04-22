package main

import (
	"fmt"

	"github.com/shubhvish4495/go-generics"
)

func main() {
	arr := make(generics.GenericSlice[int], 0)
	sliceVal := []int{8, 10, 2, 4, 834, 47, 3}

	arr = append(arr, sliceVal...)

	fmt.Println("Addition value of all the elements of slice ", arr.AddElements())

	v, ok := arr.FindElement(834)
	fmt.Println("Finds the index of the element passed ", v, " also returns bool value as ", ok)

	vnf, oknf := arr.FindElement(837)
	fmt.Println("Will return ", vnf, " incase element is not found and will return bool value as ", oknf)

	//This sorting is inplace
	arr.Sort()
	fmt.Println("Sorted array ", arr)

	//This sorting is inplace
	arr.SortDesc()
	fmt.Println("Desc sorted array ", arr)
}
