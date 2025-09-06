package w

import "structs"

type Tuple[A, B any] struct {
	_ structs.HostLayout
	A A
	B B
}

func (t *Tuple[A, B]) Get() (A, B) {
	return t.A, t.B
}