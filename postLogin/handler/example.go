package handler

import (
	"context"

	example "RentHouseWeb/postLogin/proto/example"
	"fmt"
	"RentHouseWeb/rentHouseWeb/utils"
	"github.com/astaxie/beego/orm"
	"RentHouseWeb/rentHouseWeb/models"
	"strconv"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostLogin(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("用户登录服务,PostLogin..")

	// 1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	// 2.查询数据
	var user models.User
	o := orm.NewOrm()
	qs := o.QueryTable("User")
	err := qs.Filter("mobile",req.Mobile).One(&user)
	if err != nil {
		fmt.Println("用户名查询失败",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 3.判断密码是否正确
	if utils.Getmd5String(req.Password)!= user.Password_hash {
		fmt.Println("密码错误",utils.Getmd5String(req.Password), user.Password_hash)
		rsp.Errno = utils.RECODE_PWDERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 4.生成sessionId
	sessionId := utils.Getmd5String(req.Mobile + req.Password + strconv.Itoa(int(time.Now().UnixNano())))

	// 5.连接redis
	bm,err := utils.RedisOpen(utils.G_server_name,utils.G_redis_addr,
		utils.G_redis_port,utils.G_redis_dbnum)
	if err !=nil{
		fmt.Println("连接redis失败",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 6.拼接key,将数据存入redis
	bm.Put(sessionId + "user_id",user.Id,time.Second * 600)
	bm.Put(sessionId + "user_mobile",user.Mobile,time.Second * 600)
	bm.Put(sessionId + "user_name",user.Name,time.Second * 600)

	// 7.返回sessionId
	rsp.SessionId = sessionId

	return nil
}
