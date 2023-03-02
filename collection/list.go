package collection

import "errors"

type List[T any] struct {
	elements []T
}

// Returns the number of elements in the list
func (this *List[T]) Count() int {
	return len(this.elements)
}

// Adds and element to the list
func (this *List[T]) Add(t T) *List[T] {
	this.elements = append(this.elements, t)
	return this
}

// Add adds a range of element to the list
func (this *List[T]) AddRange(t []T) *List[T] {
	this.elements = append(this.elements, t...)
	return this
}

// OrderedRemove removes the element at index i in an ordered fashion
// meaning that internal order of the lists elements will not change
func (this *List[T]) OrderedRemove(i int) error {
	if i > this.Count()-1 || i < 0 {
		return errors.New("index out of bounds")
	}
	this.elements = append(this.elements[:i], this.elements[i+1:]...)
	return nil
}

// UnorderedRemove removes the element at index i in an unordered fashion
// meaning that internal order of the lists elements may change
func (this *List[T]) UnorderedRemove(i int) error {
	if i > this.Count()-1 || i < 0 {
		return errors.New("index out of bounds")
	}
	this.elements[i] = this.elements[this.Count()-1]
	this.elements = this.elements[:this.Count()-1]
	return nil
}

// Get return element at index i
func (this *List[T]) Get(i int) (T, error) {
	var result T
	if i > this.Count()-1 || i < 0 {
		return result, errors.New("index out of bounds")
	}
	result = this.elements[i]
	return result, nil
}

// getUnsafe is only used internally and should not be used outside
// of this package. getUnsafe relies on calling func to do bounds
// checking.
func (this *List[T]) getUnsafe(i int) T {
	result := this.elements[i]
	return result
}

// Foreach loops over all elements in the list
func (this *List[T]) Foreach(f func(T)) {
	for i := 0; i < this.Count(); i++ {
		f(this.getUnsafe(i))
	}
}

// Filter filters all elements in the list based on the supplied func
func (this *List[T]) Filter(f func(T) bool) *List[T] {
	var result List[T]
	for i := 0; i < this.Count(); i++ {
		v := this.getUnsafe(i)
		if f(v) {
			result.Add(v)
		}
	}
	return &result
}
