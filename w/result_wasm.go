package w

type Result[T any, U comparable] struct {
	isErr bool
	value T
	err   U
}

func (r *Result[T, U]) IsOk() bool {
	return !r.isErr
}

func (r *Result[T, U]) IsErr() bool {
	return r.isErr
}

func (r *Result[T, U]) Get() (T, U) {
	return r.value, r.err
}

func (r *Result[T, U]) Set(value T, err U) {
	var errZero U
	if err == errZero {
		r.isErr = false
		r.value = value
		r.err = errZero
	} else {
		r.isErr = true
		var valueZero T
		r.value = valueZero
		r.err = err
	}
}
