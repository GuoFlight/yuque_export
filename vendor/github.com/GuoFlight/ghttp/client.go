package ghttp

import (
	"bytes"
	"github.com/avast/retry-go"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	Method  string
	Host    string
	Url     string
	Header  map[string]string
	ReqBody string
	UrlArgs map[string]string
	Retry
}
type Res struct {
	Body   []byte
	Detail *http.Response
}

func (my *Client) do() (Res, error) {
	client := &http.Client{}
	req, err := http.NewRequest(my.Method, my.Url, strings.NewReader(my.ReqBody))
	if err != nil {
		return Res{}, err
	}

	//Header中添加域名
	if my.Host != "" {
		req.Host = my.Host
	}

	//添加头部
	for i, v := range my.Header {
		req.Header.Add(i, v)
	}

	//添加url参数
	params := req.URL.Query()
	for argK, argV := range my.UrlArgs {
		params.Add(argK, argV)
	}
	req.URL.RawQuery = params.Encode()

	//发起请求
	response, err := client.Do(req)
	if err != nil {
		return Res{}, err
	}
	defer response.Body.Close()

	//获取响应体
	var res Res
	res.Detail = response
	res.Body, err = ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(res.Body))
	if err != nil {
		return Res{}, err
	}
	return res, nil
}
func (my *Client) Do() (Res, error) {
	// 不重试
	if my.Retry.JudgeRetryFunc == nil {
		return my.do()
	}
	// 有重试
	var res Res
	var errRetry, errHttp error
	errRetry = retry.Do(
		func() error {
			res, errHttp = my.do()
			return my.JudgeRetryFunc(&res)
		},
		my.Options...,
	)
	if errHttp != nil {
		return res, errHttp
	} else if errRetry != nil {
		return res, errRetry
	}
	return res, nil
}
