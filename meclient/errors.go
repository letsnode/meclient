package meclient

type MeClientError struct {
	ErrorString string
}

func (e *MeClientError) Error() string {
	return e.ErrorString
}

var (
	ErrInvalidQueryParam = &MeClientError{"invalid query param"}
)
