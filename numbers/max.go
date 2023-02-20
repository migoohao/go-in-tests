package numbers

type number interface {
	int64 | float64 | uint64 | int32 | float32 | uint32 | int16 | uint16 | int8 | uint8
}

func Max[T number](a, b T) T {
	if b > a {
		return b
	}
	return a
}
