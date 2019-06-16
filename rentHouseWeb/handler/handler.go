package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/micro/go-grpc"
	"context"
	GETAREA "RentHouseWeb/getArea/proto/example"
	GETIMAGECD "RentHouseWeb/getImageCd/proto/example"
	GETSMSCD "RentHouseWeb/getSmscd/proto/example"
	POSTRET "RentHouseWeb/postRet/proto/example"
	GETSESSION "RentHouseWeb/getSession/proto/example"
	"encoding/json"
	"RentHouseWeb/rentHouseWeb/models"
	"image"
	"github.com/afocus/captcha"
	"image/png"
	"regexp"
	"RentHouseWeb/rentHouseWeb/utils"
)

/*
	区域服务web端业务处理函数
	路由解析函数(返回到前端的响应,前端传入的请求,路由解析请求后的参数)
*/
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
	w.Header().Set("Content-Type", "application/json")

	// 将map集合转换为json数据,发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

/*
	登录检查服务业务处理函数
*/
func GetSession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("登录检查服务,GetSession...")

	//获取sessionId
	cookie, err := r.Cookie("userLogin")
	if err != nil {
		// 准备返回给前端的map
		response := map[string]interface{}{
			"errno":  "4101",
			"errmsg": "用户未登录",
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")

		// 将map集合转换为json数据,发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	//创建 grpc 客户端
	cli := grpc.NewService()
	//客户端初始化
	cli.Init()

	// call the backend service
	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := GETSESSION.NewExampleService("go.micro.srv.getSession", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.GetSession(context.TODO(), &GETSESSION.Request{
		SessionId: cookie.Value,
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//将名字接收到
	data := make(map[string]string)
	data["name"] = rsp.Name

	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}

	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// 准备返回给前端的map
	response := map[string]interface{}{
		"errno":  "0",
		"errmsg": "ok",
	}

	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")

	// 将map集合转换为json数据,发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

/*
	图片验证码服务web端业务处理函数
*/
func GetImageCd(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("获取图片验证码服务,GetImageCd..")

	//获取url中所带的参数,uuid
	uuid := params.ByName("uuid")
	// 创建grpc客户端
	client := grpc.NewService()
	// 初始化客户端
	client.Init()
	// 通过protobuf生成文件调用创建客户端句柄的函数(服务名,客户端默认参数)
	exampleClient := GETIMAGECD.NewExampleService("go.micro.srv.getImageCd", client.Client())
	// 通过句柄调用服务端的业务处理函数,获取响应数据
	rsp, err := exampleClient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid: uuid,
	})
	// 判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//判断返回值,如果不等于0,直接返回错误
	if rsp.Errno != "0" {
		// 准备返回给前端的map
		response := map[string]interface{}{
			"errno":  rsp.Errno,
			"errmsg": rsp.Errmsg,
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")

		// 将map集合转换为json数据,发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//接受图片数据,拼接图片结构体,并将图片发送给前端
	//定义接受图片数据的变量
	var rgba image.RGBA
	for _, value := range rsp.Pix {
		//循环接收Pix
		rgba.Pix = append(rgba.Pix, uint8(value))
	}
	//接收Stride
	rgba.Stride = int(rsp.Stride)
	//接收point
	rgba.Rect.Min.X = int(rsp.Min.X)
	rgba.Rect.Min.Y = int(rsp.Min.Y)
	rgba.Rect.Max.X = int(rsp.Max.X)
	rgba.Rect.Max.Y = int(rsp.Max.Y)
	var image captcha.Image
	image.RGBA = &rgba
	//将图片发送给前端
	png.Encode(w, image)

}

/*
	短信验证码服务web端业务处理函数
*/
func GetSmscd(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("获取短信验证码服务,GetSmscd..")

	//获取手机号
	mobile := params.ByName("mobile")
	//验证手机号
	reg := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
	bool := reg.MatchString(mobile)
	if bool == false {
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//获取图片验证码和uuid
	text := r.URL.Query()["text"][0]
	uuid := r.URL.Query()["id"][0]

	//创建grpc客户端
	client := grpc.NewService()
	//客户端初始化
	client.Init()
	// 通过protobuf生成文件调用创建客户端句柄的函数(服务名,客户端默认参数)
	exampleClient := GETSMSCD.NewExampleService("go.micro.srv.getSmscd", client.Client())
	// 通过句柄调用服务端的业务处理函数,获取响应数据
	rsp, err := exampleClient.GetSmscd(context.TODO(), &GETSMSCD.Request{
		Mobile: mobile,
		Uuid:   uuid,
		Text:   text,
	})
	// 判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// 准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")

	// 将map集合转换为json数据,发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

/*
	用户注册服务web端业务处理函数
*/
func PostRet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("用户注册服务,PostRet..")

	// 接收前端发送过来的数据
	var request map[string]interface{}
	// 将前端json数据解析到map中
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//校验数据
	if request["mobile"].(string) == "" || request["password"].(string) == "" || request["sms_code"].(string) == "" {
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建grpc客户端
	client := grpc.NewService()
	//客户端初始化
	client.Init()
	// 通过protobuf生成文件调用创建客户端句柄的函数(服务名,客户端默认参数)
	exampleClient := POSTRET.NewExampleService("go.micro.srv.postRet", client.Client())
	// 通过句柄调用服务端的业务处理函数,获取响应数据
	rsp, err := exampleClient.PostRet(context.TODO(), &POSTRET.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
		SmsCode:  request["sms_code"].(string),
	})
	// 判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// 准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	//读取存储sessionId的cookie
	cookie, err := r.Cookie("userLogin")
	// 如果读取失败或者cookie的value为空则创建cookie
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "userLogin", Value: rsp.SessionId, MaxAge: 600, Path: "/"}
		http.SetCookie(w, &cookie)
	}

	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")

	// 将map集合转换为json数据,发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
