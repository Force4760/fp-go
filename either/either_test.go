package either

import "testing"

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
