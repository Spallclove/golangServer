package routes

import (
	"fmt"

	"gos/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func testHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		fmt.Println(name)
		if name == "" {
			utils.Petition(c, 400, "服务不支持", nil)
			return
		}
		data := []any{1, "12", "5", 150}
		utils.Petition(c, 200, "请求成功", data)
	}
}
