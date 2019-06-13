package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/micro/go-grpc"
	"context"
	GETAREA "RentHouseWeb/getArea/proto/example"
	"encoding/json"
	"RentHouseWeb/rentHouseWeb/models"
)

/*
	客户端业务处理函数
*/

// 路由解析函数(返回到前端的响应,前端传入的请求,路由解析请求后的参数)
func GetArea(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("获取区域信息服务,GetArea...")
	// 创建grpc客户端
	client := grpc.NewService()
	// 初始化客户端
	client.Init()
	// 通过protobuf生成文件调用创建客户端句柄的函数(服务名,客户端默认参数)
	exampleClient := GETAREA.NewExampleService("go.micro.srv.getArea", client.Client())
	// 通过句柄调用服务端的业务处理函数
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	// 判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 定义接收数据的容器
	area_list := []models.Area{}
	// 接收srv端返回数据
	for _, value := range rsp.Data {
		temp := models.Area{Id: int(value.AId), Name: value.AName}
		area_list = append(area_list, temp)
	}

	// 准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   area_list,
	}

	//设置返回数据的格式
	w.Header().Set("Content-Type","application/json")

	// 将map集合转换为json数据,发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
