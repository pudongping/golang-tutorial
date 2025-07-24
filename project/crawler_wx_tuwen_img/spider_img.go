package crawler_wx_tuwen_img

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

var (
	reImgURL      = regexp.MustCompile(`cdn_url: '([^']+)'`)                          // 抓取微信图片链接地址
	reContentData = regexp.MustCompile(`window\.__QMTPL_SSR_DATA__=\s*(\{.*?\})\s*;`) // 解析图文类型的一些文字信息
	reTitle       = regexp.MustCompile(`title:'(.*?)'`)                               // 抓取标题
	reDesc        = regexp.MustCompile(`desc:'(.*?)'`)                                // 抓取正文内容
)

const (
	tuWenFilePathConst     = "./wx_tuwen_urls.txt"   // 微信图文链接地址文件
	imgSavePathConst       = "./download"            // 图片保存路径
	wenAnFilePathConst     = "./download/wen_an.txt" // 文案保存地址
	httpClientTimeoutConst = time.Minute * 15        // 网络请求超时时间
)

type CrawlResult struct {
	URL                string   // 需要被抓取的原始链接地址
	Number             int      // 当前子协程的编号
	Err                error    // 抓取过程中出现的错误
	Html               string   // 链接地址对应的抓取内容
	ImgSavePathSuccess []string // 图片存储的硬盘路径地址
	WriteContent       string   // 需要被写入的文字内容
}

type WXCrawler struct {
	HttpClient    *http.Client
	ImgSavePath   string
	WenAnFilePath string
}

type Tool struct {
}

func (t *Tool) HttpGet(client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "创建GET请求失败！")
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		// 当这里的 err 不为空时，res 可能会直接为 nil，因此就不能在这里 close
		return nil, errors.Wrap(err, "发送GET请求失败！")
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close() // 确保在非 200 OK 响应时关闭资源
		return nil, errors.Wrap(errors.Errorf("网络请求失败，错误码为：%d", res.StatusCode), "HTTP状态码不为200")
	}

	return res, nil
}

func (t *Tool) MkdirIfNotExist(path string) error {
	// os.Stat 用于检查目录是否存在，os.IsNotExist 判断错误类型是否是因为文件或目录不存在
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return errors.Wrap(err, "创建目录失败")
		}
		// 这一行代码必须要有，因为此时的 err 本来就不会为 nil 且此时的 err 就为文件或者目录不存在时的 error
		// 避免外层判断遭受干扰
		return nil
	}

	return err
}

func (t *Tool) CreateFileIfNotExist(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return errors.Wrap(err, "创建文件失败")
		}
		defer file.Close()
		return nil
	}

	return err
}

// 抓取每一个链接地址对应的 html 内容
func (c *WXCrawler) FetchWXHTMLContent(wxTuWenUrl string) (string, error) {
	if "" == wxTuWenUrl {
		return "", errors.Wrap(errors.New("链接地址不能为空！"), "被抓取的链接地址不能为空！")
	}
	httpResp, err := new(Tool).HttpGet(c.HttpClient, wxTuWenUrl)
	if err != nil {
		return "", errors.Wrap(err, "httpGet 方法出现错误！")
	}
	if httpResp != nil {
		defer httpResp.Body.Close()
	}

	// 解析 HTML 文档
	doc, err := goquery.NewDocumentFromReader(httpResp.Body)
	if err != nil {
		return "", errors.Wrap(err, "goquery 解析 HTML 文档出现错误")
	}

	html, err := doc.Html()
	if err != nil {
		return "", errors.Wrap(err, "获取 HTML 文档内容出现错误")
	}

	return html, err
}

func (c *WXCrawler) ParseImgUrls(html string) ([]string, error) {
	// 提取 picture_page_info_list 内容块
	reList := regexp.MustCompile(`window\.picture_page_info_list\s*=\s*(\[[\s\S]*?\]);`)
	listMatch := reList.FindStringSubmatch(html)
	if len(listMatch) < 2 {
		return nil, errors.New("未找到 picture_page_info_list 内容")
	}
	listContent := listMatch[1]

	// 提取 cdn_url，排除 watermark_info 下的
	reImg := regexp.MustCompile(`cdn_url:\s*'([^']+)'`)
	reWatermark := regexp.MustCompile(`watermark_info\s*:\s*\{[\s\S]*?cdn_url:\s*'([^']+)'[\s\S]*?\}`)

	// 先找所有 watermark_info 下的 cdn_url
	watermarkUrls := make(map[string]struct{})
	for _, wm := range reWatermark.FindAllStringSubmatch(listContent, -1) {
		if len(wm) > 1 {
			watermarkUrls[wm[1]] = struct{}{}
		}
	}

	// 再找所有 cdn_url，排除水印
	results := reImg.FindAllStringSubmatch(listContent, -1)
	urls := make([]string, 0, len(results))
	for _, v := range results {
		if len(v) > 1 {
			if _, isWatermark := watermarkUrls[v[1]]; !isWatermark {
				urls = append(urls, v[1])
			}
		}
	}
	return urls, nil
}

