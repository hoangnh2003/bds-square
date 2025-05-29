package response

const (
	ErrCodeSuccess       = 20001
	ErrCodeParamInvalid  = 20003
	ErrInvalidToken      = 30001
	ErrCodeUserHasExists = 50001
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHasExists: "User has already registered",
}
