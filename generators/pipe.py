def pipe_func(n: int) -> str:
    gen = "[" + ", ".join(f"T{i+1}" for i in range(n)) + ", R any]"

    fns = ", ".join(f"fn{i} func(T{i}) T{i+1}" for i in range(1, n)) + f", fn{n} func(T{n}) R"

    cal = "".join(f"fn{i}(" for i in range(n, 0, -1)) + "t1" + ")" * n

    return f"""// Performs left-to-right function composition of {n} functions
func Pipe{n}{gen}({fns}) func(T1) R  {{
    return func(t1 T1) R {{
        return {cal}
    }}
}}"""