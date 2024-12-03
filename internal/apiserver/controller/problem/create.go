package problem

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/ividernvi/algohub-forum/pkg/util/idutil"
	"github.com/sirupsen/logrus"
)

func (pc ProbController) Create(ctx *gin.Context) {
	logrus.WithContext(ctx)

	var p model.Problem

	if err := ctx.ShouldBind(&p); err != nil {
		logrus.Errorf("%+v", err)
		core.WriteResponse(ctx, code.ERROR_USER_BIND_FAILED, nil)
		return
	}

	p.InstanceID = idutil.GenerateUUID("prob")
	p.ObjectMeta.Name = p.Title
	p.ObjectMeta.CreatedAt = time.Now()
	p.ObjectMeta.UpdatedAt = time.Now()

	logrus.Printf("user %+v will be insert", p)

	if err := pc.ProvSrv.Problems().Create(ctx, &p, model.CreateOptions{}); err != nil {
		logrus.Errorf("create user failed %+v", err)
		core.WriteResponse(ctx, code.ERROR_DATABASE_INSERT_FAILED, nil)
		return
	}

	core.WriteResponse(ctx, code.ERROR_OK, p)
}
