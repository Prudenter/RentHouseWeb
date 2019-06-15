package handler

import (
	"context"
	example "RentHouseWeb/getImageCd/proto/example"
	"RentHouseWeb/rentHouseWeb/utils"
	"github.com/afocus/captcha"
	"image/color"
	"fmt"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetImageCd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取图片验证码服务,GetImageCd...")
	// 1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	// 2.生成随机数与图片
	//创建一个句柄
	cap := captcha.New()
	//设置字体
	if err := cap.SetFont("comic.ttf"); err != nil {
		//抛出异常
		panic(err.Error())
	}
	//设置突破大小
	cap.SetSize(90, 41)
	//设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	//设置前景色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 200})
	//设置背景色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	//随即生成图片与验证码
	img, str := cap.Create(4, captcha.NUM)
	//打印字符串
	println("随机验证码为:", str)

	// 3.连接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr,
		utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 4.存入数据
	bm.Put(req.Uuid, str, time.Second*300)

	// 5.将图片拆分,返回数据给web端
	rgba := *((*img).RGBA)
	// 赋值数据给proto
	// pix
	for _, value := range rgba.Pix {
		rsp.Pix = append(rsp.Pix, uint32(value))
	}
	// stride
	rsp.Stride = int64(rgba.Stride)
	// point
	rsp.Min = &example.ResponsePoint{
		X: int64(rgba.Rect.Min.X),
		Y: int64(rgba.Rect.Min.Y),
	}
	rsp.Max = &example.ResponsePoint{
		X: int64(rgba.Rect.Max.X),
		Y: int64(rgba.Rect.Max.Y),
	}
	return nil
}
