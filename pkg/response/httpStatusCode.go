package response

const (
	ErrCodeSuccess       = 20001
	ErrCodeParamInvalid  = 20003
	ErrInvalidToken      = 30001
	ErrCodeUserHasExists = 50001
	ErrCodeServerError   = 50000
	ErrCodeNotFound      = 404
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHasExists: "User has already registered",
	ErrCodeServerError:   "Server error",
	ErrCodeNotFound:      "Not found",
}
