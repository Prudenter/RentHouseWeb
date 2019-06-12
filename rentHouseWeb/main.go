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

	//注册路由解析函数--模板调用
	router.GET("/example/call", handler.ExampleCall)

	//将router注册到服务
	service.Handle("/", router)

	// 服务运行
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
