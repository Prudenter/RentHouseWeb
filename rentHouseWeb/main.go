package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"RentHouseWeb/rentHouseWeb/handler"
	"github.com/julienschmidt/httprouter"
	_ "RentHouseWeb/rentHouseWeb/models"
	"RentHouseWeb/rentHouseWeb/utils"
	"net/http"
)

/*
	web服务主函数
*/
func main() {
	// 创建web服务
	service := web.NewService(
		web.Name("go.micro.web.rentHouseWeb"),
		web.Version("latest"),
		web.Address(utils.G_server_addr+":"+utils.G_server_port),
	)

	// 服务初始化
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	//创建路由
	router := httprouter.New()

	//文件服务器映射静态页面
	router.NotFound = http.FileServer(http.Dir("html"))

	//注册区域服务路由解析函数
	router.GET("/api/v1.0/areas", handler.GetArea)

	//解决session报错问题
	router.GET("/api/v1.0/session", handler.GetSession)
	//解决index报错问题
	router.GET("/api/v1.0/house/index", handler.GetIndex)

	//注册图片验证码服务路由解析函数
	router.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)

	//注册短信验证码服务路由解析函数
	router.GET("/api/v1.0/smscode/:mobile", handler.GetSmscd)

	//注册用户注册服务路由解析函数
	router.POST("/api/v1.0/users", handler.PostRet)


	//将router注册到服务
	service.Handle("/", router)

	// 服务运行
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
