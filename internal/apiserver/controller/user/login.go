package user

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/auth"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/ividernvi/algohub-forum/pkg/util/jwtutil"
	"github.com/sirupsen/logrus"
)

func (uc *UserController) Login(ctx *gin.Context) {
	logrus.WithContext(ctx)

	username, passwd := core.GetAuthenticationHeader(ctx)

	logrus.Printf("%s in Authentication", ctx.GetHeader("Authorization"))
	logrus.Printf("%s:%s try to login", username, passwd)

	user, err := uc.Usrv.Users().GetByAccount(ctx, username, model.GetOptions{})
	if err != nil {
		logrus.Errorf("find user failed %+v", err)
		core.WriteResponse(ctx, code.ERROR_DATABASE_INSERT_FAILED, nil)
		return
	}

	if auth.Compare(user.Password, passwd) != nil {
		logrus.Errorf("user %#+v try to login but failed", username)
		core.WriteResponse(ctx, code.ERROR_USER_AUTHENTICATION_FAILED, nil)
		return
	}

	tokenString, _ := jwtutil.CreateToken(user.InstanceID)
	ctx.Header("Authorization", fmt.Sprintf("Bearer %s", tokenString))

	var t = &model.Token{
		UserID:    user.InstanceID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(2 * time.Hour),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = uc.Usrv.Tokens().Create(ctx, *t, model.CreateOptions{}); err != nil {
		logrus.Errorf("redis token insert failed", username)
		core.WriteResponse(ctx, code.ERROR_DATABASE_DELETE_FAILED, nil)
		return
	}

	core.WriteResponse(ctx, code.ERROR_OK, "OK")
}
