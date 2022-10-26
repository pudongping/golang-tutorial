package http_client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Result struct {
	Args    interface{}       `json:"args"`
	Data    string            `json:"data"`
	Files   interface{}       `json:"files"`
	Form    interface{}       `json:"form"`
	Headers map[string]string `json:"headers"`
	Json    interface{}       `json:"json"`
	Method  string            `json:"method"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

func TestGet(t *testing.T) {
	reqUrl := "https://httpbin.org/anything?name=alex&age=18"
	req := map[string]interface{}{
		"param1": [...]int{1, 2, 3},
		"param2": "hello world",
	}

	// var res map[string]interface{}
	var res Result

	_, body, err := Get(reqUrl, req, &res, nil)
	if err != nil {
		t.Errorf("有错误 %v \n", err)
	}

	// {
	//  "args": {
	//    "age": "18",
	//    "name": "alex"
	//  },
	//  "data": "{\"param1\":[1,2,3],\"param2\":\"hello world\",\"param3\":{\"param31\":\"value31\",\"param32\":12345,\"param33\":[11,33,44]}}",
	//  "files": {},
	//  "form": {},
	//  "headers": {
	//    "Accept-Encoding": "gzip",
	//    "Content-Length": "109",
	//    "Content-Type": "application/json",
	//    "Host": "httpbin.org",
	//    "Token": "123456",
	//    "User-Agent": "App/",
	//    "X-Amzn-Trace-Id": "Root=1-6358a5b4-5f9ceebc28bd7db245d4accc"
	//  },
	//  "json": {
	//    "param1": [
	//      1,
	//      2,
	//      3
	//    ],
	//    "param2": "hello world",
	//    "param3": {
	//      "param31": "value31",
	//      "param32": 12345,
	//      "param33": [
	//        11,
	//        33,
	//        44
	//      ]
	//    }
	//  },
	//  "method": "POST",
	//  "origin": "112.65.11.169",
	//  "url": "https://httpbin.org/anything?name=alex&age=18"
	// }
	fmt.Println(body)
	fmt.Println("=====>")
	spew.Dump(res)
}

func TestPostJSON(t *testing.T) {
	reqUrl := "https://httpbin.org/anything?name=alex&age=18"
	req := map[string]interface{}{
		"param1": [...]int{1, 2, 3},
		"param2": "hello world",
		"param3": map[string]interface{}{
			"param31": "value31",
			"param32": 12345,
			"param33": []int{11, 33, 44},
		},
	}

	var res map[string]interface{}

	headers := map[string]string{
		"User-Agent": "App/",
		"Token":      "123456",
	}

	_, body, err := PostJSON(reqUrl, req, &res, headers)
	if err != nil {
		t.Errorf("有错误 %v \n", err)
	}

	fmt.Println(body)
	fmt.Println("=====>")
	spew.Dump(res)
}

func TestPostForm(t *testing.T) {
	reqUrl := "https://httpbin.org/anything?name=alex&age=18"
	req := url.Values{
		"param1": []string{"value11", "value12", "value13"},
		"param2": []string{"value21"},
	}

	var res map[string]interface{}

	headers := map[string]string{
		"User-Agent": "App/",
		"Token":      "123456",
	}

	_, body, err := PostForm(reqUrl, req, &res, headers)
	if err != nil {
		t.Errorf("有错误 %v \n", err)
	}

	// {
	//  "args": {
	//    "age": "18",
	//    "name": "alex"
	//  },
	//  "data": "",
	//  "files": {},
	//  "form": {
	//    "param1": [
	//      "value11",
	//      "value12",
	//      "value13"
	//    ],
	//    "param2": "value21"
	//  },
	//  "headers": {
	//    "Accept-Encoding": "gzip",
	//    "Content-Length": "59",
	//    "Content-Type": "application/x-www-form-urlencoded",
	//    "Host": "httpbin.org",
	//    "Token": "123456",
	//    "User-Agent": "App/",
	//    "X-Amzn-Trace-Id": "Root=1-6358a4da-33ec2e5104331c413171db1b"
	//  },
	//  "json": null,
	//  "method": "POST",
	//  "origin": "112.65.11.169",
	//  "url": "https://httpbin.org/anything?name=alex&age=18"
	// }
	fmt.Println(body)
	fmt.Println("=====>")
	spew.Dump(res)
}

func TestAnyRequest(t *testing.T) {
	// 如果将时间设置稍微短一点，比如说 2 秒，那么可能 delete 请求会成功，但是 put 请求可能会因为上下文超时而中断掉
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	reqUrl := "https://httpbin.org/anything/123"
	var err error

	fmt.Println("DELETE start")
	var deleteRes map[string]interface{}
	_, deleteBody, err := AnyRequest(ctx, "DELETE", reqUrl, "", &deleteRes, nil)
	if err != nil {
		t.Errorf("DELETE 有错误 %v \n", err)
	}
	// {
	//  "args": {},
	//  "data": "",
	//  "files": {},
	//  "form": {},
	//  "headers": {
	//    "Accept-Encoding": "gzip",
	//    "Host": "httpbin.org",
	//    "User-Agent": "Go-http-client/2.0",
	//    "X-Amzn-Trace-Id": "Root=1-6358db58-76062ef533376565790e1d8a"
	//  },
	//  "json": null,
	//  "method": "DELETE",
	//  "origin": "112.65.11.169",
	//  "url": "https://httpbin.org/anything/123"
	// }
	fmt.Println(deleteBody)
	fmt.Println("DELETE =====>")
	spew.Dump(deleteRes)
	fmt.Println("DELETE end")

	fmt.Println("============================>")

	fmt.Println("PUT start")
	var putRes map[string]interface{}
	putReq := map[string]interface{}{
		"param1": []int{11, 22, 33},
		"param2": "value2",
		"param3": 123,
	}
	putPayload, err := json.Marshal(putReq)
	if err != nil {
		t.Errorf("PUT 有错误 %v \n", err)
		return
	}
	putHeaders := map[string]string{
		"Content-Type": "application/json",
		"Token":        "1234",
	}
	_, putBody, err := AnyRequest(ctx, "PUT", reqUrl, string(putPayload), &putRes, putHeaders)
	if err != nil {
		t.Errorf("PUT 有错误 %v \n", err)
	}
	// {
	//  "args": {},
	//  "data": "{\"param1\":[11,22,33],\"param2\":\"value2\",\"param3\":123}",
	//  "files": {},
	//  "form": {},
	//  "headers": {
	//    "Accept-Encoding": "gzip",
	//    "Content-Length": "52",
	//    "Content-Type": "application/json",
	//    "Host": "httpbin.org",
	//    "Token": "1234",
	//    "User-Agent": "Go-http-client/2.0",
	//    "X-Amzn-Trace-Id": "Root=1-6358de6f-44d849be16f5a4c477b07616"
	//  },
	//  "json": {
	//    "param1": [
	//      11,
	//      22,
	//      33
	//    ],
	//    "param2": "value2",
	//    "param3": 123
	//  },
	//  "method": "PUT",
	//  "origin": "112.65.11.169",
	//  "url": "https://httpbin.org/anything/123"
	// }
	fmt.Println(putBody)
	fmt.Println("PUT =====>")
	spew.Dump(putRes)
	fmt.Println("PUT end")
}
