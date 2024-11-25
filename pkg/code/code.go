package code

// Basic errors
const (
	ERROR_OK                = 000000
	ERROR_UNKNOWN           = 999999
	ERROR_PARAMETER_MISSING = 000001
)

// Following errors divided by moduel

// User model error
const (
	ERROR_USER_USERNAME_INVALID = 010101 + iota
	ERROR_USER_USERNAME_DUPLICATE
	ERROR_USER_BIND_FAILED
	ERROR_USER_NOT_EXISTS
)

// Database error
const (
	ERROR_DATABASE_INSERT_FAILED = 040101 + iota
	ERROR_DATABASE_SAVE_FAILED
)
