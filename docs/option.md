# Options

Option represents encapsulation of an optional value, it might be used as the return type of functions which may or may not return a meaningful value when they are applied.
You could instanciate an `opt.Option[T]` with a value with `opt.Some(val)`. If the value is missing you can use `opt.None[T]()`.

Option exports:`Some`, `None`, `IsSome`, `IsSomeAnd`, `IsNone`,  `FromPtr`, `ToPtr`, `GetOrElse`, `Match`, `Map`, `Chain`, `Filter`, `Flat`, `Eq`, `String`.

Option implements the `Stringer` interface
