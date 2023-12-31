//Write by Diego R Rocha
// Package xarray provides a simple implementation of integer arrays
// with various utility methods such as Push, Pop, ForEach, Map, Filter, Find,
// Every, Some, Reduce, and Sort.
package xarray

// IntArray represents an integer array.
type IntArray []int

// MapFunc defines the type for a mapping function.
type MapFunc func(int) int

// FilterFunc defines the type for a filtering function.
type FilterFunc func(int) bool

// ReduceFunc defines the type for a reducing function.
type ReduceFunc func(int, int) int

// Push appends an integer to the end of the array.
func (a *IntArray) Push(value int) {
	*a = append(*a, value)
}

// Pop removes and returns the last element of the array.
// It returns 0 if the array is empty.
func (a *IntArray) Pop() int {
	if len(*a) == 0 {
		return 0
	}
	lastIndex := len(*a) - 1
	value := (*a)[lastIndex]
	*a = (*a)[:lastIndex]
	return value
}

// ForEach applies a callback function to each element of the array.
func (a IntArray) ForEach(callback func(int)) {
	for _, v := range a {
		callback(v)
	}
}

// Map applies a mapping function to each element of the array
// and returns a new mapped array.
func (a IntArray) Map(callback MapFunc) IntArray {
	result := make(IntArray, len(a))
	for i, v := range a {
		result[i] = callback(v)
	}
	return result
}

// Filter applies a filtering function to each element of the array
// and returns a new array containing only the elements that satisfy the condition.
func (a IntArray) Filter(callback FilterFunc) IntArray {
	var result IntArray
	for _, v := range a {
		if callback(v) {
			result.Push(v)
		}
	}
	return result
}

// Find applies a filtering function to find the first element in the array
// that satisfies the condition. It returns 0 if no such element is found.
func (a IntArray) Find(callback FilterFunc) int {
	for _, v := range a {
		if callback(v) {
			return v
		}
	}
	return 0
}

// Every checks if all elements in the array satisfy a given condition.
func (a IntArray) Every(callback FilterFunc) bool {
	for _, v := range a {
		if !callback(v) {
			return false
		}
	}
	return true
}

// Some checks if at least one element in the array satisfies a given condition.
func (a IntArray) Some(callback FilterFunc) bool {
	for _, v := range a {
		if callback(v) {
			return true
		}
	}
	return false
}

// Reduce applies a reducing function to the elements of the array
// and returns a single accumulated result.
func (a IntArray) Reduce(callback ReduceFunc, initialValue int) int {
	accumulator := initialValue
	for _, v := range a {
		accumulator = callback(accumulator, v)
	}
	return accumulator
}

// Sort performs a bubble sort on the elements of the array.
func (a IntArray) Sort() {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
