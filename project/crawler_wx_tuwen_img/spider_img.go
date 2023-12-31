package crawler_wx_tuwen_img

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func httpGet(client *http.Client, url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("创建请求失败，错误信息为：%s", err.Error())
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("网络请求失败，错误信息为：%s", err.Error())
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("网络请求失败，错误码为：%d", res.StatusCode)
	}

	return res
}

func fetchImgUrls(client *http.Client, wxTuWenUrl string) []string {
	if "" == wxTuWenUrl {
		log.Fatal("wxTuWenUrl is empty")
	}
	httpResp := httpGet(client, wxTuWenUrl)
	defer httpResp.Body.Close()

	// 解析 HTML 文档，提取图片链接
	doc, err := goquery.NewDocumentFromReader(httpResp.Body)
	if err != nil {
		log.Fatalf("解析 HTML 文档出现错误，错误信息为：%s", err.Error())
	}

	html, err := doc.Html()
	if err != nil {
		log.Fatalf("获取 HTML 文档内容出现错误，错误信息为：%s", err.Error())
	}

	// 提取图片链接
	re := regexp.MustCompile(`cdn_url: '([^']+)'`)
	// FindAllStringSubmatch返回所有正则表达式的匹配项
	// 每个匹配项都是由原始文本和各个分组匹配项组成的slice
	// 因为这里使用了([^']+)'进行了一次分组匹配而我们想要的就是这个分组匹配的内容 所以使用了v[1]
	results := re.FindAllStringSubmatch(html, -1)
	if (len(results)) == 0 {
		log.Fatalf("图片链接提取失败")
	}
	urls := make([]string, 0, len(results))
	for _, v := range results {
		urls = append(urls, v[1]) // 输出匹配到的cdn_url的值
	}

	return urls
}

func mkdirIfNotExist(path string) {
	// os.Stat 用于检查目录是否存在，os.IsNotExist 判断错误类型是否是因为文件或目录不存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Fatalf("创建目录失败，错误信息为：%s", err.Error())
		}
	}
}

func downloadImgFile(client *http.Client, savePath, imgUrl, fileName string) {
	httpResp := httpGet(client, imgUrl)
	defer httpResp.Body.Close()
	filePath := savePath + "/" + fileName
	fmt.Printf("正在下载图片：%s\n", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("创建文件失败，错误信息为：%s", err.Error())
	}
	defer file.Close()
	// 保存文件
	_, err = io.Copy(file, httpResp.Body)
	if err != nil {
		log.Fatalf("保存文件失败，错误信息为：%s", err.Error())
	}
}

func getWXTuWenUrls(tuWenFilePath string) []string {
	urlFile, err := os.ReadFile(tuWenFilePath)
	if err != nil {
		log.Fatalf("读取文件失败，错误信息为：%s", err.Error())
	}
	if "" == string(urlFile) {
		log.Fatalf("读取文件内容为空，地址比如：https://mp.weixin.qq.com/s/zQvrsZGEJ5T7zcSyu_E94A")
	}
	wxTuWenIMGUrls := strings.Split(string(urlFile), "\n")
	if len(wxTuWenIMGUrls) == 0 {
		log.Fatalf("微信图文链接地址必须一行一个")
	}

	return wxTuWenIMGUrls
}

func fastDownloadImgFiles(client *http.Client, savePath string, imgUrls []string) {
	mkdirIfNotExist(savePath)
	wg := sync.WaitGroup{}
	for i, imgUrl := range imgUrls {
		wg.Add(1)
		go func(i int, imgUrl, savePath string, wg *sync.WaitGroup, client *http.Client) {
			defer wg.Done()
			downloadImgFile(client, savePath, imgUrl, fmt.Sprintf("%d.jpeg", i+1))
		}(i, imgUrl, savePath, &wg, client)
	}
	wg.Wait()
}

func Run() {
	tuWenFilePath := "./wx_tuwen_urls.txt" // 微信图文链接地址文件
	savePath := "./download"               // 图片保存路径

	start := time.Now()

	wxTuWenIMGUrls := getWXTuWenUrls(tuWenFilePath)
	wg := sync.WaitGroup{}
	httpClient := &http.Client{
		Timeout: time.Minute * 5,
	}

	for i, wxTuWenIMGUrl := range wxTuWenIMGUrls {
		wg.Add(1)
		go func(i int, wxTuWenIMGUrl, savePath string, wg *sync.WaitGroup, httpClient *http.Client) {
			defer wg.Done()
			imgUrls := fetchImgUrls(httpClient, wxTuWenIMGUrl) // 获取所需要下载对图片链接地址
			fastDownloadImgFiles(httpClient, fmt.Sprintf("%s/%d", savePath, i+1), imgUrls)
		}(i, wxTuWenIMGUrl, savePath, &wg, httpClient)
	}

	wg.Wait()

	castTime := time.Since(start)
	log.Printf("总耗时：%s", castTime.String())
}
