/**
* @Author: ASlowPerson  
* @Date: 19-6-12 下午2:26
*/

package utils
/*
	url请求拼接函数
	将http://IP:PORT + url进行拼接
*/
func SpliceUrl(url string) (reqUrl string) {
	reqUrl = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url
	return
}
