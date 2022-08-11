# Helpers


- [Compose](#compose)
- [Pipe](#pipe)
- [Curry](#curry)
- [If](#if)
- [Id](#id)


## Compose

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

## Pipe

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

## Curry

Allow to transfrom a function that receives a certain amount of params in single functions that take one single param each.

Variations `Curry2`, `Curry3` and `Curry4` stating the number of params will be curried individually.

```go
curryedSum := Curry2(func(a int, b int) int { return a + b })
curryedSum(1)(2)
```

## If

Implementation of the ternary operator.

```go
fp.If(true, 42, 21) // <=>   true ? 42 : 21   =>   42
fp.If(false, 2, 42) // <=>   false ? 2 : 42   =>   42
```

## Id

Id function. Takes an argument and returns it unchanged.

```go
fp.Id(42) // => 42
```