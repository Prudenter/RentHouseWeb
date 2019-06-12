package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

/*
	客户端业务处理函数
*/

// 路由解析函数(返回到前端的响应,前端传入的请求,路由解析请求后的参数)
func ExampleCall(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/*// 定义接受前端发送过来的json数据的容器
	var request map[string]interface{}
	// 将前端json数据解析到map当中
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//创建grpc客户端
	client := grpc.NewService()
	//初始化客户端
	client.Init()
	//通过protobuf生成文件调用创建客户端句柄的函数(服务名,客户端默认参数)
	exampleClient := example.NewExampleService("go.micro.srv.template", client.Client())
	//通过句柄调用服务端的业务处理函数
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 定义接受返回响应的容器
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// 将map集合转换为json数据,发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}*/
}
