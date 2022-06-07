type MethodError struct {
	ErrorString string
}

func (e *MethodError) Error() string {
	return e.ErrorString
}

var (
	ErrInvalidQueryParam = &MethodError{"invalid query param"}
)
