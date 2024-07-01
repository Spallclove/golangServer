package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/api/upload", uploadHandler(db))
	r.GET("/api/hello", helloHandler(db))
	r.GET("/api/test", testHandler(db))
	r.POST("/api/inster/add", insterAddHandler(db))
	r.POST("/api/inster/update", insterUpdateHandler(db))
}
