package main

import (
	"github.com/gin-gonic/gin"
	"k8s-platform/config"
	"k8s-platform/controller"
	"k8s-platform/middle"
	"k8s-platform/service"
	"net/http"
)

func main() {
	//初始化gin对象
	r := gin.Default()
	//初始化k8s client
	service.K8s.Init()
	//数据库初始化
	//db.Init()
	//注册中间件
	r.Use(middle.Cors())
	r.Use(middle.JWTAuth())
	//初始化路由规则
	controller.Router.InitApiRouter(r)
	//终端websocket
	go func() {
		http.HandleFunc("/ws", service.Terminal.WsHandler)
		http.ListenAndServe(":8081", nil)
	}()
	//gin程序启动
	r.Run(config.ListenAddr)
}
