<h1 align="center">xarray v0.1</h1>

<p align="center">
  [![Go Report Card](https://goreportcard.com/badge/github.com/diegorrocha221/xarray-for-golang)](https://goreportcard.com/report/github.com/diegorrocha221/xarray-for-golang)
  <a href="https://pkg.go.dev/github.com/diegoRrocha221/xarray-for-golang"><img src="https://pkg.go.dev/badge/github.com/diegoRrocha221/xarray-for-golang.svg" alt="Go Reference"></a>
  <a href="https://github.com/diegoRrocha221/xarray-for-golang/blob/main/LICENSE"><img src="https://img.shields.io/github/license/diegoRrocha221/xarray-for-golang" alt="GitHub License"></a>
</p>

<p align="center">xarray is a lightweight Go package that provides array manipulation methods inspired by TypeScript. It is designed to simplify common array operations, offering methods such as Push, Pop, ForEach, Map, Filter, Find, Every, Some, Reduce, and Sort. The package supports both integer and character arrays, making it versatile for various use cases.</p>

<h2 align="center">Installation</h2>

<p align="center">To use `xarray` in your Go project, you can use the `go get` command:</p>

```bash
go get github.com/diegoRrocha221/xarray-for-golang
```
<h2 align="center">Usage</h2>
<h3 align="center">IntArray Type</h3>
<details>
<summary><code>Push</code></summary>

```bash
// Push values to the array
arr.Push(1)
arr.Push(2)
arr.Push(3)
```
Appends the specified values to the end of the integer array.

</details>
<details>
<summary><code>Pop</code></summary>

```bash
// Pop a value from the array
popped := arr.Pop()
```
Removes and returns the last element from the integer array.
</details>
<details>
<summary><code>ForEach</code></summary>

```bash
// Iterate over the array
arr.ForEach(func(value int) {
    // Your callback logic here
})
```

Executes a provided function once for each element in the integer array.
</details>
<details>
<summary><code>Map</code></summary>

```bash
// Map values in the array
mapped := arr.Map(func(value int) int {
    return value * 2
})
```
Creates a new integer array with the results of calling a provided function on every element in the array.
</details>
<details>
<summary><code>Filter</code></summary>

```bash
// Filter values in the array
filtered := arr.Filter(func(value int) bool {
    return value%2 == 0
})
```

Creates a new integer array with all elements that pass the test implemented by the provided function.
</details>
<details>
<summary><code>Find</code></summary>

```bash
// Find a specific value in the array
found := arr.Find(func(value int) bool {
    return value > 1
})
```
 Returns the first element in the array that satisfies the provided testing function.
</details>
<details>
<summary><code>Every</code></summary>


```bash
// Check if every element satisfies a condition
allSatisfy := arr.Every(func(value int) bool {
    return value > 0
})
```
Tests whether all elements in the array pass the test implemented by the provided function.
</details>
<details>
<summary><code>Some</code></summary>

```bash
// Check if any element satisfies a condition
someSatisfy := arr.Some(func(value int) bool {
    return value == 2
})
```
Tests whether at least one element in the array passes the test implemented by the provided function.
</details>
<details>
<summary><code>Reduce</code></summary>

```bash
// Reduce the array to a single value
reduced := arr.Reduce(func(acc, value int) int {
    return acc + value
}, 0)
```
Applies a function against an accumulator and each element in the array (from left to right) to reduce it to a single value.
</details>
<details>
<summary><code>Sort</code></summary>

```bash
// Sort the array
arr.Sort()
```
 Sorts the elements of the array in ascending order.
</details>
<h3 align="center">CharArray Type</h3>
<details>
<summary><code>Push</code></summary>

```bash
// Push a value to the array
charArr.Push('d')
```
Description: Appends the specified character to the end of the character array.
</details>
<details>
<summary><code>Pop</code></summary>

```bash
// Pop a value from the array
popped := charArr.Pop()
```

Removes and returns the last character from the character array.
</details>
<details>
<summary><code>ForEach</code></summary>

```bash
// Iterate over the array
charArr.ForEach(func(value int) {
    // Your callback logic here
})
```
Executes a provided function once for each character in the character array.
</details>
<details>
<summary><code>Map</code></summary>

```bash
// Map values in the array
mapped := charArr.Map(func(value int) int {
    return value + 1
})
```
Creates a new character array with the results of calling a provided function on every character in the array.
</details>
<details>
<summary><code>Filter</code></summary>

```bash
// Filter values in the array
filtered := charArr.Filter(func(value int) bool {
    return value%2 == 0
})
```
Creates a new character array with all characters that pass the test implemented by the provided function.
</details>
<details>
<summary><code>Find</code></summary>

```bash
// Find a specific value in the array
found := charArr.Find(func(value int) bool {
    return value > 'c'
})
```

Returns the first character in the array that satisfies the provided testing function.
</details>
<details>
<summary><code>Every</code></summary>

```bash
// Check if every character satisfies a condition
allSatisfy := charArr.Every(func(value int) bool {
    return value >= 'a'
})
```
 Tests whether all characters in the array pass the test implemented by the provided function.
</details>
<details>
<summary><code>Some</code></summary>

```bash
// Check if any character satisfies a condition
someSatisfy := charArr.Some(func(value int) bool {
    return value > 'c'
})
```

Tests whether at least one character in the array passes the test implemented by the provided function.
</details>
<details>
<summary><code>Reduce</code></summary>

```bash
// Reduce the array to a single value
reduced := charArr.Reduce(func(acc, value int) int {
    return acc + int(value)
}, 0)
```

Applies a function against an accumulator and each character in the array (from left to right) to reduce it to a single value.
</details>
<details>
<summary><code>Sort</code></summary>

```bash
// Sort the array
charArr.Sort()
```

Description: Sorts the characters of the array in ascending order.
</details>
<h2 align="center">Contributing</h2>
<p align="center">If you find a bug or have a feature request, please open an issue. Contributions are welcome!</p>
<h2 align="center">License</h2>
<p align="center">`xarray` is licensed under the MIT License. See <a href="https://github.com/diegoRrocha221/xarray-for-golang/blob/main/LICENSE">LICENSE</a> for details.</p>
