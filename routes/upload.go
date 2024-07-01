package routes

import (
	"fmt"
	"gos/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 上传目录
const uploadDir = "./api/uploads"

// 创建上传目录
func init() {
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}
}

func uploadHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		fileType := file.Header.Get("Content-Type")

		switch fileType {
		case "image/jpeg", "image/png", "image/gif":
			fileType = "image"
		case "video/mp4", "video/quicktime":
			// 如果是视频类型，保存到视频目录
			fileType = "video"
		default:
			c.String(http.StatusBadRequest, "Unsupported file type")
			return
		}

		// 给文件命名以避免重复
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
		filePath := filepath.Join(uploadDir, filename)

		// 保存文件到指定路径
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		// 判断文件类型并处理
		if fileType == "video" {
			handleVideoUpload(c, filePath, filename)
		} else {
			fileURL := fmt.Sprintf("http://localhost:8080/api/uploads/%s", filename)
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"url":  fileURL,
				"type": fileType,
				"data": fileURL,
			})
		}
	}
}

func handleVideoUpload(c *gin.Context, filePath, filename string) {
	// 生成缩略图路径
	thumbnailFilename := fmt.Sprintf("%d_thumbnail.png", time.Now().UnixNano())
	thumbnailPath := filepath.Join(uploadDir, thumbnailFilename)

	// 生成缩略图
	err := utils.GenerateThumbnail(filePath, thumbnailPath)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("generate thumbnail err: %s", err.Error()))
		return
	}

	// 返回视频和缩略图的访问URL
	videoURL := fmt.Sprintf("http://localhost:8080/api/uploads/%s", filename)
	thumbnailURL := fmt.Sprintf("http://localhost:8080/api/uploads/%s", thumbnailFilename)
	c.JSON(http.StatusOK, gin.H{
		"code":         1,
		"url":          thumbnailURL,
		"type":         "video",
		"data":         thumbnailURL,
		"thumbnailUrl": videoURL,
	})
}
