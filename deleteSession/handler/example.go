package handler

import (
	"context"
	example "RentHouseWeb/deleteSession/proto/example"
	"fmt"
	"RentHouseWeb/rentHouseWeb/utils"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) DeleteSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("退出登录服务,DeleteSession..")
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

	// 3.拼接key,删除session登录信息
	sessionId := req.SessionId
	bm.Delete(sessionId + "user_id")
	bm.Delete(sessionId + "user_mobile")
	bm.Delete(sessionId + "user_name")

	return nil
}
