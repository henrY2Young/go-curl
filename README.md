# go-curl

### 使用方法

```
curl := NewCurl()
	curl.cli = &http.Client{}
	//设置cookie
	cookies := map[string]string{
		"cookie": "set_cookie",
	}
	//设置传递数据,和Content-Type去配合使用，传递json或者string类型
	//data:= map[string]interface{}{
	//	"name":"abc",
	//	"sex":"man",
	//}
	//data1,_ := json.Marshal(data)
	data2 := "name=abc&sex=man"
    //设置请求头
	header := map[string]string{
		"Connection": "keep-alive",
	}
	//设置代理
	proxy:="http://119.5.0.75:808"
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
	res, _ := curl.SetUrl("https://www.google.com").SetProxy(proxy).SetHeader(header).SetUserAgent(userAgent).SetMethod("post").SetData(string(data2)).SetContentType("application/x-www-form-urlencoded").SetCookies(cookies).send()
	fmt.Println(res.Body)
 
```
### SetProxy()
设置代理

### SetHeader()

设置请求头

### SetUserAgent()

设置user-agent


### SetMethd()

设置请求方法
### SetData()
设置请求值

### SetContentType()

设置content-type 这值和请求的data配合，确保请求端能够获取到传值


## Response

response中有header,body,status,statuscode等值
