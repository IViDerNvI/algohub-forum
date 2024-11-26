package code

// Basic errors code
const (
	ERROR_OK                = 000000
	ERROR_UNKNOWN           = 999999
	ERROR_PARAMETER_MISSING = 000001
)

// Following errors divided by moduel

// User model error code
const (
	ERROR_USER_USERNAME_INVALID = 010101 + iota
	ERROR_USER_USERNAME_DUPLICATE
	ERROR_USER_BIND_FAILED
	ERROR_USER_NOT_EXISTS
)

// Database error code
const (
	ERROR_DATABASE_INSERT_FAILED = 040101 + iota
	ERROR_DATABASE_SAVE_FAILED
	ERROR_DATABASE_DELETE_FAILED
)

// Post error code
const (
	ERROR_POST_BIND_FAILED = 010201 + iota
	ERROR_POST_NOT_EXISTS
)
