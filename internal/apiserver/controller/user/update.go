package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (uc UserController) Update(ctx *gin.Context) {
	logrus.WithContext(ctx)

	var u model.User

	if err := ctx.ShouldBind(&u); err != nil {
		logrus.Errorf("%+v", err)
		core.WriteResponse(ctx, code.ERROR_USER_BIND_FAILED, nil)
		return
	}

	old_user, err := uc.Usrv.Users().GetByAccount(ctx, ctx.Param("name"), model.GetOptions{})

	if err != nil {
		logrus.Errorf("%+v", err)
		core.WriteResponse(ctx, code.ERROR_USER_NOT_EXISTS, nil)
		return
	}

	old_user.Username = u.Username
	old_user.Email = u.Email
	old_user.UpdatedAt = time.Now()

	if err := uc.Usrv.Users().Update(ctx, *old_user, model.UpdateOptions{}); err != nil {
		logrus.Errorf("save user failed %+v", err)
		core.WriteResponse(ctx, code.ERROR_DATABASE_SAVE_FAILED, nil)
		return
	}

	core.WriteResponse(ctx, 0, old_user)
}
