# Result

Result represents encapsulation of a possibly erroneous value, it might be used as the return type of functions which may or may not return an error when they are applied.
You could instanciate an `res.Result[T]` with a value with `res.Ok(val)`. If there is an error you can use `res.Err[T](err)`.

Error exports: `Ok`, `Err`, `IsOk`, `IsErr`, `IsOkAnd`, `Map`, `Flat`, `GetOrElse`, `Eq`, `Chain`, `Match`, `String`, `Error`

Error implements the `Stringer` and the `error` interfaces.