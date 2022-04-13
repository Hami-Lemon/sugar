package sugar

type float interface {
	float32 | float64
}

type integer interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

type number interface {
	integer | float
}

func Max[T number](v1, v2 T) T {
	return Condition(v1 > v2, v1, v2)
}

func Min[T number](v1, v2 T) T {
	return Condition(v1 < v2, v1, v2)
}

func Abs[T number](num T) T {
	return Condition(num < 0, -num, num)
}
