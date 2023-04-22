package generics

import "sort"

// GenericMap is a generic implementation of map
// where we have keys which can be of type all possible ints
// all possible floats and string. For values we have
// supported types as all possible ints, all possible floats
// string and boolean. NOTE - this is not thread safe map
type GenericMap[T genericKeys, V genericVals] map[T]V

// Get value from a custom map
func (c GenericMap[T, V]) Get(key T) (V, bool) {
	v, ok := c[key]
	return v, ok
}

// Put values inside a custom map
func (c GenericMap[T, V]) Put(key T, value V) {
	c[key] = value
}

// Put values inside a custom map if key is not present
func (c GenericMap[T, V]) PutIfAbsent(key T, value V) {
	if _, present := c.Get(key); !present {
		c[key] = value
	}
}

// This function deletes element from the map for which
// the key has been passed. This is a simple wrapper function
// over the delete() function. In case the key is nil or
// the key is not present in the map, it will be no op operation
func (c GenericMap[T, V]) DeleteElement(key T) {
	delete(c, key)
}

// GetSortedKeys returns sorted keys of map
func (c GenericMap[T, V]) GetSortedKeys() []T {
	var keys []T

	// get all the keys
	for key := range c {
		keys = append(keys, key)
	}

	// sort all the keys
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// return the sorted keys
	return keys
}
