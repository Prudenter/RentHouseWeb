package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"RentHouseWeb/rentHouseWeb/handler"
	"github.com/julienschmidt/httprouter"
)

/*
	web服务主函数
*/
func main() {
	// 创建web服务
	service := web.NewService(
		web.Name("go.micro.web.rentHouseWeb"),
		web.Version("latest"),
		web.Address("127.0.0.1:8888"),
	)

	// 服务初始化
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	//创建路由
	router := httprouter.New()

	//注册路由解析函数
	router.GET("/example/call", handler.ExampleCall)

	//将router注册到服务
	service.Handle("/", router)

	// 服务运行
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
