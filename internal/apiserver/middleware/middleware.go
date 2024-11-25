package middleware

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	middlewares     = map[string]gin.HandlerFunc{}
	middlewareMutex = &sync.Mutex{}
)

type Middleware gin.HandlerFunc

func registe(label string, middleware gin.HandlerFunc) {
	middlewareMutex.Lock()
	defer middlewareMutex.Unlock()

	middlewares[label] = middleware
}

func mustRegiste(label string, middleware gin.HandlerFunc) {
	middlewareMutex.Lock()
	defer middlewareMutex.Unlock()

	if _, ok := middlewares[label]; ok {
		panic(fmt.Sprintf("middleware: %s already exist", label))
	}

	middlewares[label] = middleware
}

func Install(label string, g *gin.Engine) {
	middlewareMutex.Lock()
	defer middlewareMutex.Unlock()

	fn := middlewares[label]

	g.Use(fn)
}

func MustInstall(lable string, g *gin.Engine) {
	middlewareMutex.Lock()
	defer middlewareMutex.Unlock()
	if _, ok := middlewares[lable]; !ok {
		panic(fmt.Sprintf("middleware: %s is not exists"))
	}

	g.Use(middlewares[lable])
}
