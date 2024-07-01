package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Petition(c *gin.Context, state int, message, data any) {
	flag := http.StatusOK
	switch state {
	case 100:
		flag = http.StatusContinue
	case 200:
		flag = http.StatusOK
	case 201:
		flag = http.StatusCreated
	case 202:
		flag = http.StatusAccepted
	case 300:
		flag = http.StatusMovedPermanently
	case 302:
		flag = http.StatusFound
	case 400:
		flag = http.StatusBadRequest
	case 401:
		flag = http.StatusUnauthorized
	case 403:
		flag = http.StatusForbidden
	case 503:
		flag = http.StatusServiceUnavailable
	default:
		flag = http.StatusInternalServerError
	}

	if data == nil {
		c.JSON(flag, gin.H{
			"code":    flag,
			"message": message,
		})
		return
	}

	c.JSON(flag, gin.H{
		"data":    data,
		"code":    flag,
		"message": message,
	})

}
