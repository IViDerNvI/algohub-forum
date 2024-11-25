package errors

import (
	"errors"
	"fmt"
)

const (
	BASE_ERROR_REFERENCE_URL = "https://github.com/ividernvi/errors"
)

type ErrCode struct {
	ErrC     int
	HttpCode int
	Message  string
	Refer    string
}

var _ Coder = &ErrCode{}

func (e *ErrCode) String() string {
	return e.Message
}

func (e *ErrCode) HTTPStatus() int {
	return e.HttpCode
}

func (e *ErrCode) Code() int {
	return e.ErrC
}

func (e *ErrCode) Reference() string {
	return fmt.Sprintf("%s%s", BASE_ERROR_REFERENCE_URL, e.Refer)
}

func (e *ErrCode) Error() string {
	return e.Message
}

func Registe(errcode int, httpcode int, message string, reference string) {
	code := &ErrCode{
		ErrC:     errcode,
		HttpCode: httpcode,
		Message:  message,
		Refer:    reference,
	}
	mustRegist(code)
}

func New(code int) error {
	if codes[code].String() == "" {
		panic(fmt.Sprintf("code %s is not defined", code))
	}
	return errors.New(codes[code].String())
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Parse(code int) Coder {
	return parser(code)
}
