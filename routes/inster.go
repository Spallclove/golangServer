package routes

import (
	"fmt"
	"gos/surface/dbo"
	"gos/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FormData struct {
	Name     string `form:"name"`
	ImageUrl string `form:"image_url"`
	VideoUrl string `form:"video_url"`
	ImageId  string `form:"image_id"`
	VideoId  string `form:"video_id"`
}

// 新增
func insterAddHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := map[string]string{
			"name":      string(c.PostForm("name")),
			"image_url": string(c.PostForm("image_url")),
			"video_url": string(c.PostForm("video_url")),
			"image_id":  string(c.PostForm("image_id")),
			"video_id":  string(c.PostForm("video_id")),
		}
		err := dbo.AddOneDataInfo(db, data)
		if err != nil {
			utils.Petition(c, 500, "操作失败", nil)
			// log.Fatal("失败了", err)
			return
		}
		utils.Petition(c, 200, "操作成功", nil)
	}
}

func insterUpdateHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("id")
		fmt.Println(c.PostForm("id"), "idididid")
		trimmedId := strings.TrimSpace(id)

		if trimmedId == "" {
			utils.Petition(c, 500, "id 不能为空", nil)
			return
		}
		data := map[string]string{
			"id":        c.PostForm("id"),
			"name":      c.PostForm("name"),
			"image_url": c.PostForm("image_url"),
			"video_url": c.PostForm("video_url"),
			"image_id":  c.PostForm("image_id"),
			"video_id":  c.PostForm("video_id"),
		}
		fmt.Print("6666:", data)
		err := dbo.UpdateOneDataInfo(db, data)
		if err != nil {
			utils.Petition(c, 500, "操作失败", err)
			return
		}
		utils.Petition(c, 200, "操作成功", nil)
	}
}
