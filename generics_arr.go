package generics

import "sort"

// This is a generic slice, it provides out of the
// box methods like AddElements, Sort, SortDesc
// Supported values for this GenericSlice are
// all types of integers and all types of floats
type GenericSlice[T addable] []T

// adds all the elements from the generic type
func (arr GenericSlice[T]) AddElements() T {
	var finalSum T
	for _, ele := range arr {
		finalSum += ele
	}

	return finalSum
}

// Sort sorts a GenericSlice into ascending order
func (arr GenericSlice[T]) Sort() {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

// SortDesc sorts a GenericSlice into descending order
func (arr GenericSlice[T]) SortDesc() {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}

// FindElement finds the element passed in this function and
// returns the index of the same element. It also returns
// a boolean indicating if the element was found or not.
// NOTE - This function finds the first index of the element
// passed.
func (arr GenericSlice[T]) FindElement(ele T) (int, bool) {
	indx := -1
	var isFound bool

	for i := range arr {
		if arr[i] == ele {
			indx = i
			isFound = true
			break
		}
	}

	return indx, isFound
}
