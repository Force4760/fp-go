package res

import "testing"

func TestOk(t *testing.T) {
	res := Ok(42)
	if res.isOk != true {
		t.Error("Ok should have returned a struct with isOk set to true. Received:", res.isOk)
	}
}

func TestErr(t *testing.T) {
	res := Err[int]("42")
	if res.isOk != false {
		t.Error("Err should have returned a struct with isOk set to false. Received:", res.isOk)
	}
}

func TestIsOk_Ok(t *testing.T) {
	res := IsOk(Ok(42))
	if res != true {
		t.Error("IsOk should have returned true. Received:", res)
	}
}
func TestIsOk_Err(t *testing.T) {
	res := IsOk(Err[int]("42"))
	if res != false {
		t.Error("IsOk should have returned false. Received:", res)
	}
}

func TestIsErr_Ok(t *testing.T) {
	res := IsErr(Ok(42))
	if res != false {
		t.Error("IsErr should have returned false. Received:", res)
	}
}
func TestIsErr_Err(t *testing.T) {
	res := IsErr(Err[int]("42"))
	if res != true {
		t.Error("IsErr should have returned true. Received:", res)
	}
}

func TestIsOkAnd_Ok_True(t *testing.T) {
	res := IsOkAnd(func(x int) bool { return x == 42 })(Ok(42))
	if res != true {
		t.Error("IsOkAnd should have returned true. Received:", res)
	}
}
func TestIsOkAnd_Ok_False(t *testing.T) {
	res := IsOkAnd(func(x int) bool { return x != 42 })(Ok(42))
	if res != false {
		t.Error("IsOkAnd should have returned false. Received:", res)
	}
}
func TestIsOkAnd_Err(t *testing.T) {
	res := IsOkAnd(func(x int) bool { return x == 42 })(Err[int]("42"))
	if res != false {
		t.Error("IsOkAnd should have returned false. Received:", res)
	}
}

func TestMap_Ok(t *testing.T) {
	res := Map(func(x int) int { return x + 1 })(Ok(41))
	if res.ok != 42 {
		t.Error("Map should have returned a struct with the ok value set to 42. Received:", res.ok)
	}
}
func TestMap_Err(t *testing.T) {
	res := Map(func(x int) int { return x + 1 })(Err[int]("42"))
	if res.err != "42" {
		t.Error("Map should have returned a struct with the err value set to \"42\". Received:", res.err)
	}
}

func TestFlat_Ok_Ok(t *testing.T) {
	res := Flat(Ok(Ok(42)))
	if res.ok != 42 {
		t.Error("Flat should have returned a struct with the ok value set to 42. Received:", res.ok)
	}
}
func TestFlat_Ok_Err(t *testing.T) {
	res := Flat(Ok(Err[int]("42")))
	if res.err != "42" {
		t.Error("Flat should have returned a struct with the err value set to \"42\". Received:", res.ok)
	}
}
func TestFlat_Err(t *testing.T) {
	res := Flat(Err[Result[int]]("42"))
	if res.err != "42" {
		t.Error("Flat should have returned a struct with the err value set to \"42\". Received:", res.ok)
	}
}

func TestGetOrElse_Ok(t *testing.T) {
	res := GetOrElse(21)(Ok(42))
	if res != 42 {
		t.Error("GetOrElse should have returned 42. Received:", res)
	}
}
func TestGetOrElse_Err(t *testing.T) {
	res := GetOrElse(42)(Err[int]("error"))
	if res != 42 {
		t.Error("GetOrElse should have returned 42. Received:", res)
	}
}

func TestEq_Ok_Ok_True(t *testing.T) {
	res := Eq(Ok(42))(Ok(42))
	if res != true {
		t.Error("Eq should have returned true. Received:", res)
	}
}
func TestEq_Ok_Ok_False(t *testing.T) {
	res := Eq(Ok(42))(Ok(21))
	if res != false {
		t.Error("Eq should have returned false. Received:", res)
	}
}
func TestEq_Ok_Err(t *testing.T) {
	res := Eq(Ok(42))(Err[int]("42"))
	if res != false {
		t.Error("Eq should have returned false. Received:", res)
	}
}
func TestEq_Err_Ok(t *testing.T) {
	res := Eq(Err[int]("42"))(Ok(42))
	if res != false {
		t.Error("Eq should have returned false. Received:", res)
	}
}
func TestEq_Err_Err(t *testing.T) {
	res := Eq(Err[int]("42"))(Err[int]("42"))
	if res != false {
		t.Error("Eq should have returned false. Received:", res)
	}
}

func TestChain_Ok(t *testing.T) {
	res := Chain(
		func(x int) Result[int] { return Ok(x + 1) },
	)(Ok(41))

	if res.ok != 42 {
		t.Error("Chain should have returned an Ok value with ok set to 42. Received:", res)
	}
}
func TestChain_Err(t *testing.T) {
	res := Chain(
		func(x int) Result[int] { return Ok(x + 1) },
	)(Err[int]("error"))

	if res.err != "error" {
		t.Error("Chain should have returned an Err value with err set to \"error\". Received:", res)
	}
}

func TestMatch_Ok(t *testing.T) {
	res := Match(
		func(s string) int { return 10 },
		func(i int) int { return i + 1 },
	)(Ok(41))

	if res != 42 {
		t.Error("Match should have returned 42. Received:", res)
	}
}
func TestMatch_Err(t *testing.T) {
	res := Match(
		func(s string) int { return 10 },
		func(i int) int { return i + 1 },
	)(Err[int]("42"))

	if res != 10 {
		t.Error("Match should have returned 10. Received:", res)
	}
}

func TestString_Ok(t *testing.T) {
	res := String(Ok(42))
	want := "Ok(42)"

	if res != want {
		t.Error("String should have returned", want, ". Received:", res)
	}
}
func TestString_Err(t *testing.T) {
	res := String(Err[int]("error"))
	want := "Err(error)"

	if res != want {
		t.Error("String should have returned", want, ". Received:", res)
	}
}

func TestError_Ok(t *testing.T) {
	res := Error(Ok(42))
	want := "No Error"

	if res != want {
		t.Error("Error should have returned", want, ". Received:", res)
	}
}
func TestError_Err(t *testing.T) {
	res := Error(Err[int]("error"))
	want := "error"

	if res != want {
		t.Error("Error should have returned", want, ". Received:", res)
	}
}
