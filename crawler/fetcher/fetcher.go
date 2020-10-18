package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) { //定义个抓取函数，输入是一个url，输出是一个字节数组和错误信息
	<-rateLimiter              //给爬虫设置一个运行速度，防止执行过快
	resp, err := http.Get(url) //使用http包Get该url链接的response
	if err != nil {
		return nil, err //如果get错误，返回一个空值和错误信息
	}
	defer resp.Body.Close() //函数执行完毕后关闭response

	if resp.StatusCode != http.StatusOK { //判断response的状态码是否正常
		return nil, fmt.Errorf("wrong statuscode: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body) //使用ioutil中的readall读取response的body部分，返回一个字节数组和错误信息

}
