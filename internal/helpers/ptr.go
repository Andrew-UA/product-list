package helpers

// Ptr Get pointer from value
func Ptr[T any](v T) *T {
	return &v
}
