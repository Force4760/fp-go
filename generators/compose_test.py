def compose_test(n: int) -> str:
    fns = ", ".join(f"fn{i % 2}" for i in range(n))

	expect = 0
	for i in range(n-1, -1, -1):
		if i % 2 == 0:
			expect += 1
		else:
			expect *= 2

    return f"""func TestCompose{n}_Example(t *testing.T) {{
    fn0 := func(x int) int {{ return x + 1 }}
	fn1 := func(x int) int {{ return x * 2 }}
    res := Compose{n}({fns})(0)
    if res != {expect} {{
        t.Error("Should perform right-to-left function composition of {n} functions. Received:", res)
    }}
}}"""