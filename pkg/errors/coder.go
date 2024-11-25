package errors

import (
	"fmt"
	"sync"
)

type Coder interface {
	HTTPStatus() int
	Code() int
	String() string
	Reference() string
}

var codes = map[int]Coder{}
var codeMux = &sync.Mutex{}

func register(coder Coder) {
	if coder.Code() == 0 || coder.Code() == 999999 {
		fmt.Sprintf("code: %d reserved", coder.Code())
		return
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	codes[coder.Code()] = coder
}

func mustRegist(coder Coder) {
	if coder.Code() == 0 || coder.Code() == 999999 {
		fmt.Sprintf("code: %d reserved", coder.Code())
		return
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}

	codes[coder.Code()] = coder
}

func parser(code int) Coder {
	return codes[code]
}
