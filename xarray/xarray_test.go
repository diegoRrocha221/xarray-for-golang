//Write By Diego R Rocha
// Package xarray provides tests for the xarray package, which implements
// integer arrays with various utility methods.
package xarray

import (
	"reflect"
	"testing"
)

// TestPush tests the Push method of IntArray.
func TestPush(t *testing.T) {
	var arr IntArray
	arr.Push(1)

	if len(arr) != 1 || arr[0] != 1 {
		t.Errorf("Push failed. Expected [1], got %v", arr)
	}
}

// TestPop tests the Pop method of IntArray.
func TestPop(t *testing.T) {
	arr := IntArray{1, 2, 3}
	popped := arr.Pop()

	if len(arr) != 2 || popped != 3 {
		t.Errorf("Pop failed. Expected popped value 3 and array [1 2], got popped: %v, array: %v", popped, arr)
	}
}

// TestForEach tests the ForEach method of IntArray.
func TestForEach(t *testing.T) {
	arr := IntArray{1, 2, 3}
	var result IntArray

	arr.ForEach(func(value int) {
		result.Push(value * 2)
	})

	expected := IntArray{2, 4, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ForEach failed. Expected %v, got %v", expected, result)
	}
}

// TestMap tests the Map method of IntArray.
func TestMap(t *testing.T) {
	arr := IntArray{1, 2, 3}
	result := arr.Map(func(value int) int {
		return value * 2
	})

	expected := IntArray{2, 4, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map failed. Expected %v, got %v", expected, result)
	}
}

// TestFilter tests the Filter method of IntArray.
func TestFilter(t *testing.T) {
	arr := IntArray{1, 2, 3, 4, 5}
	result := arr.Filter(func(value int) bool {
		return value%2 == 0
	})

	expected := IntArray{2, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter failed. Expected %v, got %v", expected, result)
	}
}

// TestFind tests the Find method of IntArray.
func TestFind(t *testing.T) {
	arr := IntArray{1, 2, 3, 4, 5}
	result := arr.Find(func(value int) bool {
		return value > 3
	})

	expected := 4
	if result != expected {
		t.Errorf("Find failed. Expected %v, got %v", expected, result)
	}
}

// TestEvery tests the Every method of IntArray.
func TestEvery(t *testing.T) {
	arr := IntArray{2, 4, 6}
	result := arr.Every(func(value int) bool {
		return value%2 == 0
	})

	if !result {
		t.Errorf("Every failed. Expected true, got false")
	}

	arr = IntArray{2, 4, 5}
	result = arr.Every(func(value int) bool {
		return value%2 == 0
	})

	if result {
		t.Errorf("Every failed. Expected false, got true")
	}
}

// TestSome tests the Some method of IntArray.
func TestSome(t *testing.T) {
	arr := IntArray{1, 3, 5}
	result := arr.Some(func(value int) bool {
		return value%2 == 0
	})

	if result {
		t.Errorf("Some failed. Expected false, got true")
	}

	arr = IntArray{1, 2, 3}
	result = arr.Some(func(value int) bool {
		return value%2 == 0
	})

	if !result {
		t.Errorf("Some failed. Expected true, got false")
	}
}

// TestReduce tests the Reduce method of IntArray.
func TestReduce(t *testing.T) {
	arr := IntArray{1, 2, 3, 4}
	result := arr.Reduce(func(acc, value int) int {
		return acc + value
	}, 0)

	expected := 10
	if result != expected {
		t.Errorf("Reduce failed. Expected %v, got %v", expected, result)
	}
}

// TestSort tests the Sort method of IntArray.
func TestSort(t *testing.T) {
	arr := IntArray{3, 1, 4, 1, 5}
	arr.Sort()

	expected := IntArray{1, 1, 3, 4, 5}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Sort failed. Expected %v, got %v", expected, arr)
	}
}

// TestCharOperations tests the character operations on IntArray.
func TestCharOperations(t *testing.T) {
	charArr := IntArray{'a', 'b', 'c', 'd'}

	charArr.Push('e')
	expected := IntArray{'a', 'b', 'c', 'd', 'e'}
	if !reflect.DeepEqual(charArr, expected) {
		t.Errorf("Push failed. Expected %v, got %v", expected, charArr)
	}

	popped := charArr.Pop()
	if len(charArr) != 4 || popped != 'e' {
		t.Errorf("Pop failed. Expected popped value 'e' and array %v, got popped: %v, array: %v", expected, popped, charArr)
	}

	var result IntArray
	charArr.ForEach(func(value int) {
		result.Push(value + 1)
	})
	expected = IntArray{'b', 'c', 'd', 'e'}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ForEach failed. Expected %v, got %v", expected, result)
	}

	result = charArr.Map(func(value int) int {
		return value + 1
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map failed. Expected %v, got %v", expected, result)
	}

	result = charArr.Filter(func(value int) bool {
		return value%2 == 0
	})
	expected = IntArray{'b', 'd'}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter failed. Expected %v, got %v", expected, result)
	}

	resultChar := charArr.Find(func(value int) bool {
		return value > 'c'
	})
	expectedChar := 'd'
	if int(resultChar) != int(expectedChar) {
		t.Errorf("Find failed. Expected %v, got %v", expectedChar, resultChar)
	}
	resultBool := charArr.Every(func(value int) bool {
		return value >= 'a'
	})
	if !resultBool {
		t.Errorf("Every failed. Expected true, got false")
	}

	resultBool = charArr.Some(func(value int) bool {
		return value > 'c'
	})
	if !resultBool {
		t.Errorf("Some failed. Expected true, got false")
	}

	resultInt := charArr.Reduce(func(acc, value int) int {
		return acc + int(value)
	}, 0)
	expectedInt := int('a' + 'b' + 'c' + 'd')
	if resultInt != expectedInt {
		t.Errorf("Reduce failed. Expected %v, got %v", expectedInt, resultInt)
	}

	charArr.Sort()
	expectedSorted := IntArray{'a', 'b', 'c', 'd'}
	if !reflect.DeepEqual(charArr, expectedSorted) {
		t.Errorf("Sort failed. Expected %v, got %v", expectedSorted, charArr)
	}
}
