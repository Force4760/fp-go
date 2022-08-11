# Either

Either represents encapsulation of a value that can be of two types.

You could instanciate an `either.Either[T]` with `either.Left[R](val)` or with `either.Right[L](val)`.

Either exports: `Left`, `Right`, `IsLeft`, `IsRight`, `IsLeftAnd`,  `IsRightAnd`, `GetOrElseL`, `GetOrElseR`, `Eq`, `Flip`, `Match`, `MapL`, `MapR`, `ChainL`, `ChainR`, `FlatL`, `FlatR`, `String`.

Either implements the `Stringer` interface
