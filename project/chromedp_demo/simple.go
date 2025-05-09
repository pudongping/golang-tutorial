package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// MacOS Chrome Path
const ChromePath = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"

// Windows Chrome Path
// const ChromePath = "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(ChromePath),    // 设置 Chrome 的路径
		chromedp.Flag("headless", false), // 设置为无头模式
	)

	// 连接本地 Chrome
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// 连接远程 Chrome
	// allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://192.168.0.105:9222/")
	// defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithDebugf(log.Printf))
	defer cancel()

	var title string
	err := chromedp.Run(ctx,
		// chromedp.Emulate(device.IPhone13ProMax),
		chromedp.Navigate("https://www.baidu.com"),
		chromedp.Evaluate(`document.title`, &title),
	)
	if err != nil {
		log.Fatalf("Failed to run chromedp: %v \n", err)
	}

	fmt.Printf("Page title: %s\n", title)

	err = chromedp.Run(ctx,
		chromedp.Navigate("https://www.baidu.com"),
		// 等待页面的 footer 元素可见（即，页面已经加载完成后）
		// chromedp.WaitVisible(`body > footer`),
		chromedp.Click(`input#kw`),
		chromedp.SendKeys(`input#kw`, "chromedp可以做什么？"),
		chromedp.Submit(`input#su`),
	)
	if err != nil {
		log.Fatalf("Failed to run chromedp: %v \n", err)
	}

	time.Sleep(time.Second * 2)

	// 截图
	// fullScreenshot(ctx)

	// 保存 PDF
	// savePDF(ctx)

}

func fullScreenshot(ctx context.Context) {
	// 截图
	var res []byte
	err := chromedp.Run(ctx,
		// chromedp.Screenshot(`body`, &res, chromedp.NodeVisible),
		chromedp.Tasks{
			chromedp.FullScreenshot(&res, 90),
		},
	)
	if err != nil {
		log.Fatalf("截图Failed to run chromedp: %v \n", err)
	}
	if err = os.WriteFile("./baidu.png", res, 0644); err != nil {
		log.Fatalf("Failed to save screenshot: %v \n", err)
	}

}

func savePDF(ctx context.Context) {
	// 保存 PDF
	var res []byte
	err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			res = buf
			return nil
		}),
	)
	if err != nil {
		log.Fatalf("Failed to run chromedp: %v \n", err)
	}
	if err = os.WriteFile("./baidu.pdf", res, 0644); err != nil {
		log.Fatalf("Failed to save screenshot: %v \n", err)
	}
}
