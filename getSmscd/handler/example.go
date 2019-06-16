package handler

import (
	"context"

	example "RentHouseWeb/getSmscd/proto/example"
	"fmt"
	"RentHouseWeb/rentHouseWeb/utils"
	"github.com/astaxie/beego/orm"
	"RentHouseWeb/rentHouseWeb/models"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"math/rand"
	"time"
	"github.com/SubmailDem/submail"
	"strconv"
	"strings"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetSmscd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取短信验证码服务,GetSmscd..")

	// 1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	// 2.验证手机号,如果数据库中存在说明当前用户时老用户
	//创建数据库句柄
	o := orm.NewOrm()
	//给查询条件赋值
	user := models.User{Mobile: req.Mobile}
	err := o.Read(&user, "Mobile")
	if err == nil {
		fmt.Println("该用户已经注册!", req.Mobile)
		rsp.Errno = utils.RECODE_USERONERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 3.连接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr,
		utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 4.获取图片验证码
	value := bm.Get(req.Uuid)
	fmt.Println(reflect.TypeOf(value), value)
	imageCode, _ := redis.String(value, nil)

	// 5.数据比对,验证图片验证码是否正确
	if req.Text != imageCode {
		fmt.Println("图片验证码错误!", req.Text, imageCode)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	// 6.生成随机数短信验证码
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成四位随机数,1000-9999
	num := r.Intn(8999) + 1000
	fmt.Println("产生的随机数:", num)

	// 7.调用短信接口发送短信
	//发送短信的配置信息
	messageConfig := make(map[string]string)
	//预先创建好的appid
	messageConfig["appid"] = "29672"
	//预先获得的app的key
	messageConfig["appkey"] = "89d90165cbea8cae80137d7584179bdb"
	//加密方式默认
	messageConfig["signtype"] = "md5"

	//messagexsend
	//创建短信发送的句柄
	messagexsend := submail.CreateMessageXSend()
	//短信发送的手机号
	submail.MessageXSendAddTo(messagexsend, req.Mobile)
	//短信发送的模板
	submail.MessageXSendSetProject(messagexsend, "NQ1J94")
	//验证码
	submail.MessageXSendAddVar(messagexsend, "code", strconv.Itoa(num))
	//发送短信的请求
	send := submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageConfig)

	fmt.Println("MessageXSend ", send)

	// 8.验证短信验证码是否发送成功
	bool := strings.Contains(send, "success")
	if bool != true {
		fmt.Println("短息验证码发送失败!")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	// 9.将随机数与手机号码存入redis中方便后续使用
	err = bm.Put(req.Mobile, num, time.Second*300)
	if err != nil {
		fmt.Println("随机数存储失败!")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	return nil
}
