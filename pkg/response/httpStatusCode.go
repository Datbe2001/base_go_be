package response

const (
	ErrCodeSuccess       = 2001  //Success
	ErrCodeInvalidParams = 2002  //Email invalid
	ErrInvalidToken      = 3001  //Token invalid
	ErrCodeUserHasExists = 50001 // User already exist
	ErrCodeUserNotFound  = 4000  // User not found
)

var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrInvalidToken:      "Token invalid",
	ErrCodeInvalidParams: "Email invalid",
	ErrCodeUserHasExists: "User already exist",
	ErrCodeUserNotFound:  "User not found",
}
