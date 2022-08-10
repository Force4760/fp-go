def compose_func(n: int) -> str:
    gen = "[" + ", ".join(f"T{i+1}" for i in range(n)) + ", R any]"

    fns = f"fn{n} func(T{n}) R, " + ", ".join(f"fn{i} func(T{i}) T{i+1}" for i in range(n-1, 0, -1))

    cal = "".join(f"fn{i}(" for i in range(n, 0, -1)) + "t1" + ")" * n

    return f"""// Performs right-to-left function composition of {n} functions
func Compose{n}{gen}({fns}) func(T1) R  {{
    return func(t1 T1) R {{
        return {cal}
    }}
}}"""
