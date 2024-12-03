package code

import "github.com/ividernvi/algohub-forum/pkg/errors"

func init() {
	normalCodeRegiste()
	userCodeRegiste()
	dbCodeRegiste()
	postCodeRegiste()
}

func normalCodeRegiste() {
	errors.Registe(ERROR_OK, 200, "OK", "OK")
	errors.Registe(ERROR_UNKNOWN, 500, "unknown error occured", "ERROR_UNKNOWN")
	errors.Registe(ERROR_PARAMETER_MISSING, 400, "parameter missing", "ERROR_PARAMETER_MISSING")
}

func userCodeRegiste() {
	errors.Registe(ERROR_USER_USERNAME_INVALID, 400, "username invalid", "ERROR_USER_USERNAME_INVALID")
	errors.Registe(ERROR_USER_USERNAME_DUPLICATE, 400, "username already exists", "ERROR_USER_USERNAME_DUPLICATE")
	errors.Registe(ERROR_USER_BIND_FAILED, 400, "json format error", "ERROR_USER_BIND_FAILED")
	errors.Registe(ERROR_USER_NOT_EXISTS, 404, "user not exists", "ERROR_USER_NOT_EXISTS")
	errors.Registe(ERROR_USER_AUTHENTICATION_FAILED, 400, "wrong password or account", "ERROR_USER_AUTHENTICATION_FAILED")
}

func dbCodeRegiste() {
	errors.Registe(ERROR_DATABASE_INSERT_FAILED, 500, "database error", "ERROR_DATABASE_INSERT_FAILED")
	errors.Registe(ERROR_DATABASE_SAVE_FAILED, 500, "database save failed", "ERROR_DATABASE_SAVE_FAILED")
	errors.Registe(ERROR_DATABASE_DELETE_FAILED, 500, "database delete failed", "ERROR_DATABASE_DELETE_FAILED")
	errors.Registe(ERROR_DATABASE_GET_FAILED, 500, "database get failed", "ERROR_DATABASE_GET_FAILED")
}

func postCodeRegiste() {
	errors.Registe(ERROR_POST_BIND_FAILED, 500, "json format error", "ERROR_POST_BIND_FAILED")
	errors.Registe(ERROR_POST_NOT_EXISTS, 404, "post not exists", "ERROR_POST_NOT_EXISTS")
}
