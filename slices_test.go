package fp

import (
	"reflect"
	"testing"

	opt "github.com/repeale/fp-go/option"
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

func TestIsNull_True(t *testing.T) {
	res := IsNull([]int{})
	if res != true {
		t.Error("IsNull should have returned true. Received:", res)
	}
}
func TestIsNull_False(t *testing.T) {
	res := IsNull([]int{1})
	if res != false {
		t.Error("IsNull should have returned false. Received:", res)
	}
}

func TestAnd_AllTrue(t *testing.T) {
	res := And([]bool{true, true, true, true})
	if res != true {
		t.Error("And should have returned true. Received:", res)
	}
}
func TestAnd_AllFalse(t *testing.T) {
	res := And([]bool{false, false, false, false})
	if res != false {
		t.Error("And should have returned false. Received:", res)
	}
}
func TestAnd_OneTrue(t *testing.T) {
	res := And([]bool{false, true, false, false})
	if res != false {
		t.Error("And should have returned false. Received:", res)
	}
}
func TestAnd_OneFalse(t *testing.T) {
	res := And([]bool{true, true, false, true})
	if res != false {
		t.Error("And should have returned false. Received:", res)
	}
}
func TestAnd_Empty(t *testing.T) {
	res := And([]bool{})
	if res != true {
		t.Error("And should have returned true. Received:", res)
	}
}

func TestOr_AllTrue(t *testing.T) {
	res := Or([]bool{true, true, true, true})
	if res != true {
		t.Error("Or should have returned true. Received:", res)
	}
}
func TestOr_AllFalse(t *testing.T) {
	res := Or([]bool{false, false, false, false})
	if res != false {
		t.Error("Or should have returned false. Received:", res)
	}
}
func TestOr_OneTrue(t *testing.T) {
	res := Or([]bool{false, true, false, false})
	if res != true {
		t.Error("Or should have returned true. Received:", res)
	}
}
func TestOr_OneFalse(t *testing.T) {
	res := Or([]bool{true, true, false, true})
	if res != true {
		t.Error("Or should have returned true. Received:", res)
	}
}
func TestOr_Empty(t *testing.T) {
	res := Or([]bool{})
	if res != false {
		t.Error("Or should have returned false. Received:", res)
	}
}

func TestElem_True(t *testing.T) {
	res := Elem(3)([]int{1, 2, 3, 4, 5})
	if res != true {
		t.Error("Elem should have returned true. Received:", res)
	}
}
func TestElem_False(t *testing.T) {
	res := Elem(9)([]int{1, 2, 3, 4, 5})
	if res != false {
		t.Error("Elem should have returned false. Received:", res)
	}
}
func TestElem_Empty(t *testing.T) {
	res := Elem(3)([]int{})
	if res != false {
		t.Error("Elem should have returned false. Received:", res)
	}
}

func TestSum(t *testing.T) {
	res := Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if res != 45 {
		t.Error("Sum should have returned 45. Received:", res)
	}
}
func TestSum_Empty(t *testing.T) {
	res := Sum([]int{})
	if res != 0 {
		t.Error("Sum should have returned 0. Received:", res)
	}
}

func TestProd(t *testing.T) {
	res := Prod([]int{1, 2, 3, 4, 5})
	if res != 120 {
		t.Error("Prod should have returned 120. Received:", res)
	}
}
func TestProd_Empty(t *testing.T) {
	res := Prod([]int{})
	if res != 1 {
		t.Error("Prod should have returned 1. Received:", res)
	}
}

func TestReplicate(t *testing.T) {
	res := Replicate[int](5)(1)
	want := []int{1, 1, 1, 1, 1}
	if reflect.DeepEqual(res, want) != true {
		t.Error("Replicate should have returned", want, ". Received:", res)
	}
}

func TestSafeHead(t *testing.T) {
	res := SafeHead([]int{1, 2})
	if !opt.Eq(res)(opt.Some(1)) {
		t.Error("SafeHead should have returned the first element of the slice as Some(1). Received:", res)
	}
}
func TestSafeHead_Single(t *testing.T) {
	res := SafeHead([]int{1, 2})
	if !opt.Eq(res)(opt.Some(1)) {
		t.Error("SafeHead should have returned the first element of the slice as Some(1). Received:", res)
	}
}
func TestSafeHead_Empty(t *testing.T) {
	res := SafeHead([]int{})
	if !opt.Eq(res)(opt.None[int]()) {
		t.Error("SafeHead should have returned None. Received:", res)
	}
}

func TestSafeLast(t *testing.T) {
	res := SafeLast([]int{1, 2})
	if !opt.Eq(res)(opt.Some(2)) {
		t.Error("SafeLast should have returned the last element of the slice as Some(1). Received:", res)
	}
}
func TestSafeLast_Single(t *testing.T) {
	res := SafeLast([]int{1})
	if !opt.Eq(res)(opt.Some(1)) {
		t.Error("SafeHead should have returned the last element of the slice as Some(1). Received:", res)
	}
}
func TestSafeLast_Empty(t *testing.T) {
	res := SafeLast([]int{})
	if !opt.Eq(res)(opt.None[int]()) {
		t.Error("SafeLast should have returned None. Received:", res)
	}
}
