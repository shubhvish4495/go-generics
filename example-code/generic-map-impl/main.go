package main

import (
	"fmt"

	"github.com/shubhvish4495/go-generics"
)

func main() {

	m := make(generics.GenericMap[int, string])

	m.Put(1, "test")
	if v, ok := m.Get(1); ok {
		fmt.Println(v)
	}
	m.DeleteElement(1)

	m2 := make(generics.GenericMap[string, bool])
	m2.Put("b-test", false)
	m2.Put("a-test", true)

	fmt.Println(m2.GetSortedKeys())
	//this will be ignored as key is already stored
	m2.PutIfAbsent("b-test", true)
	m2.PutIfAbsent("c-test", false)

	vb, _ := m2.Get("b-test")
	vc, _ := m2.Get("c-test")

	fmt.Println("Insert with PutIfAbsent for key not present in map->", vc)
	fmt.Println("Insert with PutIfAbsent will be ignored if key present-> ", vb)
}
