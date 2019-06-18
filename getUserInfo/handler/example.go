package handler

import (
	"context"
	example "RentHouseWeb/getUserInfo/proto/example"
	"fmt"
	"RentHouseWeb/rentHouseWeb/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"RentHouseWeb/rentHouseWeb/models"
	"strconv"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取用户信息服务,GetUserInfo..")

	// 1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno )

	// 2.获取sessionId
	sessionId := req.SessionId

	// 3.连接redis
	bm ,err :=utils.RedisOpen(utils.G_server_name ,utils.G_redis_addr ,utils.G_redis_port ,utils.G_redis_dbnum)
	if err!=nil{
		fmt.Println("redis 连接失败",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno )
		return nil
	}

	// 4.拼接key,查询user_id
	value := bm.Get(sessionId + "user_id")
	value_int,_ := redis.Int(value,nil)

	// 5.连接数据库
	o := orm.NewOrm()
	user := models.User{Id:value_int}

	// 6.查询数据库
	err = o.Read(&user)
	err =o.Read(&user)
	if err !=nil{
		fmt.Println("用户数据查询失败",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno )
		return nil
	}

	// 7.返回数据
	rsp.UserId = strconv.Itoa(user.Id)
	rsp.Name = user.Name
	rsp.Mobile = user.Mobile
	rsp.RealName = user.Real_name
	rsp.IdCard = user.Id_card
	rsp.AvatarUrl = user.Avatar_url

	return nil
}
