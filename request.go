package go_curl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	cli         *http.Client
	req         *http.Request
	res         *http.Response
	Url         string            `json:"url 请求地址"`
	Method      string            `json:"method 方法"`
	UserAgent   string            `json:"user_agent"`
	TimeOut     time.Duration     `json:"time_out 超时"`
	Cookies     map[string]string `json:"cookies"`
	ContentType string            `json:"content_type"`
	Proxy       string
	Data        string `json:"传输数据"`
	Header      map[string]string
}

type Response struct {
	Header     map[string][]string
	Body       string
	Status     string
	StatusCode int
}

func NewCurl() *Request {
	req := &Request{}
	req.TimeOut = time.Duration(5 * time.Second)

	return req
}
func (this *Request) SetUrl(url string) *Request {
	this.Url = url
	return this
}
func (this *Request) SetData(data string) *Request {
	this.Data = data
	return this
}
func (this *Request) SetMethod(method string) *Request {
	this.Method = strings.ToUpper(method)
	return this
}
func (this *Request) SetUserAgent(string2 string) *Request {
	this.UserAgent = string2
	return this
}
func (this *Request) SetTimeOut(timeInt int64) *Request {
	this.TimeOut = time.Duration(timeInt)
	return this
}
func (this *Request) SetHeader(headers map[string]string) *Request {
	this.Header = headers
	return this
}
func (this *Request) SetCookies(cookies map[string]string) *Request {
	this.Cookies = cookies
	return this
}
func (this *Request) SetProxy(urlstr string) *Request {
	return this
}
func (this *Request) SetContentType(contentType string) *Request {
	this.ContentType = contentType
	return this
}
func (this *Request) setCookies() error {
	for k, v := range this.Cookies {
		this.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
	return nil
}

func (this *Request, ) send() (*Response, error) {
	cli := &http.Client{}
	cli.Timeout = this.TimeOut * time.Second
	url := this.Url
	method := this.Method
	payload := bytes.NewBufferString(this.Data)
	if req, err := http.NewRequest(method, url, payload); err != nil {
		return nil, err
	} else {
		this.req = req
	}
	this.setCookies()
	this.req.Header.Set(strings.ToUpper("content-type"), this.ContentType)
	this.req.Header.Set(strings.ToUpper("User-Agent"), this.UserAgent)
	for k, v := range this.Header {
		this.req.Header.Set(strings.ToUpper(k), v)
	}
	res, err := this.cli.Do(this.req)
	resp := &Response{}
	if err == nil {
		body, err := ioutil.ReadAll(res.Body)
		if err == nil {
			resp = &Response{Header: (res.Header), Body: string(body), Status: res.Status, StatusCode: res.StatusCode}
		}
	} else {
		panic(err)
	}
	this.res = res
	return resp, err
}
