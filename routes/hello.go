package routes

import (
	"gos/surface/dbo"
	"gos/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRow interface{}

func helloHandler(db *gorm.DB) gin.HandlerFunc {
	data, err := dbo.GetAllUserInfo(db)
	if err != nil {
		return func(c *gin.Context) {
			utils.Petition(c, 500, "服务器异常", nil)
		}
	}
	return func(c *gin.Context) {
		utils.Petition(c, 200, "请求成功", data)
	}
}
