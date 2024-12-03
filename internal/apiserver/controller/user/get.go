package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (uc UserController) Get(ctx *gin.Context) {
	logrus.WithContext(ctx)

	username := ctx.Param("name")

	if user, err := uc.Usrv.Users().GetByAccount(ctx, username, model.GetOptions{}); err != nil {
		core.WriteResponse(ctx, code.ERROR_USER_NOT_EXISTS, nil)
		return
	} else {
		core.WriteResponse(ctx, code.ERROR_OK, user)
		return
	}
}
