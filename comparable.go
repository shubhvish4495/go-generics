package generics

import (
	c "golang.org/x/exp/constraints"
)

// Generic type created using constraints
// It is an interface for all types of integers and floats
type addable interface {
	c.Integer | c.Float
}

// Generic Vals is an interface for all integers, floats
// string and boolean
type genericVals interface {
	c.Ordered | ~bool
}

// Generic keys is an interface for all addable types and
// string as well
type genericKeys interface {
	addable | ~string
}
