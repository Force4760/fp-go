# Either

Either represents encapsulation of a value that can be of two types.

You could instanciate an `either.Either[T]` with `either.Left[R](val)` or with `either.Right[L](val)`.

Either exports: 
* `Left`, `Right`
* `IsLeft`, `IsRight`
* `IsLeftAnd`,  `IsRightAnd`
* `GetOrElseL`, `GetOrElseR`
* `MapL`, `MapR`
* `ChainL`, `ChainR`
* `FlatL`, `FlatR`
* `Eq`
* `Flip`
* `Match`
* `String`

Either implements the `Stringer` interface
