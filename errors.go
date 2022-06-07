package meclient

type ClientError struct {
	ErrorString string
}

func (e *ClientError) Error() string {
	return e.ErrorString
}

var (
	ErrInvalidQueryParam = &ClientError{"invalid query param"}
)
