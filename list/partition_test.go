package lst

import (
	"reflect"
	"testing"

	"github.com/force4760/fp-go/pair"
)

func TestPartition(t *testing.T) {
	res := Partition(func(x int) bool { return x >= 4 })([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{4, 5}, []int{1, 2, 3})

	if reflect.DeepEqual(res, want) != true {
		t.Error("Partition should have returned", want, ". Received:", res)
	}
}
func TestPartition_All(t *testing.T) {
	res := Partition(func(x int) bool { return x < 10 })([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{1, 2, 3, 4, 5}, []int{})

	if reflect.DeepEqual(res, want) != true {
		t.Error("Partition should have returned", want, ". Received:", res)
	}
}
func TestPartition_None(t *testing.T) {
	res := Partition(func(x int) bool { return x > 10 })([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{}, []int{1, 2, 3, 4, 5})

	if reflect.DeepEqual(res, want) != true {
		t.Error("Partition should have returned", want, ". Received:", res)
	}
}

func TestPartitionWithIndex(t *testing.T) {
	res := PartitionWithIndex(
		func(x int, i int) bool { return x+i >= 4 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{3, 4, 5}, []int{1, 2})

	if reflect.DeepEqual(res, want) != true {
		t.Error("PartitionWithIndex should have returned", want, ". Received:", res)
	}
}
func TestPartitionWithIndex_All(t *testing.T) {
	res := PartitionWithIndex(
		func(x int, i int) bool { return x+1 < 10 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{1, 2, 3, 4, 5}, []int{})

	if reflect.DeepEqual(res, want) != true {
		t.Error("PartitionWithIndex should have returned", want, ". Received:", res)
	}
}
func TestPartitionWithIndex_None(t *testing.T) {
	res := PartitionWithIndex(
		func(x int, i int) bool { return x+i > 10 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{}, []int{1, 2, 3, 4, 5})

	if reflect.DeepEqual(res, want) != true {
		t.Error("PartitionWithIndex should have returned", want, ". Received:", res)
	}
}

func TestPartitionWithSlice(t *testing.T) {
	res := PartitionWithSlice(
		func(x int, i int, xs []int) bool { return x+i >= 4 && len(xs) == 5 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{3, 4, 5}, []int{1, 2})

	if reflect.DeepEqual(res, want) != true {
		t.Error("PartitionWithSlice should have returned", want, ". Received:", res)
	}
}
func TestPartitionWithSlice_All(t *testing.T) {
	res := PartitionWithSlice(
		func(x int, i int, xs []int) bool { return x+1 < 10 && len(xs) == 5 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{1, 2, 3, 4, 5}, []int{})

	if reflect.DeepEqual(res, want) != true {
		t.Error("PartitionWithSlice should have returned", want, ". Received:", res)
	}
}
func TestPartitionWithSlice_None(t *testing.T) {
	res := PartitionWithSlice(
		func(x int, i int, xs []int) bool { return x+i < 10 && len(xs) != 5 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{}, []int{1, 2, 3, 4, 5})

	if reflect.DeepEqual(res, want) != true {
		t.Error("PartitionWithSlice should have returned", want, ". Received:", res)
	}
}
