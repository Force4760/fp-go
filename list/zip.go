package lst

import "github.com/repeale/fp-go/pair"

// Take 2 lists of A and B and merge them into a single list of Pair[A, B]
// If the lists don't have the same size, the final list will have the same size as the smaller one
func Zip[B, A any](as []A) func([]B) []pair.Pair[A, B] {
	return func(bs []B) (res []pair.Pair[A, B]) {
		i := 0

		for i < len(as) && i < len(bs) {
			res = append(res, pair.New(as[i], bs[i]))
			i++
		}

		return
	}
}

// Take an array of Pair[A, B] and return a Pair of arrays, the first of type []A and the second of []B
func Unzip[A, B any](xs []pair.Pair[A, B]) pair.Pair[[]A, []B] {
	as := make([]A, len(xs))
	bs := make([]B, len(xs))

	for i, x := range xs {
		as[i] = pair.Fst(x)
		bs[i] = pair.Snd(x)
	}

	return pair.New(as, bs)
}

// Take 2 lists of A and B and merge them into a single list of C with a function f :: (A, B) -> C
// If the lists don't have the same size, the final list will have the same size as the smaller one
func ZipWith[A, B, C any](fn func(A, B) C) func([]A) func([]B) []C {
	return func(as []A) func([]B) []C {
		return func(bs []B) (res []C) {
			i := 0

			for i < len(as) && i < len(bs) {
				res = append(res, fn(as[i], bs[i]))
				i++
			}

			return
		}
	}
}
