# fp-go

[![Go Reference](https://pkg.go.dev/badge/github.com/repeale/fp-go.svg)](https://pkg.go.dev/github.com/repeale/fp-go)
[![Go Report](https://goreportcard.com/badge/github.com/repeale/fp-go)](https://goreportcard.com/badge/github.com/repeale/fp-go)
[![codecov](https://codecov.io/gh/repeale/fp-go/branch/main/graph/badge.svg?token=ORP8NR634Q)](https://codecov.io/gh/repeale/fp-go)

Fp-go is a collection of Functional Programming helpers powered by Golang [1.18](https://tip.golang.org/doc/go1.18)+ [generics](https://tip.golang.org/doc/go1.18#generics) created by [Repeale](https://github.com/repeale)

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

- [Helpers](./docs/helpers.md)

- Types
  - [Lists](./docs/list.md)
  - [Option](./docs/option.md)
  - [Pair](./docs/pair.md)
  - [Either](./docs/either.md)
  - [Result](./docs/result.md)

---

## Install

Requires go 1.18+

```sh
go get github.com/force4760/fp-go
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

---

## TO DO

- [X] Either
- [X] Result
- [ ] Ranges