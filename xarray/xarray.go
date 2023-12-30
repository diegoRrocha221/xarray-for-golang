package xarray

type IntArray []int

type MapFunc func(int) int

type FilterFunc func(int) bool

type ReduceFunc func(int, int) int

func (a *IntArray) Push(value int) {
	*a = append(*a, value)
}

func (a *IntArray) Pop() int {
	if len(*a) == 0 {
		return 0
	}
	lastIndex := len(*a) - 1
	value := (*a)[lastIndex]
	*a = (*a)[:lastIndex]
	return value
}

func (a IntArray) ForEach(callback func(int)) {
	for _, v := range a {
		callback(v)
	}
}

func (a IntArray) Map(callback MapFunc) IntArray {
	result := make(IntArray, len(a))
	for i, v := range a {
		result[i] = callback(v)
	}
	return result
}

func (a IntArray) Filter(callback FilterFunc) IntArray {
	var result IntArray
	for _, v := range a {
		if callback(v) {
			result.Push(v)
		}
	}
	return result
}

func (a IntArray) Find(callback FilterFunc) int {
	for _, v := range a {
		if callback(v) {
			return v
		}
	}
	return 0
}

func (a IntArray) Every(callback FilterFunc) bool {
	for _, v := range a {
		if !callback(v) {
			return false
		}
	}
	return true
}

func (a IntArray) Some(callback FilterFunc) bool {
	for _, v := range a {
		if callback(v) {
			return true
		}
	}
	return false
}

func (a IntArray) Reduce(callback ReduceFunc, initialValue int) int {
	accumulator := initialValue
	for _, v := range a {
		accumulator = callback(accumulator, v)
	}
	return accumulator
}

func (a IntArray) Sort() {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}