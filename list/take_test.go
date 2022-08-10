package lst

import (
	"reflect"
	"testing"
)

func TestTake(t *testing.T) {
	res := Take[int](3)([]int{1, 2, 3, 4, 5})
	want := []int{1, 2, 3}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Take should have returned", want, ". Received:", res)
	}
}
func TestTake_Eq(t *testing.T) {
	res := Take[int](5)([]int{1, 2, 3, 4, 5})
	want := []int{1, 2, 3, 4, 5}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Take should have returned", want, ". Received:", res)
	}
}
func TestTake_More(t *testing.T) {
	res := Take[int](10)([]int{1, 2, 3, 4, 5})
	want := []int{1, 2, 3, 4, 5}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Take should have returned", want, ". Received:", res)
	}
}

func TestTakeWhile(t *testing.T) {
	res := TakeWhile(func(x int) bool { return x < 4 })([]int{1, 2, 3, 4, 5})
	want := []int{1, 2, 3}
	if reflect.DeepEqual(res, want) != true {
		t.Error("TakeWhile should have returned", want, ". Received:", res)
	}
}
func TestTakeWhile_All(t *testing.T) {
	res := TakeWhile(func(x int) bool { return x < 10 })([]int{1, 2, 3, 4, 5})
	want := []int{1, 2, 3, 4, 5}
	if reflect.DeepEqual(res, want) != true {
		t.Error("TakeWhile should have returned", want, ". Received:", res)
	}
}
func TestTakeWhile_None(t *testing.T) {
	res := TakeWhile(func(x int) bool { return x > 10 })([]int{1, 2, 3, 4, 5})
	want := []int{}
	if reflect.DeepEqual(res, want) != true {
		t.Error("TakeWhile should have returned", want, ". Received:", res)
	}
}

func TestTakeWhileWithIndex(t *testing.T) {
	res := TakeWhileWithIndex(
		func(x int, i int) bool { return x+i < 4 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{1, 2}

	if reflect.DeepEqual(res, want) != true {
		t.Error("TakeWhileWithIndex should have returned", want, ". Received:", res)
	}
}
func TestTakeWhileWithIndex_All(t *testing.T) {
	res := TakeWhileWithIndex(
		func(x int, i int) bool { return x+1 < 10 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{1, 2, 3, 4, 5}

	if reflect.DeepEqual(res, want) != true {
		t.Error("TakeWhileWithIndex should have returned", want, ". Received:", res)
	}
}
func TestTakeWhileWithIndex_None(t *testing.T) {
	res := TakeWhileWithIndex(
		func(x int, i int) bool { return x+i > 10 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{}

	if reflect.DeepEqual(res, want) != true {
		t.Error("TakeWhileWithIndex should have returned", want, ". Received:", res)
	}
}

func TestTakeWhileWithSlice(t *testing.T) {
	res := TakeWhileWithSlice(
		func(x int, i int, xs []int) bool { return x+i < 4 && len(xs) == 5 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{1, 2}

	if reflect.DeepEqual(res, want) != true {
		t.Error("TakeWhileWithIndex should have returned", want, ". Received:", res)
	}
}
func TestTakeWhileWithSlice_All(t *testing.T) {
	res := TakeWhileWithSlice(
		func(x int, i int, xs []int) bool { return x+1 < 10 && len(xs) == 5 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{1, 2, 3, 4, 5}

	if reflect.DeepEqual(res, want) != true {
		t.Error("TakeWhileWithIndex should have returned", want, ". Received:", res)
	}
}
func TestTakeWhileWithSlice_None(t *testing.T) {
	res := TakeWhileWithSlice(
		func(x int, i int, xs []int) bool { return x+i < 10 && len(xs) != 5 },
	)([]int{1, 2, 3, 4, 5})

	want := []int{}

	if reflect.DeepEqual(res, want) != true {
		t.Error("TakeWhileWithIndex should have returned", want, ". Received:", res)
	}
}
