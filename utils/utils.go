package utils

func PointerAny[T comparable](v T) *T {
	if v == nil {
		return nil
	}
	return &v
}
