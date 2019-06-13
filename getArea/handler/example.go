package handler

import (
	"context"
	example "RentHouseWeb/getArea/proto/example"
	"fmt"
	"RentHouseWeb/rentHouseWeb/utils"
	"github.com/astaxie/beego/orm"
	"RentHouseWeb/rentHouseWeb/models"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取区域信息服务,GetArea...")
	// 1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	// 2.查询数据库
	o := orm.NewOrm()
	// 定义接收数据的容器
	var areas []models.Area
	//设置查询条件
	qs := o.QueryTable("Area")
	//查询全部
	n, err := qs.All(&areas)
	if err != nil {
		fmt.Println("查询数据库错误", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if n == 0 {
		fmt.Println("无数据", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 3.将查询到的数据转化类型
	for _, value := range areas {
		area := example.ResponseAddress{AId: int32(value.Id), AName: string(value.Name)}
		// 4.数据返回给web端
		rsp.Data = append(rsp.Data, &area)
	}
	return nil
}
