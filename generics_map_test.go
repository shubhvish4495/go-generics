package generics

import (
	"testing"
)

func TestGet(t *testing.T) {
	testMap := make(GenericMap[int, string], 0)
	testMap.Put(1, "test")
	val, ok := testMap.Get(1)

	if ok != true && val != "test" {
		t.Fail()
	}
}

func TestGetWithNoData(t *testing.T) {
	testMap := make(GenericMap[int, string], 0)
	val, ok := testMap.Get(1)

	if ok != false && val != "" {
		t.Fail()
	}
}

func TestPut(t *testing.T) {
	testMap := make(GenericMap[int, string], 0)
	testMap.Put(1, "test")
	val, ok := testMap.Get(1)

	if ok != false && val != "test" {
		t.Fail()
	}
}

func TestPutForSameKey(t *testing.T) {
	testMap := make(GenericMap[int, string], 0)
	testMap.Put(1, "test-1")
	testMap.Put(1, "test-2")
	val, ok := testMap.Get(1)

	if ok != false && val != "test-2" {
		t.Fail()
	}
}

func TestPutIfAbsent(t *testing.T) {
	testMap := make(GenericMap[int, string], 0)
	testMap.Put(1, "test-1")
	testMap.PutIfAbsent(1, "test-2")

	val, _ := testMap.Get(1)

	if val != "test-1" {
		t.Fail()
	}
}

func TestDeleteElement(t *testing.T) {
	testMap := make(GenericMap[int, string], 0)
	testMap.Put(1, "test-1")
	testMap.DeleteElement(1)

	v, ok := testMap.Get(1)
	if ok != false && v != "" {
		t.Fail()
	}
}

// TestDeleteElementInvalidKey is the test to check
// if delete works in case of invalid key or not.
// This test does not have any t.Fail stub as if it were to fail
// we woould expect a panic.
func TestDeleteElementInvalidKey(t *testing.T) {
	testMap := make(GenericMap[int, string], 0)

	testMap.DeleteElement(1)
}

func TestGetSortedKeysForIntKeys(t *testing.T) {
	testMap := make(GenericMap[int, string], 0)
	testMap.Put(10, "test-1")
	testMap.Put(2, "test-2")
	testMap.Put(0, "test-3")

	sortedKeys := testMap.GetSortedKeys()
	expected := []int{0, 2, 10}

	for i := range sortedKeys {
		if sortedKeys[i] != expected[i] {
			t.Fail()
			break
		}
	}
}

func TestGetSortedKeysForStringKeys(t *testing.T) {
	testMap := make(GenericMap[string, string], 0)
	testMap.Put("m-test", "1")
	testMap.Put("c-test", "2")
	testMap.Put("x-test", "3")

	sortedKeys := testMap.GetSortedKeys()
	expected := []string{"c-test", "m-test", "x-test"}

	for i := range sortedKeys {
		if sortedKeys[i] != expected[i] {
			t.Fail()
			break
		}
	}
}
