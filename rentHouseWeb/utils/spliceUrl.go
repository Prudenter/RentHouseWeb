/**
* @Author: ASlowPerson  
* @Date: 19-6-12 下午2:26
*/

package utils

import (
	"github.com/astaxie/beego/cache"
	"encoding/json"
	"fmt"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/astaxie/beego/cache/redis"
)

/*
	url请求拼接函数
	将http://IP:PORT + url进行拼接
*/
func SpliceUrl(url string) (reqUrl string) {
	reqUrl = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url
	return
}

/*
	redis连接函数
*/
func RedisOpen(server_name, redis_addr, redis_port, redis_dbnum string) (bm cache.Cache, err error) {
	redis_config_map := map[string]string{
		"key":   server_name,
		"conn":  redis_addr + ":" + redis_port,
		"dbNum": redis_dbnum,
	}
	// 将配置信息的map转化为json
	redis_config, _ := json.Marshal(redis_config_map)
	// 连接redis数据库
	bm, err = cache.NewCache("redis", string(redis_config))
	if err != nil {
		fmt.Println("连接redis错误", err)
		return nil, err
	}
	return bm, nil
}
