# fp-go

[![Go Reference](https://pkg.go.dev/badge/github.com/repeale/fp-go.svg)](https://pkg.go.dev/github.com/repeale/fp-go)
[![Go Report](https://goreportcard.com/badge/github.com/repeale/fp-go)](https://goreportcard.com/badge/github.com/repeale/fp-go)
[![codecov](https://codecov.io/gh/repeale/fp-go/branch/main/graph/badge.svg?token=ORP8NR634Q)](https://codecov.io/gh/repeale/fp-go)

Fp-go is a collection of Functional Programming helpers powered by Golang [1.18](https://tip.golang.org/doc/go1.18)+ [generics](https://tip.golang.org/doc/go1.18#generics).

<p align="center">
  <img 
    width="500"
    height="313"
    src="https://user-images.githubusercontent.com/9580458/162070974-4367f4b8-bb3d-451c-8114-dd578bad4e46.png"
  >
</p>

## Contents

- [Install](#install)
- [Features](#features)
  - [Currying](#currying)
  - [Variations](#variations)
- [Lists](#lists)
  - [Every](#every)
  - [Filter](#filter)
  - [Flat](#flat)
  - [FlatMap](#flatmap)
  - [Map](#map)
  - [Reduce](#reduce)
  - [Some](#some)
  - [Take](#take)
  - [TakeWhile](#takewhile)
  - [Drop](#drop)
  - [DropWhile](#dropwhile)
  - [Manipulation] (#manipulation)
  - [IsNull] (#isnull)
  - [And] (#and)
  - [Or] (#or)
  - [Elem] (#elem)
  - [Sum] (#sum)
  - [Prod] (#prod)
  - [Replicate] (#replicate)
  

  - [SplitAt] (#splitat) //////////
  - [Span] (#span) //////////
  - [Partition] (#elem) //////////
  - [Concat] (#sum) //////////
  - [ZipWith] (#zipwith) //////////
  - [Zip] (#zip) //////////
  - [Unzip] (#unzip) //////////
  

- [Helpers](#helpers)
  - [Compose](#compose)
  - [Pipe](#pipe)
  - [Curry](#curry)
  - [If](#if)

- [Options](#options)
- [Pairs](#pairs)

---

## Install

Requires go 1.18+

```sh
go get github.com/repeale/fp-go
```

---

## Features

- [Currying](#currying)
- [Variations](#variations)

### Currying

By default! Data last!

```go
func isPositive(x int) bool {
	return x > 0
}

func main() {
    filterPositive := fp.Filter(isPositive)
    numbers := []int{1, 2, 3, 4, 5}

    positives := filterPositive(numbers)
}
```

### Variations

Variations allows you to get different parameters in the callback function so that you get only only what is really needed.

[Default](#default) \
[WithIndex](#withindex) \
[WithSlice](#withslice)

#### Default

Only the current item is available:

```go
fp.Map[int, string](func(x int) { ... })
```

#### WithIndex

Current element and index are available:

```go
fp.MapWithIndex[int, string](func(x int, i int) { ... })
```

#### WithSlice

Current element, index and the whole slice are available:

```go
fp.MapWithSlice[int, string](func(x int, i int, xs: []int) { ... })
```

## Lists

### Every

Determines whether all the members of an array satisfy the specified test.

Variations `EveryWithIndex` and `EveryWithSlice`

```go
fp.Every(func(x int) bool { return x > 0 })([]int{1, 2, 3})

// => true
```

### Filter

Returns the elements of an array that meet the condition specified in a callback function.

Variations `FilterWithIndex` and `FilterWithSlice`

```go
fp.Filter(func(x int) bool { return x > 0 })([]int{-1, 2, -3, 4})

// => []int{2, 4}
```

### Flat

Returns a new array with all sub-array elements concatenated into it.

```go
fp.Flat([][]int{{1, 2}, {3, 4}})

// => []int{1, 2, 3, 4}
```

### FlatMap

Calls a defined callback function on each element of an array. Then, flattens the result into a new array. This is identical to a map followed by flat.

Variations `FlatMapWithIndex` and `FlatMapWithSlice`

```go
fp.FlatMap(func(x int) []int { return []int{x, x} })([]int{1, 2})

// => []int{1, 1, 2, 2}
```

### Map

Calls a defined callback function on each element of an array, and returns an array that contains the results.

Variations `MapWithIndex` and `MapWithSlice`

```go
fp.Map(func(x int64) string {
    return strconv.FormatInt(x, 10)
})([]int64{1, 2, 3})

// => []string{"1", "2", "3", "4"}
```

### Reduce

Calls the specified callback function for all the elements in an array. The return value of the callback function is the accumulated result, and is provided as an argument in the next call to the callback function.

Variations `ReduceWithIndex` and `ReduceWithSlice`

```go
fp.Reduce(func(acc int, curr int) int { return acc + curr }, 0)([]int{1, 2, 3})

// => 6
```

### Some

Determines whether the specified callback function returns true for any element of an array.

Variations `SomeWithIndex` and `SomeWithSlice`

```go
fp.Some(func(x int) bool { return x < 0 })([]int{1, 2, 3})

// => false
```

### Take

Returns the prefix of length n of an array, or the array itself if the length of the array is smaller than n.

```go
fp.Take[int](3)([]int{1, 2, 3, 4, 5})

// => []int{1, 2, 3}
```

### TakeWhile

Returns the longest prefix of an array such that every of those elements satisfy a predicate

Variations `TakeWhileIndex` and `TakeWhileWithSlice`

```go
fp.TakeWhile(func(x int) bool {return x < 5})([]int{1, 2, 3, 4, 5})

// => []int{1, 2, 3, 4}
```

### Drop

Returns the suffix of an array after the first n elements, or the empty array if the length of the array is smaller than n.

```go
fp.Drop[int](3)([]int{1, 2, 3, 4, 5})

// => []int{4, 5}
```

### DropWhile

Returns the suffix of an array that remains after `TakeWhile`

Variations `DropWhileIndex` and `DropWhileWithSlice`

```go
fp.DropWhile(func(x int) bool {return x < 5})([]int{1, 2, 3, 4, 5})

// => []int{5}
```

### Manipulation

```
┌head┐┌─────tail─────┐
[   1,   2,   3,   4 ]   
└─────init─────┘└last┘
```

```go
fp.Head([]int{1, 2, 3, 4, 5})  // => 1
fp.Tail([]int{1, 2, 3, 4, 5})  // => []int{2, 3, 4, 5}
fp.Last([]int{1, 2, 3, 4, 5})  // => 5
fp.Init([]int{1, 2, 3, 4, 5})  // => []int{1, 2, 3, 4}

fp.SafeHead([]int{1, 2, 3})    // => Some(1)
fp.SafeHead([]int{})           // => None()
fp.SafeLast([]int{1, 2, 3})    // => Some(3)
fp.SafeLast([]int{})           // => None()
```

### IsNull

Check if the list is empty

```go
fp.IsNull([]int{1, 2, 3, 4}) // => false
fp.IsNull([]int{})           // => true
```

### And

Check if all of the elements in a boolean array are true

```go
fp.And([]bool{true, true, true, true})  // => true
fp.And([]bool{true, true, true, false}) // => false
```

### Or

Check if any of the elements in a boolean array is true

```go
fp.And([]bool{false, false, true, false})  // => true
fp.And([]bool{false, false, false, false}) // => false
```

### Elem

Check if the provided value is an element of the array

```go
fp.Elem(42)([]int{1, 2, 3, 4, 5}) // => false
fp.Elem(42)([]int{3, 21, 42, 84}) // => true
```

### Sum

Computes the sum of all the numbers in the array

```go
fp.Sum([]int{21, 10, 11})

// => 42
```

### Prod

Computes the product of all the numbers in the array

```go
fp.Prod([]int{2, 3, 7})

// => 42
```

### Replicate

Create an array with a given value repeated n times

```go
fp.Replicate[bool](5)(true)

// => []bool{true, true, true, true, true}
```

---

## Options

Option represents encapsulation of an optional value, it might be used as the return type of functions which may or may not return a meaningful value when they are applied.
You could instanciate an `opt.Option[T]` with a value with `opt.Some(val)`. If the value is missing you can use `opt.None[T]()`.

Option exports:`Some`, `None`, `IsSome`, `IsSomeAnd`, `IsNone`,  `FromPtr`, `ToPtr`, `GetOrElse`, `Match`, `Map`, `Chain`, `Filter`, `Flat`, `Eq`.


## Pairs

Pair allows you to group 2 values into a single struct. Can be used to make a function with multiple returns return a sigle wrapped value

Pair exports `New`, `Fst`, `Snd`, `Get`, `MapFst`, `MapSnd`, `MapBoth`, `CheckFst`, `CheckSnd`, `CheckBoth`, `Merge`, `MergeC`, `Eq`, `Zip`.

---

## Helpers

### Compose

Performs right-to-left function composition.

Variations `Compose2` to `Compose16` stating the number of functions you are going to compose.

```go
func isPositive(x int) bool {
	return x > 0
}

func sumTwo(x int) int {
	return x + 2
}

Compose2(fp.Filter(isPositive), fp.Map(sumTwo))([]int{1, 2, 3, -1})

// => []int{3,4,5,1}
```

### Pipe

Performs left-to-right function composition.

Variations `Pipe2` to `Pipe16` stating the number of functions you are going to compose.

```go
func isPositive(x int) bool {
	return x > 0
}

func sumTwo(x int) int {
	return x + 2
}

Pipe2(fp.Filter(isPositive), fp.Map(sumTwo))([]int{1, 2, 3, -1})

// => []int{3,4,5}
```

### Curry

Allow to transfrom a function that receives a certain amount of params in single functions that take one single param each.

Variations `Curry2`, `Curry3` and `Curry4` stating the number of params will be curried individually.

```go
curryedSum := Curry2(func(a int, b int) int { return a + b })
curryedSum(1)(2)
```

### If

Implementation of the ternary operator.

```go
fp.If(true, 42, 21) // <=>   true ? 42 : 21   =>   42
fp.If(false, 2, 42) // <=>   false ? 2 : 42   =>   42
```

---

## TO DO

- [ ] Do Notation (?)
- [ ] Either
- [ ] Result
- [ ] Ranges
- [ ] Split at
- [ ] Span (takeWhile p, dropWhile p)
- [ ] Partition (filter p, filter $ not . p)
- [ ] Concat
- [ ] ZipWith
- [ ] Zip
- [ ] Unzip