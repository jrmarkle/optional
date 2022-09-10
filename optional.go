package optional

type Optional[T any] struct {
	value T
	valid bool
}

func None[T any]() Optional[T] {
	return Optional[T]{valid: false}
}

func Some[T any](v T) Optional[T] {
	return Optional[T]{valid: true, value: v}
}

func (o *Optional[T]) Is() bool {
	return o.valid
}

func (o *Optional[T]) Get() T {
	if o.valid {
		return o.value
	}
	var zeroValue T
	return zeroValue
}

func (o *Optional[T]) GetOr(defaultValue T) T {
	if o.valid {
		return o.value
	}
	return defaultValue
}
