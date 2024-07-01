package routes

import (
	"fmt"
	"strconv"

	"gos/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Moder struct {
	Name  string
	Code  int
	Child Childer
}

type Childer struct {
	Name string
	Code int
}

func pakerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data []Moder
		for i := 0; i < 10; i++ {
			child := Childer{
				Name: "child" + strconv.Itoa(i),
				Code: i,
			}
			moder := Moder{
				Name:  "moder" + strconv.Itoa(i),
				Code:  i,
				Child: child,
			}
			data = append(data, moder)
		}
		utils.Petition(c, 200, "操作成功", data)
	}
}

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
