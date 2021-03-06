/*
下载器，接受请求并返回请求
*/

package downloader

import (
	"basic"
	"fmt"
	"net/http"
)

type GenDownloader interface {
	Download(req *basic.Request) *basic.Response
}

type Downloader struct {
	//用于处理http请求
	client *http.Client
}

func NewDownloader() GenDownloader {
	return &Downloader{&http.Client{}}
}

//接受构造请求，返回构造响应
func (self *Downloader) Download(req *basic.Request) *basic.Response {
	for k, v := range basic.Config.HttpHeader {
		fmt.Println(k, v)
		req.GetReq().Header.Set(k, v)
	}
	httpRes, err := self.client.Do(req.GetReq())
	if err != nil {
		return nil
	}
	response := basic.NewResponse(httpRes, req.GetIndex())
	return response
}
