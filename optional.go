package optional

// ---

// New constructs a new optional Value.
func New[T any](value T, valid bool) Value[T] {
	if valid {
		return Some(value)
	}

	return None[T]()
}

// Some returns an optional Value that has provided inner value.
func Some[T any](value T) Value[T] {
	return Value[T]{
		value,
		true,
	}
}

// None returns an optional Value that has no inner value.
func None[T any]() Value[T] {
	return Value[T]{}
}

// ---

// Map transforms optional Value[T] to optional Value[U] using the given function.
func Map[T, U any, F ~func(T) U](v Value[T], f F) Value[U] {
	if v.IsSome() {
		return Some(f(v.value))
	}

	return None[U]()
}

// ---

// Value is an optional Value containing a value of type T inside.
type Value[T any] struct {
	value T
	valid bool
}

// Unwrap returns the inner value of type T and a true if present.
func (o Value[T]) Unwrap() (T, bool) {
	return o.value, o.valid
}

// IsSome returns true if the optional Value has inner value.
func (o Value[T]) IsSome() bool {
	return o.valid
}

// IsNone returns true if the optional Value has no inner value.
func (o Value[T]) IsNone() bool {
	return !o.IsSome()
}

// Or returns the Value if it is Some or other Value.
func (o Value[T]) Or(other Value[T]) Value[T] {
	if o.IsSome() {
		return o
	}

	return other
}

// OrSome returns the inner value T if present, otherwise it returns provided value.
func (o Value[T]) OrSome(value T) T {
	if o.IsSome() {
		return o.value
	}

	return value
}

// OrZero returns the inner value T if present, otherwise it returns zero initialized value.
func (o Value[T]) OrZero() T {
	return o.value
}

// OrElse returns Some(value) if the inner value is present, otherwise it calls provided function and returns its result.
func (o Value[T]) OrElse(value func() Value[T]) Value[T] {
	if o.IsSome() {
		return Some(o.value)
	}

	return value()
}

// Reset resets the optional Value to None.
func (o *Value[T]) Reset() {
	*o = None[T]()
}

// Take returns a copy of optional Value and resets it to None.
func (o *Value[T]) Take() Value[T] {
	result := *o
	*o = None[T]()

	return result
}

// Replace returns a copy of optional Value and resets it to Some(value).
func (o *Value[T]) Replace(value T) Value[T] {
	result := *o
	*o = Some(value)

	return result
}
