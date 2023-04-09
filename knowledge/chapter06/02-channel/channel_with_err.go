package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Url  string
	Resp *http.Response
	Err  error
}

func main() {
	var urls = []string{
		"https://httpbin.org/get",
		"https://httpbin.org/get?a=1",
		"https://httpbin.org/get?a=2",
		"https://httpbin.org/get?a=3",
		"https://httpbin.org/get?a=4",
		"https://httpbin.org/get?a=5",
		"https://www.baidu.com",
	}

	results, err := CrawlWithTimeout(urls, time.Second*10)
	if err != nil {
		log.Fatalf("有错误！%v \n", err)
	}

	for _, r := range results {

		if r.Err != nil {
			fmt.Printf("抓取的过程中出现错误，url ==> %s 错误信息为：%v\n", r.Url, r.Err)
			continue
		}

		resBody, err := ioutil.ReadAll(r.Resp.Body)
		if err != nil {
			fmt.Printf("读取返回体出现错误：%v \n", err)
			continue
		}

		fmt.Printf("URL: %s, Status Code: %d \n Response Content: \n %v \n ====> \n", r.Url, r.Resp.StatusCode, string(resBody))
	}

}

func CrawlWithTimeout(urls []string, timeout time.Duration) ([]*Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan *Result, len(urls))

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		// wg.Add(1)
		go func(url string) {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					ch <- &Result{
						Url: url,
						Err: fmt.Errorf("内部出现错误：%v", err),
					}
				}
			}()

			client := &http.Client{Timeout: 5 * time.Second}

			res, err := client.Get(url)
			if err != nil {
				ch <- &Result{
					Url: url,
					Err: err,
				}
				return
			}
			// 如果需要在外层使用 `*http.Response` 那么可不用关闭 res.Body
			// defer res.Body.Close()

			ch <- &Result{url, res, nil}

		}(url)
	}
	// 因此这一段 goroutine 无论是放到 for _, url := range urls 之上，还是之下都一样，
	// 因为只需要保证协程执行完毕之后再去关闭通道即可，但是为了方便理解，这里还是放到下面
	go func() {
		wg.Wait() // 要等待所有的 wg.Done() 执行完毕，才会去执行 close(ch)
		close(ch)
	}()

	// 等待所有 goroutine 完成或超时发生
	var results []*Result
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("程序超时：%s", ctx.Err())
		case result, ok := <-ch:
			// 变量 ok 如果为 true 表示 channel 没有关闭，如果为 false 表示 channel 已经关闭
			if !ok {
				return results, nil
			}
			results = append(results, result)
		}
	}

}
