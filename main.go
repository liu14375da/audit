package main

import (
	_ "AuditInformation/docs"
	"AuditInformation/internal/audit"
	Cors "AuditInformation/internal/cors"
	"fmt"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title 获取审核信息
// @version 1.0
// @description 点击信息按钮，获取审核人，意见等信息
// @termsOfService http://swagger.io/terms/

// @contact.name 有问题请找前端刘娇娇,联系电话：123567890
// @contact.url http://www.swagger.io/support

func main() {
	fmt.Println("测试")
	router := gin.Default()
	router.Use(Cors.Cors()) // 允许使用跨域请求	 全局中间件
	url := gs.URL("http://10.28.83.123:8863/swagger/doc.json")
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler,url))
	router.GET("/Audit", audit.AuInfo)
	router.Run(":8863")
}
