package lst

import (
	"reflect"
	"testing"

	"github.com/repeale/fp-go"
	"github.com/repeale/fp-go/pair"
)

func TestZip_EqLen(t *testing.T) {
	res := Zip[bool]([]int{1, 2, 3, 4, 5})([]bool{true, true, false, false, true})
	want := []pair.Pair[int, bool]{pair.New(1, true), pair.New(2, true), pair.New(3, false), pair.New(4, false), pair.New(5, true)}
	if !reflect.DeepEqual(res, want) {
		t.Error("Zip should have returned", want, ". Received:", res)
	}
}
func TestZip_Large_Fst(t *testing.T) {
	res := Zip[bool]([]int{1, 2, 3, 4, 5})([]bool{true, true, false})
	want := []pair.Pair[int, bool]{pair.New(1, true), pair.New(2, true), pair.New(3, false)}
	if !reflect.DeepEqual(res, want) {
		t.Error("Zip should have returned", want, ". Received:", res)
	}
}
func TestZip_Large_Snd(t *testing.T) {
	res := Zip[bool]([]int{1, 2, 3})([]bool{true, true, false, false, true})
	want := []pair.Pair[int, bool]{pair.New(1, true), pair.New(2, true), pair.New(3, false)}
	if !reflect.DeepEqual(res, want) {
		t.Error("Zip should have returned", want, ". Received:", res)
	}
}

func TestZipWith_EqLen(t *testing.T) {
	res := ZipWith(
		func(a int, b bool) int { return fp.If(b, a+1, a-1) },
	)([]int{1, 2, 3, 4, 5})([]bool{true, true, false, false, true})

	want := []int{2, 3, 2, 3, 6}

	if !reflect.DeepEqual(res, want) {
		t.Error("ZipWith should have returned", want, ". Received:", res)
	}
}
func TestZipWith_Large_Fst(t *testing.T) {
	res := ZipWith(
		func(a int, b bool) int { return fp.If(b, a+1, a-1) },
	)([]int{1, 2, 3, 4, 5})([]bool{true, true, false})

	want := []int{2, 3, 2}

	if !reflect.DeepEqual(res, want) {
		t.Error("ZipWith should have returned", want, ". Received:", res)
	}
}
func TestZipWith_Large_Snd(t *testing.T) {
	res := ZipWith(
		func(a int, b bool) int { return fp.If(b, a+1, a-1) },
	)([]int{1, 2, 3})([]bool{true, true, false, false, true})

	want := []int{2, 3, 2}

	if !reflect.DeepEqual(res, want) {
		t.Error("Zip should have returned", want, ". Received:", res)
	}
}

func TestUnzip(t *testing.T) {
	res := Unzip([]pair.Pair[int, bool]{
		pair.New(1, true), pair.New(2, false), pair.New(3, true), pair.New(4, false),
	})

	want := pair.New([]int{1, 2, 3, 4}, []bool{true, false, true, false})

	if !reflect.DeepEqual(res, want) {
		t.Error("ZipWith should have returned", want, ". Received:", res)
	}
}
