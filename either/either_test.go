package either

import (
	"testing"
)

func TestLeft(t *testing.T) {
	res := Left[int](42)
	if res.isLeft != true {
		t.Error("Left should return a struct with isLeft set to true. Received:", res.isLeft)
	}
}

func TestRight(t *testing.T) {
	res := Right[int](42)
	if res.isLeft != false {
		t.Error("Right should return a struct with isLeft set to false. Received:", res.isLeft)
	}
}

func TestIsLeft_True(t *testing.T) {
	res := IsLeft(Left[int](42))
	if res != true {
		t.Error("IsLeft should return true. Received:", res)
	}
}
func TestIsLeft_False(t *testing.T) {
	res := IsLeft(Right[int](42))
	if res != false {
		t.Error("IsLeft should return false. Received:", res)
	}
}

func TestIsRight_True(t *testing.T) {
	res := IsRight(Right[int](42))
	if res != true {
		t.Error("IsRight should return true. Received:", res)
	}
}
func TestIsRight_False(t *testing.T) {
	res := IsRight(Left[int](42))
	if res != false {
		t.Error("IsRight should return false. Received:", res)
	}
}

func TestIsLeftAnd_True_True(t *testing.T) {
	res := IsLeftAnd[int](
		func(x int) bool { return x == 42 },
	)(Left[int](42))

	if res != true {
		t.Error("IsLeftAnd should return true. Received:", res)
	}
}
func TestIsLeftAnd_True_False(t *testing.T) {
	res := IsLeftAnd[int](
		func(x int) bool { return x != 42 },
	)(Left[int](42))

	if res != false {
		t.Error("IsLeftAnd should return false. Received:", res)
	}
}
func TestIsLeftAnd_False(t *testing.T) {
	res := IsLeftAnd[int](
		func(x int) bool { return x == 42 },
	)(Right[int](42))

	if res != false {
		t.Error("IsLeftAnd should return false. Received:", res)
	}
}

func TestIsRightAnd_True_True(t *testing.T) {
	res := IsRightAnd[int](
		func(x int) bool { return x == 42 },
	)(Right[int](42))

	if res != true {
		t.Error("IsRightAnd should return true. Received:", res)
	}
}
func TestIsRightAnd_True_False(t *testing.T) {
	res := IsRightAnd[int](
		func(x int) bool { return x != 42 },
	)(Right[int](42))

	if res != false {
		t.Error("IsRightAnd should return false. Received:", res)
	}
}
func TestIsRightAnd_False(t *testing.T) {
	res := IsRightAnd[int](
		func(x int) bool { return x == 42 },
	)(Left[int](42))

	if res != false {
		t.Error("IsRightAnd should return false. Received:", res)
	}
}

func TestEq_Left_Left_True(t *testing.T) {
	res := Eq(Left[int](42))(Left[int](42))
	if res != true {
		t.Error("Eq should return true. Received:", res)
	}
}
func TestEq_Left_Left_False(t *testing.T) {
	res := Eq(Left[int](42))(Left[int](21))
	if res != false {
		t.Error("Eq should return false. Received:", res)
	}
}
func TestEq_Right_Right_True(t *testing.T) {
	res := Eq(Right[int](42))(Right[int](42))
	if res != true {
		t.Error("Eq should return true. Received:", res)
	}
}
func TestEq_Right_Right_False(t *testing.T) {
	res := Eq(Right[int](42))(Right[int](21))
	if res != false {
		t.Error("Eq should return false. Received:", res)
	}
}
func TestEq_Right_Left(t *testing.T) {
	res := Eq(Right[int](42))(Left[int](42))
	if res != false {
		t.Error("Eq should return false. Received:", res)
	}
}
func TestEq_Left_Right(t *testing.T) {
	res := Eq(Left[int](42))(Right[int](42))
	if res != false {
		t.Error("Eq should return false. Received:", res)
	}
}

func TestFlip_Left(t *testing.T) {
	res := Flip(Left[int](42))
	if res.isLeft != false {
		t.Error("Flip should a struct with isLeft set to false. Received:", res.isLeft)
	}
}
func TestFlip_Right(t *testing.T) {
	res := Flip(Right[int](42))
	if res.isLeft != true {
		t.Error("Flip should a struct with isLeft set to true. Received:", res.isLeft)
	}
}

func TestMapL_Left(t *testing.T) {
	res := MapL[int](
		func(x int) int { return x + 1 },
	)(Left[int](41))

	if res.left != 42 {
		t.Error("MapL should have returned 42. Received:", res.left)
	}
}
func TestMapL_Right(t *testing.T) {
	res := MapL[int](
		func(x int) int { return x + 1 },
	)(Right[int](42))

	if res.right != 42 {
		t.Error("MapL should have returned 42. Received:", res.right)
	}
}