func (c *WXCrawler) FastDownloadImgFiles(imgUrls []string, num int) (imgFilePaths []string, err error) {
	savePath := fmt.Sprintf("%s/%d", c.ImgSavePath, num) // 以批次分组成不同的文件夹
	// 先检查目录是否存在
	if err = new(Tool).MkdirIfNotExist(savePath); err != nil {
		return nil, errors.Wrap(err, "批量下载图片时，检查目录是否存在出现错误")
	}

	// 并发下载
	wg := &sync.WaitGroup{}
	sem := make(chan struct{}, 10) // 最多同时10个并发下载
	type imgDownRes struct {
		imgFilePath string
		err         error
	}
	filePathChan := make(chan imgDownRes, len(imgUrls))

	for i, imgUrl := range imgUrls {
		wg.Add(1)

		go func(i int, imgUrl, savePath string, wg *sync.WaitGroup) {
			defer wg.Done()
			sem <- struct{}{} // 请求信号量
			// 释放信号量
			defer func() { <-sem }()

			imgFilePath := fmt.Sprintf("%s/%d.jpeg", savePath, i+1)
			// 一张一张的下载图片
			_, err := c.DownloadImgFile(imgUrl, imgFilePath)
			filePathChan <- imgDownRes{
				imgFilePath: imgFilePath,
				err:         err,
			}

		}(i, imgUrl, savePath, wg)

	}

	wg.Wait()
	close(filePathChan)

	var errStr string
	for downloadRes := range filePathChan {
		if downloadRes.err != nil {
			errStr += "Path: " + downloadRes.imgFilePath + " Err: " + downloadRes.err.Error() + " | "
		} else {
			imgFilePaths = append(imgFilePaths, downloadRes.imgFilePath)
		}
	}

	if "" != errStr {
		return nil, errors.Wrap(errors.New(errStr), "批量下载图片时，可能下载某一张图片时出现错误")
	}

	return imgFilePaths, nil
}

// 一张一张的下载图片
func (c *WXCrawler) DownloadImgFile(imgUrl, imgFilePath string) (string, error) {
	httpResp, err := new(Tool).HttpGet(c.HttpClient, imgUrl)
	if err != nil {
		return "", errors.Wrap(err, "一张一张下载图片时，出现错误")
	}
	if httpResp != nil {
		defer httpResp.Body.Close()
	}
	log.Printf("正在下载图片：%s\n", imgFilePath)
	file, err := os.Create(imgFilePath)
	if err != nil {
		return "", errors.Wrap(err, "下载图片时，创建文件失败")
	}
	defer file.Close()
	// 保存文件
	_, err = io.Copy(file, httpResp.Body)
	if err != nil {
		return "", errors.Wrap(err, "保存下载的图片文件失败")
	}
	return imgFilePath, nil
}

func (c *WXCrawler) GetWriteContent(html string, num int) string {
	// 查找匹配的部分
	matches := reContentData.FindStringSubmatch(html)
	if len(matches) <= 1 {
		// 没有匹配到内容
		return ""
	}
	contentStr := matches[1] // 匹配到的内容

	// 提取 title 和 desc 的值
	// 因为提取的 jsonStr 内容中是一定会含有 title 和 desc 字段的，因此以下代码可不用做边界值的判断
	// 这里不能直接通过解析 json 字符串的方式来提取内容，因为这里的内容不是一个合法的 json 字符串，它仅仅是一个 js 代码（尤其注意）
	title := reTitle.FindStringSubmatch(contentStr)[1]
	desc := reDesc.FindStringSubmatch(contentStr)[1]

	content := fmt.Sprintf("第 %d ====> \r\n", num)
	content += "标题： " + title + "\r\n"
	content += "正文内容 --------------- \r\n " + desc + "\r\n ------------- \r\n"

	return content
}

