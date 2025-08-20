package utils

func PointerAny[T any](v T) *T {
	r := new(T)
	*r = v
	return r
}
