package code

import "github.com/ividernvi/algohub-forum/pkg/errors"

func init() {
	normalCodeRegiste()
	userCodeRegiste()
	dbCodeRegiste()
	postCodeRegiste()
}

func normalCodeRegiste() {
	errors.Registe(ERROR_OK, 200, "OK", "ok")
	errors.Registe(ERROR_UNKNOWN, 500, "unknown error occured", "unkown")
	errors.Registe(ERROR_PARAMETER_MISSING, 400, "parameter missing", "missparam")
}

func userCodeRegiste() {
	errors.Registe(ERROR_USER_USERNAME_INVALID, 400, "username invalid", "usernameinvalid")
	errors.Registe(ERROR_USER_USERNAME_DUPLICATE, 400, "username already exists", "usernameduplicate")
	errors.Registe(ERROR_USER_BIND_FAILED, 400, "json format error", "bindfailed")
	errors.Registe(ERROR_USER_NOT_EXISTS, 404, "user not exists", "usernotexists")
	errors.Registe(ERROR_USER_AUTHENTICATION_FAILED, 400, "wrong password or account", "authentication")
}

func dbCodeRegiste() {
	errors.Registe(ERROR_DATABASE_INSERT_FAILED, 500, "database error", "dbfailed")
	errors.Registe(ERROR_DATABASE_SAVE_FAILED, 500, "database save failed", "dbsavefail")
	errors.Registe(ERROR_DATABASE_DELETE_FAILED, 500, "database delete failed", "dbdeletefail")
	errors.Registe(ERROR_DATABASE_GET_FAILED, 500, "database get failed", "ERROR_DATABASE_GET_FAILED")
}

func postCodeRegiste() {
	errors.Registe(ERROR_POST_BIND_FAILED, 500, "json format error", "bindfailed")
	errors.Registe(ERROR_POST_NOT_EXISTS, 404, "post not exists", "postnotexists")
}
