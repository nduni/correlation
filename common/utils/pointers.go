package utils

func FromSliceToPointer[T any](s []T, i int) *T {
	if len(s)-1 < i {
		return nil
	}
	return &s[i]
}