func (c *WXCrawler) WriteWenAnContent(contents []CrawlResult) error {
	if err := (&Tool{}).CreateFileIfNotExist(c.WenAnFilePath); err != nil {
		return errors.Wrap(err, "写入文案时，创建文本文件出现异常")
	}

	// 先按照 number 字段进行从小到大排序
	sort.Slice(contents, func(i, j int) bool {
		return contents[i].Number < contents[j].Number
	})

	result := ""
	for _, content := range contents {
		result += content.WriteContent + "\r\n\r\n"
	}

	if err := ioutil.WriteFile(c.WenAnFilePath, []byte(result), 0644); err != nil {
		return errors.Wrapf(err, "写入 %s 文件时，发生错误：%+v", c.WenAnFilePath, err)
	}

	return nil
}

// 从指定文件中读取所需要抓取的链接地址
func getWXTuWenUrls(tuWenFilePath string) ([]string, error) {
	urlFile, err := os.ReadFile(tuWenFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "获取所有原始链接地址时，读取文件失败")
	}

	if "" == string(urlFile) {
		return nil, errors.Wrap(errors.New("读取文件内容为空，需要填写的地址比如：https://mp.weixin.qq.com/s/zQvrsZGEJ5T7zcSyu_E94A"), "存放链接的文件不能为空！")
	}

	wxTuWenIMGUrls := strings.Split(string(urlFile), "\n")
	if len(wxTuWenIMGUrls) == 0 {
		return nil, errors.Wrap(errors.New("微信图文链接地址必须一行一个"), "链接地址解析错误")
	}

	return wxTuWenIMGUrls, nil
}

func work(i int, wxTuWenIMGUrl string, wg *sync.WaitGroup, wxCrawler *WXCrawler, crawlResultChan chan CrawlResult) {
	defer wg.Done()

	num := i + 1 // 标记每个子协程的序号
	var err error

	crawlRes := CrawlResult{
		URL:    wxTuWenIMGUrl,
		Number: num,
		Err:    nil,
	}
	// 先一个一个的抓取每一个链接地址对应的 html 内容
	html, err := wxCrawler.FetchWXHTMLContent(wxTuWenIMGUrl)
	if err != nil {
		crawlRes.Err = err
		crawlResultChan <- crawlRes
		return
	}
	crawlRes.Html = html
	// 从抓取后的 html 内容中解析出所有的图片链接地址
	imgUrls, err := wxCrawler.ParseImgUrls(html)
	if err != nil {
		crawlRes.Err = err
		crawlResultChan <- crawlRes
		return
	}
	// 批量下载图片
	imgFilePaths, err := wxCrawler.FastDownloadImgFiles(imgUrls, num)
	if err != nil {
		crawlRes.Err = err
		crawlResultChan <- crawlRes
		return
	}
	crawlRes.ImgSavePathSuccess = imgFilePaths
	// 提取想要记录的文本信息
	crawlRes.WriteContent = wxCrawler.GetWriteContent(html, num)

	crawlResultChan <- crawlRes
}

func RunSpiderImg() (err error) {
	start := time.Now()

	wxTuWenIMGUrls, err := getWXTuWenUrls(tuWenFilePathConst) // 从指定文件中读取想要被抓取的链接地址
	if err != nil {
		return errors.Wrap(err, "读取文件，获取被抓取的链接地址出错")
	}

	wg := sync.WaitGroup{}
	wxCrawler := &WXCrawler{
		HttpClient:    &http.Client{Timeout: httpClientTimeoutConst},
		ImgSavePath:   imgSavePathConst,
		WenAnFilePath: wenAnFilePathConst,
	}

	crawlResultChan := make(chan CrawlResult, len(wxTuWenIMGUrls)) // 收集信息
	wg.Add(len(wxTuWenIMGUrls))

	for i, wxTuWenIMGUrl := range wxTuWenIMGUrls {
		go work(i, wxTuWenIMGUrl, &wg, wxCrawler, crawlResultChan)
	}

	// 等待所有子协程完成
	wg.Wait()
	close(crawlResultChan) // 关闭通道

	var spiderResults []CrawlResult
	// 读取通道中所有的数据
	for workRes := range crawlResultChan {
		spiderResults = append(spiderResults, workRes)
	}

	// 处理最终抓取的结果
	for _, item := range spiderResults {
		if item.Err != nil {
			log.Printf("Num: %d ==> Url: %s ==> Err: %+v \n", item.Number, item.URL, item.Err)
		}
	}

	// 将一些文案写入文件中
	if err := wxCrawler.WriteWenAnContent(spiderResults); err != nil {
		return errors.Wrap(err, "将文案写入时，出现异常")
	}

	castTime := time.Since(start)
	log.Printf("总耗时：%s", castTime.String())

	return nil
}
