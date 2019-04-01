package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//获取要爬取的网站首页面
	resp,err:=http.Get("http://www.zhenai.com/zhenghun")

	if err!=nil{
		panic(err)
	}

	//最后关闭返回的Body
	defer resp.Body.Close()

	contents,err := ioutil.ReadAll(resp.Body)

	if err !=nil{
		panic(err)
	}

	fmt.Printf("%s",contents);

}
