package handler

import (
	"context"
	example "RentHouseWeb/getArea/proto/example"
	"fmt"
	"RentHouseWeb/rentHouseWeb/utils"
	"github.com/astaxie/beego/orm"
	"RentHouseWeb/rentHouseWeb/models"
	"encoding/json"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取区域信息服务,GetArea...")
	// 1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	// 2.调用工具类函数,获取存入redis中的数据
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	// 定义区域信息存入redis的key值
	key := "area_info"
	// 先从redis中获取数据
	area_info_value := bm.Get(key)
	// 校验数据
	if area_info_value != nil {
		fmt.Println("已从redis中获取到数据,准备发送给web端..")
		// 接收数据
		var areas []models.Area
		// 解码
		err := json.Unmarshal(area_info_value.([]byte), &areas)
		if err != nil {
			fmt.Println("redis数据解码错误", err)
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(rsp.Errno)
			return nil
		}
		// 循环转换数据发送给web
		for _, value := range areas {
			area := example.ResponseAddress{AId: int32(value.Id), AName: string(value.Name)}
			// 数据返回给web端
			rsp.Data = append(rsp.Data, &area)
		}
		return nil
	}

	// 3.如果redis中获取不到数据,则需要查询数据库
	o := orm.NewOrm()
	// 定义接收数据的容器
	var areas []models.Area
	//设置查询条件
	qs := o.QueryTable("Area")
	//查询全部
	n, err := qs.All(&areas)
	if err != nil {
		fmt.Println("查询数据库错误!", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if n == 0 {
		fmt.Println("无数据!", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 4.将查询到的数据存入redis中,方便后续使用
	// 对数据进行json编码
	area_info_json, _ := json.Marshal(areas)
	//存入数据
	bm.Put(key, area_info_json, time.Second*7200)
	if err != nil {
		fmt.Println("存入redis数据失败!", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 5.将查询到的数据转化类型
	for _, value := range areas {
		area := example.ResponseAddress{AId: int32(value.Id), AName: string(value.Name)}
		// 6.数据返回给web端
		rsp.Data = append(rsp.Data, &area)
	}
	return nil
}
