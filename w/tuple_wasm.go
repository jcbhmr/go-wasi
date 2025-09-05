package w

import "structs"

type Tuple[A, B any] struct {
	_ structs.HostLayout
	A A
	B B
}