func TestMapR_Left(t *testing.T) {
	res := MapR[int](
		func(x int) int { return x + 1 },
	)(Left[int](42))

	if res.left != 42 {
		t.Error("MapR should have returned 42. Received:", res.left)
	}
}
func TestMapR_Right(t *testing.T) {
	res := MapR[int](
		func(x int) int { return x + 1 },
	)(Right[int](41))

	if res.right != 42 {
		t.Error("MapL should have returned 42. Received:", res.right)
	}
}

func TestGetOrElseL_Left(t *testing.T) {
	res := GetOrElseL[int](21)(Left[int](42))

	if res != 42 {
		t.Error("GetOrElseL should have returned 42. Received:", res)
	}
}
func TestGetOrElseL_Right(t *testing.T) {
	res := GetOrElseL[int](42)(Right[int](21))

	if res != 42 {
		t.Error("GetOrElseL should have returned 42. Received:", res)
	}
}

func TestGetOrElseR_Left(t *testing.T) {
	res := GetOrElseR[int](42)(Left[int](21))

	if res != 42 {
		t.Error("GetOrElseR should have returned 42. Received:", res)
	}
}
func TestGetOrElseR_Right(t *testing.T) {
	res := GetOrElseR[int](21)(Right[int](42))

	if res != 42 {
		t.Error("GetOrElseR should have returned 42. Received:", res)
	}
}

func TestMatch_Left(t *testing.T) {
	res := Match(
		func(x int) int { return x + 1 },
		func(x int) int { return x - 1 },
	)(Left[int](41))

	if res != 42 {
		t.Error("Match should have returned 42. Received:", res)
	}
}
func TestMatch_Right(t *testing.T) {
	res := Match(
		func(x int) int { return x + 1 },
		func(x int) int { return x - 1 },
	)(Right[int](43))

	if res != 42 {
		t.Error("Match should have returned 42. Received:", res)
	}
}

func TestFlatL_Left_Left(t *testing.T) {
	res := FlatL(Left[int](Left[int](42)))

	if res.left != 42 {
		t.Error("FlatL should have returned a left value with 42. Received:", res.left)
	}
}
func TestFlatL_Left_Right(t *testing.T) {
	res := FlatL(Left[int](Right[int](42)))

	if res.right != 42 {
		t.Error("FlatL should have returned a right value with 42. Received:", res.right)
	}
}
func TestFlatL_Right(t *testing.T) {
	res := FlatL(Right[Either[int, int]](42))

	if res.right != 42 {
		t.Error("FlatL should have returned a right value with 42. Received:", res.right)
	}
}

func TestFlatR_Right_Left(t *testing.T) {
	res := FlatR(Right[int](Left[int](42)))

	if res.left != 42 {
		t.Error("FlatR should have returned a left value with 42. Received:", res.left)
	}
}
func TestFlatR_Right_Right(t *testing.T) {
	res := FlatR(Right[int](Right[int](42)))

	if res.right != 42 {
		t.Error("FlatR should have returned a right value with 42. Received:", res.right)
	}
}
func TestFlatR_Left(t *testing.T) {
	res := FlatR(Left[Either[int, int]](42))

	if res.left != 42 {
		t.Error("FlatR should have returned a left value with 42. Received:", res.left)
	}
}

func TestChainL_Left(t *testing.T) {
	res := ChainL(
		func(x int) Either[int, int] { return Left[int](x + 1) },
	)(Left[int](41))

	if res.left != 42 {
		t.Error("ChainL should have returned a left value with 42. Received:", res.left)
	}
}
func TestChainL_Right(t *testing.T) {
	res := ChainL(
		func(x int) Either[int, int] { return Left[int](x + 1) },
	)(Right[int](42))

	if res.right != 42 {
		t.Error("ChainL should have returned a right value with 42. Received:", res.right)
	}
}

func TestChainR_Left(t *testing.T) {
	res := ChainR(
		func(x int) Either[int, int] { return Right[int](x + 1) },
	)(Left[int](42))

	if res.left != 42 {
		t.Error("ChainR should have returned a left value with 42. Received:", res.left)
	}
}
func TestChainR_Right(t *testing.T) {
	res := ChainR(
		func(x int) Either[int, int] { return Right[int](x + 1) },
	)(Right[int](41))

	if res.right != 42 {
		t.Error("ChainR should have returned a right value with 42. Received:", res.right)
	}
}

func TestString_Left(t *testing.T) {
	res := String(Left[int](42))
	want := "Left(42)"

	if res != want {
		t.Error("String should have returned", want, ". Received:", res)
	}
}

func TestString_Right(t *testing.T) {
	res := String(Right[int](42))
	want := "Right(42)"

	if res != want {
		t.Error("String should have returned", want, ". Received:", res)
	}
}

// Test Chain
