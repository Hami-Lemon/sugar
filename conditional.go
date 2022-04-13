package sugar

// Condition 三目运算符,条件con为true时，返回v1，否则返回v2
func Condition[T any](con bool, v1, v2 T) T {
	if con {
		return v1
	} else {
		return v2
	}
}
