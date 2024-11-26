package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/controller/user"
	"github.com/ividernvi/algohub-forum/internal/apiserver/middleware"
	"github.com/ividernvi/algohub-forum/internal/apiserver/store/mysql"
	"github.com/ividernvi/algohub-forum/internal/model/options"
)

func InitRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) *gin.Engine {
	middleware.MustInstall("logger", g)
	return g
}

func installController(g *gin.Engine) *gin.Engine {
	storeIns, _ := mysql.GetMySQLFactoryOr(*options.NewDBOptions())
	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)

			userv1.POST("", userController.Create)
			userv1.PUT(":name", userController.Update)
			userv1.DELETE(":name", userController.Delete)
			userv1.GET(":name", userController.Get)
		}
	}
	return g
}
