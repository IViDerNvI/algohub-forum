package problem

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (pc ProbController) List(ctx *gin.Context) {
	var opts model.ListOptions

	opts.Kind = ctx.Query("kind")
	opts.APIVersion = ctx.Query("apiVersion")
	opts.QueryKey = ctx.Query("queryKey")
	opts.Pattern = ctx.Query("pattern")

	timeoutSecondsStr := ctx.Query("timeoutSeconds")
	if timeoutSecondsStr != "" {
		timeout, _ := strconv.ParseInt(timeoutSecondsStr, 10, 64)
		opts.TimeoutSeconds = int64(timeout)
	}

	offsetStr := ctx.Query("offset")
	if offsetStr != "" {
		offset, _ := strconv.ParseInt(offsetStr, 10, 64)
		opts.Offset = int64(offset)
	}

	limitStr := ctx.Query("limit")
	if limitStr != "" {
		limit, _ := strconv.ParseInt(limitStr, 10, 64)
		opts.Limit = int64(limit)
	}

	logrus.Printf("search with opt: %#+v", opts)

	var probs *model.ProblemList
	var err error
	if probs, err = pc.ProvSrv.Problems().List(ctx, opts); err != nil {
		logrus.Printf("list failed: $s", err)
		core.WriteResponse(ctx, code.ERROR_DATABASE_GET_FAILED, nil)
		return
	}
	core.WriteResponse(ctx, code.ERROR_OK, probs)
}
