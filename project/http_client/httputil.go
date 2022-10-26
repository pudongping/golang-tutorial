package http_client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

func Get(url string, req, res interface{}, headers map[string]string) (*http.Response, string, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return AnyRequest(context.TODO(), "GET", url, string(reqBody), res, headers)
}

func PostJSON(url string, req, res interface{}, headers map[string]string) (*http.Response, string, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	headers["Content-Type"] = "application/json"

	return AnyRequest(context.TODO(), "POST", url, string(reqBody), res, headers)
}

func PostForm(url string, req url.Values, res interface{}, headers map[string]string) (*http.Response, string, error) {
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	return AnyRequest(context.TODO(), "POST", url, req.Encode(), res, headers)
}

func AnyRequest(ctx context.Context, method, url, req string, res interface{}, headers map[string]string) (responseObj *http.Response, responseBody string, err error) {
	payload := strings.NewReader(req)
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, "", err
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}

	// 请求超时时间为 5 秒钟
	client := &http.Client{Timeout: time.Second * 5}
	// response, err := client.Do(request)
	response, err := ctxhttp.Do(ctx, client, request)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	if err = json.Unmarshal(resBody, res); err != nil {
		return nil, "", err
	}

	return response, string(resBody), nil
}
