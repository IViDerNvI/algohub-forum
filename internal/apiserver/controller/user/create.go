package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/auth"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/ividernvi/algohub-forum/pkg/util/idutil"
	"github.com/sirupsen/logrus"
)

func (uc UserController) Create(ctx *gin.Context) {
	logrus.WithContext(ctx)

	var u model.User

	if err := ctx.ShouldBind(&u); err != nil {
		logrus.Errorf("%+v", err)
		core.WriteResponse(ctx, code.ERROR_USER_BIND_FAILED, nil)
		return
	}

	u.ObjectMeta.InstanceID = idutil.GenerateUUID(u.GetDatabaseName())
	u.ObjectMeta.Name = u.Username
	u.ObjectMeta.CreatedAt = time.Now()
	u.ObjectMeta.UpdatedAt = time.Now()

	u.Password, _ = auth.Encrypt(u.Password)
	u.TotalScore = 0
	u.Status = 1

	logrus.Printf("user %+v will be insert", u)

	if err := uc.Usrv.Users().Create(ctx, u, model.CreateOptions{}); err != nil {
		logrus.Errorf("create user failed %+v", err)
		core.WriteResponse(ctx, code.ERROR_DATABASE_INSERT_FAILED, nil)
		return
	}

	core.WriteResponse(ctx, code.ERROR_OK, u)
}
