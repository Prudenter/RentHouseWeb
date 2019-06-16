package handler

import (
	"context"

	example "RentHouseWeb/postRet/proto/example"
	"fmt"
	"RentHouseWeb/rentHouseWeb/utils"
	"github.com/garyburd/redigo/redis"
	"RentHouseWeb/rentHouseWeb/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostRet(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("用户注册服务,PostRet..")

	// 1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	// 2.连接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		fmt.Println("redis连接失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 3.获取短信验证码
	value := bm.Get(req.Mobile)
	value_string, _ := redis.String(value, nil)
	// 4.验证短信验证码是否正确
	if value_string != req.SmsCode {
		fmt.Println("短信验证码错误", value_string, req.SmsCode)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 5.密码加密
	user := models.User{}
	user.Password_hash = utils.Getmd5String(req.Password)

	// 6.插入数据到数据库中
	user.Mobile = req.Mobile
	user.Name = req.Mobile
	o := orm.NewOrm()
	id,err := o.Insert(&user)
	if err != nil {
		fmt.Println("用户注册数据插入数据库失败!",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 7.生成sessionId
	sessionId := utils.Getmd5String(req.Mobile + req.Password + strconv.Itoa(int(time.Now().UnixNano())))
	rsp.SessionId = sessionId

	// 8.通过sessionId将用户登录信息存入到redis
	bm.Put(sessionId + "user_id",id,time.Second * 600)
	bm.Put(sessionId + "user_mobile",user.Mobile,time.Second * 600)
	bm.Put(sessionId + "user_name",user.Name,time.Second * 600)

	return nil
}
