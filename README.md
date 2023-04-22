# <img src="logo.png" width="50" height="50"> <i><b>Go-Generics</b></i>
This is a repository for generic implementation of Map and Slice. This repo provides out of the box some general use methods which can reduce our code. Just import this once and use it's generic implementation for different types of slices and maps

### Requirements
Supported for Go <b>1.18+</b>


Get this package

```shell
go get github.com/shubhvish4495/go-generics
```

### How to use this

```golang
package main

import (
	"fmt"

	"github.com/shubhvish4495/go-generics"
)

func main() {
	arr := make(generics.GenericSlice[int], 0)
	sliceVal := []int{8, 10, 2, 4, 834, 47, 3}

	arr = append(arr, sliceVal...)

    	//This sorting is inplace
	arr.Sort()
	fmt.Println("Sorted array ", arr)
}
```

Complete implementation example can be found in `example-code` sub-dir of this project
