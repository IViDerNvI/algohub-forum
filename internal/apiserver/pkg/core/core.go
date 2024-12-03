package core

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/pkg/errors"
	"github.com/sirupsen/logrus"
)

func GetAuthenticationHeader(c *gin.Context) (username, password string) {
	authHeader := c.GetHeader("Authorization")

	if !strings.HasPrefix(authHeader, "Basic ") {
		return
	}

	credentials := authHeader[len("Basic "):]
	decoded, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		return
	}

	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return
	}

	username, password = parts[0], parts[1]
	return
}

func WriteResponse(c *gin.Context, errcode int, data interface{}) {
	if errcode != 0 || data == nil {
		err := errors.Parse(errcode)
		logrus.Errorf(err.String())
		c.JSON(err.HTTPStatus(), err)
		return
	}

	c.JSON(http.StatusOK, data)
}
