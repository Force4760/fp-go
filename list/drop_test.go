package lst

import (
	"reflect"
	"testing"
)

func TestDrop(t *testing.T) {
	res := Drop[int](3)([]int{1, 2, 3, 4, 5})
	want := []int{4, 5}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Drop should have returned", want, ". Received:", res)
	}
}
func TestDrop_Eq(t *testing.T) {
	res := Drop[int](5)([]int{1, 2, 3, 4, 5})
	want := []int{}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Drop should have returned", want, ". Received:", res)
	}
}
func TestDrop_More(t *testing.T) {
	res := Drop[int](10)([]int{1, 2, 3, 4, 5})
	want := []int{}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Drop should have returned", want, ". Received:", res)
	}
}

func TestDropWhile(t *testing.T) {
	res := DropWhile(func(x int) bool { return x < 4 })([]int{1, 2, 3, 4, 5})
	want := []int{4, 5}
	if reflect.DeepEqual(res, want) != true {
		t.Error("DropWhile should have returned", want, ". Received:", res)
	}
}
func TestDropWhile_All(t *testing.T) {
	res := DropWhile(func(x int) bool { return x < 10 })([]int{1, 2, 3, 4, 5})
	want := []int{}
	if reflect.DeepEqual(res, want) != true {
		t.Error("DropWhile should have returned", want, ". Received:", res)
	}
}
func TestDropWhile_None(t *testing.T) {
	res := DropWhile(func(x int) bool { return x > 10 })([]int{1, 2, 3, 4, 5})
	want := []int{1, 2, 3, 4, 5}
	if reflect.DeepEqual(res, want) != true {
		t.Error("DropWhile should have returned", want, ". Received:", res)
	}
}

func TestDropWhileWithIndex(t *testing.T) {
	res := DropWhileWithIndex(
		func(x int, i int) bool { return x+i < 4 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{3, 4, 5}

	if reflect.DeepEqual(res, want) != true {
		t.Error("DropWhileWithIndex should have returned", want, ". Received:", res)
	}
}
func TestDropWhileWithIndex_All(t *testing.T) {
	res := DropWhileWithIndex(
		func(x int, i int) bool { return x+1 < 10 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{}

	if reflect.DeepEqual(res, want) != true {
		t.Error("DropWhileWithIndex should have returned", want, ". Received:", res)
	}
}
func TestDropWhileWithIndex_None(t *testing.T) {
	res := DropWhileWithIndex(
		func(x int, i int) bool { return x+i > 10 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{1, 2, 3, 4, 5}

	if reflect.DeepEqual(res, want) != true {
		t.Error("DropWhileWithIndex should have returned", want, ". Received:", res)
	}
}

func TestDropWhileWithSlice(t *testing.T) {
	res := DropWhileWithSlice(
		func(x int, i int, xs []int) bool { return x+i < 4 && len(xs) == 5 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{3, 4, 5}

	if reflect.DeepEqual(res, want) != true {
		t.Error("DropWhileWithIndex should have returned", want, ". Received:", res)
	}
}
func TestDropWhileWithSlice_All(t *testing.T) {
	res := DropWhileWithSlice(
		func(x int, i int, xs []int) bool { return x+1 < 10 && len(xs) == 5 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{}

	if reflect.DeepEqual(res, want) != true {
		t.Error("DropWhileWithIndex should have returned", want, ". Received:", res)
	}
}
func TestDropWhileWithSlice_None(t *testing.T) {
	res := DropWhileWithSlice(
		func(x int, i int, xs []int) bool { return x+i < 10 && len(xs) != 5 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{1, 2, 3, 4, 5}

	if reflect.DeepEqual(res, want) != true {
		t.Error("DropWhileWithIndex should have returned", want, ". Received:", res)
	}
}
