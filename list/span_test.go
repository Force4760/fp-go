package lst

import (
	"reflect"
	"testing"

	"github.com/repeale/fp-go/pair"
)

func TestSpan(t *testing.T) {
	res := Span(func(x int) bool { return x < 4 })([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{1, 2, 3}, []int{4, 5})

	if reflect.DeepEqual(res, want) != true {
		t.Error("Span should have returned", want, ". Received:", res)
	}
}
func TestSpan_All(t *testing.T) {
	res := Span(func(x int) bool { return x < 10 })([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{1, 2, 3, 4, 5}, []int{})

	if reflect.DeepEqual(res, want) != true {
		t.Error("Span should have returned", want, ". Received:", res)
	}
}
func TestSpan_None(t *testing.T) {
	res := Span(func(x int) bool { return x > 10 })([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{}, []int{1, 2, 3, 4, 5})

	if reflect.DeepEqual(res, want) != true {
		t.Error("Span should have returned", want, ". Received:", res)
	}
}

func TestSpanWithIndex(t *testing.T) {
	res := SpanWithIndex(
		func(x int, i int) bool { return x+i < 4 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{1, 2}, []int{3, 4, 5})

	if reflect.DeepEqual(res, want) != true {
		t.Error("SpanWithIndex should have returned", want, ". Received:", res)
	}
}
func TestSpanWithIndex_All(t *testing.T) {
	res := SpanWithIndex(
		func(x int, i int) bool { return x+1 < 10 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{1, 2, 3, 4, 5}, []int{})

	if reflect.DeepEqual(res, want) != true {
		t.Error("SpanWithIndex should have returned", want, ". Received:", res)
	}
}
func TestSpanWithIndex_None(t *testing.T) {
	res := SpanWithIndex(
		func(x int, i int) bool { return x+i > 10 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{}, []int{1, 2, 3, 4, 5})

	if reflect.DeepEqual(res, want) != true {
		t.Error("SpanWithIndex should have returned", want, ". Received:", res)
	}
}

func TestSpanWithSlice(t *testing.T) {
	res := SpanWithSlice(
		func(x int, i int, xs []int) bool { return x+i < 4 && len(xs) == 5 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{1, 2}, []int{3, 4, 5})

	if reflect.DeepEqual(res, want) != true {
		t.Error("SpanWithSlice should have returned", want, ". Received:", res)
	}
}
func TestSpanWithSlice_All(t *testing.T) {
	res := SpanWithSlice(
		func(x int, i int, xs []int) bool { return x+1 < 10 && len(xs) == 5 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{1, 2, 3, 4, 5}, []int{})

	if reflect.DeepEqual(res, want) != true {
		t.Error("SpanWithSlice should have returned", want, ". Received:", res)
	}
}
func TestSpanWithSlice_None(t *testing.T) {
	res := SpanWithSlice(
		func(x int, i int, xs []int) bool { return x+i < 10 && len(xs) != 5 },
	)([]int{1, 2, 3, 4, 5})

	want := pair.New([]int{}, []int{1, 2, 3, 4, 5})

	if reflect.DeepEqual(res, want) != true {
		t.Error("SpanWithSlice should have returned", want, ". Received:", res)
	}
}
