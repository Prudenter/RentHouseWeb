package handler

import (
	"context"
	example "RentHouseWeb/getSession/proto/example"
	"fmt"
	"RentHouseWeb/rentHouseWeb/utils"
	"github.com/garyburd/redigo/redis"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("登录检查服务,GetSession...")

	// 1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno )

	// 2.连接redis
	bm ,err :=utils.RedisOpen(utils.G_server_name,utils.G_redis_addr ,utils.G_redis_port ,utils.G_redis_dbnum)
	if err!=nil{
		fmt.Println("连接redis失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno )
		return nil
	}

	// 3.接收web发送过来的sessionId
	sessionId := req.SessionId

	// 4.拼接对应key
	nameKey := sessionId + "user_name"

	// 5.查询name
	value := bm.Get(nameKey)
	name,_ := redis.String(value,nil)

	// 6.返回数据到web端
	rsp.Name = name

	return nil
}
