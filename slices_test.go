package fp

import (
	"reflect"
	"testing"
)

func TestHead(t *testing.T) {
	res := Head([]int{1, 2})
	if res != 1 {
		t.Error("Head should have returned the first element of the slice (1). Received:", res)
	}
}
func TestHead_Single(t *testing.T) {
	res := Head([]int{1})
	if res != 1 {
		t.Error("Head should have returned the first element of the slice (1). Received:", res)
	}
}

func TestTail(t *testing.T) {
	res := Tail([]int{1, 2, 3, 4, 5})
	want := []int{2, 3, 4, 5}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Tail should have returned the slice without the first element", want, ". Received:", res)
	}
}
func TestTail_Empty(t *testing.T) {
	res := Tail([]int{})
	want := []int{}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Tail should have returned the empty slice. Received:", res)
	}
}
func TestTail_Single(t *testing.T) {
	res := Tail([]int{1})
	want := []int{}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Tail should have returned the slice without the first element", want, ". Received:", res)
	}
}

func TestLast(t *testing.T) {
	res := Last([]int{1, 2})
	if res != 2 {
		t.Error("Last should have returned the first element of the slice (2). Received:", res)
	}
}
func TestLast_Single(t *testing.T) {
	res := Last([]int{2})
	if res != 2 {
		t.Error("Last should have returned the first element of the slice (2). Received:", res)
	}
}

func TestInit(t *testing.T) {
	res := Init([]int{1, 2, 3, 4, 5})
	want := []int{1, 2, 3, 4}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Init should have returned the slice without the last element", want, ". Received:", res)
	}
}
func TestInit_Empty(t *testing.T) {
	res := Init([]int{})
	want := []int{}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Init should have returned the slice without the last element", want, ". Received:", res)
	}
}
func TestInit_Single(t *testing.T) {
	res := Init([]int{1})
	want := []int{}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Init should have returned the slice without the last element", want, ". Received:", res)
	}
}
