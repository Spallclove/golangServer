package main

import (
	"fmt"
	"gos/database"
	"gos/middleware"
	"gos/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(middleware.Cors())
	db, err := database.ConnectSql()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	// 设置路由组
	routes.SetupRoutes(r, db)
	// 静态文件服务
	r.Static("/api/uploads", "./api/uploads")
	// 运行服务器
	r.Run(":8080")
}
