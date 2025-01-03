package response

const (
	ErrCodeSuccess       = 2001
	ErrCodeInvalidParams = 2002
	ErrInvalidToken      = 3001
)

var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrInvalidToken:      "Token invalid",
	ErrCodeInvalidParams: "Email invalid",
}
