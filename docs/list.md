# Lists


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
- [Manipulation](#manipulation)
- [IsNull](#isnull)
- [And](#and)
- [Or](#or)
- [Elem](#elem)
- [Sum](#sum)
- [Prod](#prod)
- [Replicate](#replicate)

- [SplitAt] (#splitat)
- [Span] (#span)
- [Partition] (#elem)
- [Concat] (#sum)
- [ZipWith] (#zipwith)
- [Zip] (#zip)
- [Unzip] (#unzip)


## Every

Determines whether all the members of an array satisfy the specified test.

Variations `EveryWithIndex` and `EveryWithSlice`

```go
fp.Every(func(x int) bool { return x > 0 })([]int{1, 2, 3})

// => true
```

## Filter

Returns the elements of an array that meet the condition specified in a callback function.

Variations `FilterWithIndex` and `FilterWithSlice`

```go
fp.Filter(func(x int) bool { return x > 0 })([]int{-1, 2, -3, 4})

// => []int{2, 4}
```

## Flat

Returns a new array with all sub-array elements concatenated into it.

```go
fp.Flat([][]int{{1, 2}, {3, 4}})

// => []int{1, 2, 3, 4}
```

## FlatMap

Calls a defined callback function on each element of an array. Then, flattens the result into a new array. This is identical to a map followed by flat.

Variations `FlatMapWithIndex` and `FlatMapWithSlice`

```go
fp.FlatMap(func(x int) []int { return []int{x, x} })([]int{1, 2})

// => []int{1, 1, 2, 2}
```

## Map

Calls a defined callback function on each element of an array, and returns an array that contains the results.

Variations `MapWithIndex` and `MapWithSlice`

```go
fp.Map(func(x int64) string {
    return strconv.FormatInt(x, 10)
})([]int64{1, 2, 3})

// => []string{"1", "2", "3", "4"}
```

## Reduce

Calls the specified callback function for all the elements in an array. The return value of the callback function is the accumulated result, and is provided as an argument in the next call to the callback function.

Variations `ReduceWithIndex` and `ReduceWithSlice`

```go
fp.Reduce(func(acc int, curr int) int { return acc + curr }, 0)([]int{1, 2, 3})

// => 6
```

## Some

Determines whether the specified callback function returns true for any element of an array.

Variations `SomeWithIndex` and `SomeWithSlice`

```go
fp.Some(func(x int) bool { return x < 0 })([]int{1, 2, 3})

// => false
```

## Take

Returns the prefix of length n of an array, or the array itself if the length of the array is smaller than n.

```go
fp.Take[int](3)([]int{1, 2, 3, 4, 5})

// => []int{1, 2, 3}
```

## TakeWhile

Returns the longest prefix of an array such that every of those elements satisfy a predicate

Variations `TakeWhileIndex` and `TakeWhileWithSlice`

```go
fp.TakeWhile(func(x int) bool {return x < 5})([]int{1, 2, 3, 4, 5})

// => []int{1, 2, 3, 4}
```

## Drop

Returns the suffix of an array after the first n elements, or the empty array if the length of the array is smaller than n.

```go
fp.Drop[int](3)([]int{1, 2, 3, 4, 5})

// => []int{4, 5}
```

## DropWhile

Returns the suffix of an array that remains after `TakeWhile`

Variations `DropWhileIndex` and `DropWhileWithSlice`

```go
fp.DropWhile(func(x int) bool {return x < 5})([]int{1, 2, 3, 4, 5})

// => []int{5}
```

## Manipulation

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

## IsNull

Check if the list is empty

```go
fp.IsNull([]int{1, 2, 3, 4}) // => false
fp.IsNull([]int{})           // => true
```

## And

Check if all of the elements in a boolean array are true

```go
fp.And([]bool{true, true, true, true})  // => true
fp.And([]bool{true, true, true, false}) // => false
```

## Or

Check if any of the elements in a boolean array is true

```go
fp.And([]bool{false, false, true, false})  // => true
fp.And([]bool{false, false, false, false}) // => false
```

## Elem

Check if the provided value is an element of the array

```go
fp.Elem(42)([]int{1, 2, 3, 4, 5}) // => false
fp.Elem(42)([]int{3, 21, 42, 84}) // => true
```

## Sum

Computes the sum of all the numbers in the array

```go
fp.Sum([]int{21, 10, 11})

// => 42
```

## Prod

Computes the product of all the numbers in the array

```go
fp.Prod([]int{2, 3, 7})

// => 42
```

## Replicate

Create an array with a given value repeated n times

```go
fp.Replicate[bool](5)(true)

// => []bool{true, true, true, true, true}
```