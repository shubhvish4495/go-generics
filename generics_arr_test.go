package generics

import "testing"

func TestAddElements(t *testing.T) {

	sliceVal := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	testSlice := make(GenericSlice[int], 0)
	testSlice = append(testSlice, sliceVal...)

	addVal := testSlice.AddElements()
	if addVal != 55 {
		t.Fail()
	}

}

func TestAddElementsEmptySlice(t *testing.T) {

	testSlice := make(GenericSlice[int], 0)
	addVal := testSlice.AddElements()

	if addVal != 0 {
		t.Fail()
	}
}

func TestSort(t *testing.T) {

	sliceVal := []int{8, 10, 2, 4, 834, 47, 3}
	expectedSliceVal := []int{2, 3, 4, 8, 10, 47, 834}

	testSlice := make(GenericSlice[int], 0)
	testSlice = append(testSlice, sliceVal...)

	testSlice.Sort()

	for i := range testSlice {
		if testSlice[i] != expectedSliceVal[i] {
			t.Fail()
		}
	}

}

func TestSortDesc(t *testing.T) {

	sliceVal := []int{8, 10, 2, 4, 834, 47, 3}
	expectedSliceVal := []int{2, 3, 4, 8, 10, 47, 834}
	n := len(expectedSliceVal) - 1

	testSlice := make(GenericSlice[int], 0)
	testSlice = append(testSlice, sliceVal...)

	testSlice.SortDesc()

	for i := range testSlice {
		if testSlice[i] != expectedSliceVal[n-i] {
			t.Fail()
		}
	}

}

func TestFindElem(t *testing.T) {

	sliceVal := []int{8, 10, 2, 4, 834, 47, 3}

	testSlice := make(GenericSlice[int], 0)
	testSlice = append(testSlice, sliceVal...)

	indx, found := testSlice.FindElement(834)

	if indx != 4 && found != true {
		t.Fail()
	}
}

func TestFindElemWithNonExistentElem(t *testing.T) {

	sliceVal := []int{8, 10, 2, 4, 834, 47, 3}

	testSlice := make(GenericSlice[int], 0)
	testSlice = append(testSlice, sliceVal...)

	indx, found := testSlice.FindElement(-9)

	if indx != -1 && found != false {
		t.Fail()
	}
}

func TestFindElemWithRepeatedElem(t *testing.T) {

	sliceVal := []int{8, 10, 834, 4, 834, 47, 3}

	testSlice := make(GenericSlice[int], 0)
	testSlice = append(testSlice, sliceVal...)

	indx, found := testSlice.FindElement(834)

	if indx != 2 && found != false {
		t.Fail()
	}
}
