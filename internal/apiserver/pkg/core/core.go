package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/pkg/errors"
	"github.com/sirupsen/logrus"
)

func WriteResponse(c *gin.Context, errcode int, data interface{}) {
	if errcode != 0 {
		err := errors.Parse(errcode)
		logrus.Errorf(err.String())
		c.JSON(err.HTTPStatus(), err)
		return
	}

	c.JSON(http.StatusOK, data)
}
