/*
GitHub： https://github.com/olivere/elastic
官方文档示例： https://olivere.github.io/elastic/
下载：go get github.com/olivere/elastic/v7
*/
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/olivere/elastic/v7"
)

var (
	ESClient *elastic.Client
	once     sync.Once
)

// ConnectES 创建 es 连接
func ConnectES(options ...elastic.ClientOptionFunc) {
	once.Do(func() {
		// client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
		var err error
		ESClient, err = elastic.NewClient(options...)
		if err != nil {
			panic(err)
		}
	})
}

func Ping(url string) (*elastic.PingResult, int, error) {
	return ESClient.Ping(url).Do(context.Background())
}

// CreateIndexIfNotExists 索引不存在时，创建索引
// index 索引名称
// mapping 数据类型
func CreateIndexIfNotExists(index, mapping string) error {
	ctx := context.Background()
	exists, err := ESClient.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	info, err := ESClient.CreateIndex(index).BodyString(mapping).Do(ctx)
	// info, err := ESClient.CreateIndex(index).Do(ctx)  // 如果只是想创建索引时，那么就不需要 BodyString() 方法
	if err != nil {
		return err
	}
	if !info.Acknowledged {
		return errors.New(fmt.Sprintf("ES 创建索引 [%s] 失败", index))
	}
	return nil
}

// DeleteIndex 删除索引
// index 索引名称
func DeleteIndex(index string) (*elastic.IndicesDeleteResponse, error) {
	info, err := ESClient.DeleteIndex(index).Do(context.Background())
	if err != nil {
		return nil, err
	}
	if !info.Acknowledged {
		return nil, errors.New(fmt.Sprintf("ES 删除索引 [%s] 失败", index))
	}
	return info, err
}

// CreateDoc 单条添加
// index 索引
// id 文档 id（可以直接为空字符串，当实参为空字符串时，es 会主动随机生成）
// body 需要添加的内容
func CreateDoc(index, id string, body interface{}) (*elastic.IndexResponse, error) {
	client := ESClient.Index().Index(index)
	if "" != id {
		client = client.Id(id)
	}
	return client.BodyJson(body).Do(context.Background())
}

// UpdateDoc 单条更新
// index 索引
// id 记录 id
// body 需要更新的内容 （建议只使用 map[string]interface{} 进行更新指定字段且需要注意 map 中的 key 需要和 es 中的 key 完全匹配，否则 es 会认为新增字段，不要使用 struct 否则会将某些值初始化零值）
func UpdateDoc(index, id string, body interface{}) (*elastic.UpdateResponse, error) {
	return ESClient.Update().Index(index).Id(id).Doc(body).Do(context.Background())
}

// DeleteDoc 删除文档
// index 索引
// id 需要删除的文档记录 id
func DeleteDoc(index, id string) (*elastic.DeleteResponse, error) {
	return ESClient.Delete().Index(index).Id(id).Do(context.Background())
}

// CreateBulkDoc 批量添加
// index 索引
// ids 需要新建的 id 数组（可以为空的字符串切片）
// body 需要添加的内容
// 需要注意：ids 和 body 的顺序要一一对应
func CreateBulkDoc(index string, ids []string, body []interface{}) (*elastic.BulkResponse, error) {
	bulkRequest := ESClient.Bulk()
	for k, v := range body {
		tmp := v
		doc := elastic.NewBulkIndexRequest().Index(index).Doc(tmp)
		if len(ids) > 0 {
			doc = doc.Id(ids[k])
		}
		bulkRequest = bulkRequest.Add(doc)
	}
	return bulkRequest.Do(context.Background())
}

// UpdateBulkDoc 批量更新
// index 索引
// ids 需要更新的 id 数组
// body 需要更新的 id 对应的数据 （建议只使用 []map[string]interface{} 进行更新指定字段且需要注意 map 中的 key 需要和 es 中的 key 完全匹配，否则 es 会认为新增字段，不要使用 struct 否则会将某些值初始化零值）
// 需要注意：ids 和 body 的顺序要一一对应
func UpdateBulkDoc(index string, ids []string, body []interface{}) (*elastic.BulkResponse, error) {
	bulkRequest := ESClient.Bulk()
	for k, v := range body {
		tmp := v
		doc := elastic.NewBulkUpdateRequest().Index(index).Id(ids[k]).Doc(tmp).DocAsUpsert(true)
		bulkRequest = bulkRequest.Add(doc)
	}
	return bulkRequest.Do(context.Background())
}

// DeleteBulkDoc 批量删除
// index 索引
// ids 需要删除的 id 数组
func DeleteBulkDoc(index string, ids []string) (*elastic.BulkResponse, error) {
	bulkRequest := ESClient.Bulk()
	for _, v := range ids {
		tmp := v
		req := elastic.NewBulkDeleteRequest().Index(index).Id(tmp)
		bulkRequest = bulkRequest.Add(req)
	}
	return bulkRequest.Do(context.Background())
}

// FirstDoc 通过 id 取出数据
// index 索引
// id 需要取的文档记录 id
func FirstDoc(index, id string) (*elastic.GetResult, error) {
	return ESClient.Get().Index(index).Id(id).Do(context.Background())
}

type RcpGoodsImgChecksES struct {
	AppName     int    `json:"app_name"`
	GoodsId     string `json:"goods_id"`
	SiteId      int    `json:"site_id"`
	CheckStatus int    `json:"check_status"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}

const RcpGoodsImgChecksESIndex = "rcp_goods_img_checks"

const RcpGoodsImgChecksESMapping = `
{
	"mappings":{
		"properties":{
			"app_name":{
				"type": "integer"
			},
			"goods_id":{
				"type": "keyword"
			},
			"site_id":{
				"type": "keyword"
			},
			"check_status":{
				"type": "integer"
			},
			"created_at":{
				"type": "date"
			},
			"updated_at":{
				"type": "date"
			}
		}
	}
}`

const esUrl = "http://127.0.0.1:9200"

func main() {
	var err error

	// 创建 es 连接
	ConnectES(
		// 如果 es 是通过 docker 安装，如果不设置 `elastic.SetSniff(false)` 那么则会报错
		elastic.SetSniff(false),            // 允许指定弹性是否应该定期检查集群，默认为 true, 会把请求 http://ip:port/_nodes/http，并将其返回的 publish_address 作为请求路径
		elastic.SetURL([]string{esUrl}...), // 服务地址
		elastic.SetBasicAuth("", ""),       // 设置认证账号和密码
		// elastic.SetHealthcheckInterval(time.Second*5), // 心跳检查，间隔时间
		// elastic.SetGzip(true),                         // 启用 gzip 压缩
	)

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := Ping(esUrl)
	if err != nil {
		// Handle error
		panic(err)
	}
	// Elasticsearch returned with code 200 and version 7.9.3
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// 直接打印出 es 版本号
	esVersion, err := ESClient.ElasticsearchVersion(esUrl)
	if err != nil {
		panic(err)
	}
	// Elasticsearch version 7.9.3
	fmt.Printf("Elasticsearch version %s\n", esVersion)

	// 删除索引
	// testDeleteIndex()

	// 判断索引是否存在，如果不存在时，则创建
	err = CreateIndexIfNotExists(RcpGoodsImgChecksESIndex, RcpGoodsImgChecksESMapping)
	if err != nil {
		panic(err)
	}

	// 创建文档
	testCreateDoc()

	// 通过文档 id 的形式更新文档
	// testUpdateDoc()

	// 通过 Script 方式更新文档（单字段更新，借助文档 id 更新）
	// testUpdateDocScript()

	// 通过条件 Script 方式更新文档（单字段更新，根据查询条件批量更新）
	// testUpdateDocScriptQuery()

	// 通过文档 id 查找文档
	// testFirstDoc()

	// 通过文档 id 删除文档
	// testDeleteDoc()

	// 批量创建
	// testCreateBulkDoc()

	// 批量更新
	// testUpdateBulkDoc()

	// term
	// testTermQuery()
	// terms
	// testTermsQuery()
	// range
	// testRangeQuery()
	// match_all
	// testMatchAllQuery()
	// match
	// testMatchQuery()
	// match_phrase
	// testMatchPhraseQuery()
	// match_phrase_prefix
	// testMatchPhrasePrefixQuery()
	// regexp
	// testRegexpQuery()

	// bool query
	// testBoolQuery()

	// 分页查询，并排序
	// testPageSort()
	// 多字段排序
	// testMultiFieldSort()
	// 返回指定字段
	// testFetchSource()

	// 通过文档 id 批量删除
	// testDeleteBulkDoc()

	// 按照条件删除
	// testDeleteByQuery()

}

func PrintQuery(src interface{}) {
	fmt.Println("开始打印参数 ====>")
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println("打印参数结束 ====>")
}

func querySearch(query elastic.Query) {
	if querySrc, err := query.Source(); err == nil {
		PrintQuery(querySrc)
	}
	queryRet, err := ESClient.Search().Index(RcpGoodsImgChecksESIndex).Query(query).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("查询到的结果总数为 %v \n", queryRet.TotalHits())
	for _, v := range queryRet.Hits.Hits {
		var tmp RcpGoodsImgChecksES
		json.Unmarshal(v.Source, &tmp)
		fmt.Printf("已经命中查询的数据为 ==> %+v \n %+v \n\n", v.Id, tmp)
	}
}

func testDeleteIndex() {
	// 删除索引
	deleteIndexRet, err := DeleteIndex(RcpGoodsImgChecksESIndex)
	if err != nil {
		panic(err)
	}
	// deleteIndexRet  ==> &{Acknowledged:true}
	fmt.Printf("deleteIndexRet  ==> %+v \n\n", deleteIndexRet)
}

func testCreateDoc() {
	// 创建文档
	now := time.Now().Unix()

	createDocRet, err := CreateDoc(RcpGoodsImgChecksESIndex, "2_18_alex111", RcpGoodsImgChecksES{
		AppName:     2,
		GoodsId:     "alex111",
		SiteId:      18,
		CheckStatus: 1,
		CreatedAt:   int(now),
		UpdatedAt:   int(now),
	})
	if err != nil {
		panic(err)
	}
	// CreateDoc ==> &{Index:rcp_goods_img_checks Type:_doc Id:2_18_alex111 Version:1 Result:created Shards:0xc00020c2c0 SeqNo:0 PrimaryTerm:1 Status:0 ForcedRefresh:false}
	fmt.Printf("CreateDoc ==> %+v \n\n", createDocRet)
}

func testUpdateDoc() {
	// 通过文档 id 的形式更新文档
	updateDocRet, err := UpdateDoc(RcpGoodsImgChecksESIndex, "2_18_alex111", map[string]interface{}{
		"check_status": 2,
		"updated_at":   int(time.Now().Unix()),
	})
	if err != nil {
		panic(err)
	}
	// UpdateDoc ==> &{Index:rcp_goods_img_checks Type:_doc Id:2_18_alex111 Version:2 Result:updated Shards:0xc0002bc280 SeqNo:1 PrimaryTerm:1 Status:0 ForcedRefresh:false GetResult:<nil>}
	fmt.Printf("UpdateDoc ==> %+v \n\n", updateDocRet)
}

func testUpdateDocScript() {
	// 通过 Script 方式更新文档（单字段更新，借助文档 id 更新）
	updateDocScript, err := ESClient.Update().
		Index(RcpGoodsImgChecksESIndex).
		Id("2_18_alex111").
		Script(elastic.NewScript("ctx._source.site_id=11")).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	// updateDocScript  ==> &{Index:rcp_goods_img_checks Type:_doc Id:2_18_alex111 Version:3 Result:updated Shards:0xc000098280 SeqNo:2 PrimaryTerm:1 Status:0 ForcedRefresh:false GetResult:<nil>}
	fmt.Printf("updateDocScript  ==> %+v \n\n", updateDocScript)
}

func testUpdateDocScriptQuery() {
	// 通过条件 Script 方式更新文档（单字段更新，根据查询条件批量更新字段）
	updateDocScriptQuery, err := ESClient.UpdateByQuery(RcpGoodsImgChecksESIndex).
		Query(elastic.NewTermQuery("goods_id", "alex111")).
		Script(elastic.NewScript("ctx._source.check_status=23")).
		ProceedOnVersionConflict().Do(context.Background())
	if err != nil {
		panic(err)
	}
	// updateDocScriptQuery  ==> &{Header:map[] Took:47 SliceId:<nil> TimedOut:false Total:2 Updated:2 Created:0 Deleted:0 Batches:1 VersionConflicts:0 Noops:0 Retries:{Bulk:0 Search:0} Throttled: ThrottledMillis:0 RequestsPerSecond:-1 Canceled: ThrottledUntil: ThrottledUntilMillis:0 Failures:[]}
	fmt.Printf("updateDocScriptQuery  ==> %+v \n\n", updateDocScriptQuery)
}

func testFirstDoc() {
	// 通过文档 id 查找文档
	firstDocRet, err := FirstDoc(RcpGoodsImgChecksESIndex, "2_18_alex111")
	if err != nil {
		panic(err)
	}
	if firstDocRet.Found { // 表示找到了数据
		// FirstDoc ==>  &{Index:rcp_goods_img_checks Type:_doc Id:2_18_alex111 Uid: Routing: Parent: Version:0xc000282a10 SeqNo:0xc000282a18 PrimaryTerm:0xc000282a20 Source:[123 34 97 112 112 95 110 97 109 101 34 58 50 44 34 117 112 100 97 116 101 100 95 97 116 34 58 49 54 54 48 53 55 57 52 56 54 44 34 115 105 116 101 95 105 100 34 58 49 56 44 34 103 111 111 100 115 95 105 100 34 58 34 97 108 101 120 49 49 49 34 44 34 99 114 101 97 116 101 100 95 97 116 34 58 49 54 54 48 53 55 57 52 56 54 44 34 99 104 101 99 107 95 115 116 97 116 117 115 34 58 50 51 125] Found:true Fields:map[] Error:<nil>} Source: {"app_name":2,"updated_at":1660579486,"site_id":18,"goods_id":"alex111","created_at":1660579486,"check_status":23}
		fmt.Printf("FirstDoc ==>  %+v Source: %+v \n\n", firstDocRet, string(firstDocRet.Source))
	}
}

func testDeleteDoc() {
	// 通过文档 id 删除文档
	deleteDocRet, err := DeleteDoc(RcpGoodsImgChecksESIndex, "2_18_alex111")
	if err != nil {
		panic(err)
	}
	// DeleteDoc  ==> &{Index:rcp_goods_img_checks Type:_doc Id:2_18_alex111 Version:6 Result:deleted Shards:0xc00007e2c0 SeqNo:7 PrimaryTerm:1 Status:0 ForcedRefresh:false}
	fmt.Printf("DeleteDoc  ==> %+v \n\n", deleteDocRet)
}

func testCreateBulkDoc() {
	now := time.Now().Unix()
	// 批量创建
	createBulkDocRet, err := CreateBulkDoc(RcpGoodsImgChecksESIndex, []string{"h1", "h2", "h3"}, []interface{}{
		RcpGoodsImgChecksES{
			AppName:     2,
			GoodsId:     "h1_goods_id",
			SiteId:      17,
			CheckStatus: 1,
			CreatedAt:   int(now),
			UpdatedAt:   int(now),
		},
		RcpGoodsImgChecksES{
			AppName:     1,
			GoodsId:     "h2_goods_id",
			SiteId:      19,
			CheckStatus: 4,
			CreatedAt:   int(now),
			UpdatedAt:   int(now),
		},
		RcpGoodsImgChecksES{
			AppName:     3,
			GoodsId:     "h3_goods_id",
			SiteId:      19,
			CheckStatus: 2,
			CreatedAt:   int(now),
			UpdatedAt:   int(now),
		},
	})
	if err != nil {
		panic(err)
	}
	// CreateBulkDoc ==> &{Took:5 Errors:false Items:[map[index:0xc00019c200] map[index:0xc00019c280] map[index:0xc00019c300]]}
	fmt.Printf("CreateBulkDoc ==> %+v \n\n", createBulkDocRet)
}

func testUpdateBulkDoc() {
	// 批量更新
	updateBulkDocRet, err := UpdateBulkDoc(RcpGoodsImgChecksESIndex, []string{"h1", "h3"}, []interface{}{
		map[string]interface{}{
			"check_status": 2,
			"updated_at":   int(time.Now().Unix()),
		},
		map[string]interface{}{
			"site_id":    20,
			"updated_at": int(time.Now().Unix()),
		},
	})
	if err != nil {
		panic(err)
	}
	// UpdateBulkDoc ==> &{Took:6 Errors:false Items:[map[update:0xc0001e2080] map[update:0xc0001e2100]]}
	fmt.Printf("UpdateBulkDoc ==> %+v \n\n", updateBulkDocRet)
}

func testTermQuery() {
	// term
	query1 := elastic.NewTermQuery("goods_id", "h2_goods_id")
	querySearch(query1)

	// 开始打印参数 ====>
	// {
	//  "term": {
	//    "goods_id": "h2_goods_id"
	//  }
	// }
	// 打印参数结束 ====>
	// 查询到的结果总数为 1
	// 已经命中查询的数据为 ==> h2
	// {AppName:1 GoodsId:h2_goods_id SiteId:19 CheckStatus:4 CreatedAt:1660579860 UpdatedAt:1660579860}
}

func testTermsQuery() {
	// terms [where goods_id in ('h3_goods_id', 'h2_goods_id')]
	query2 := elastic.NewTermsQuery("goods_id", []interface{}{"h3_goods_id", "h2_goods_id"}...)
	querySearch(query2)

	// 开始打印参数 ====>
	// {
	//  "terms": {
	//    "goods_id": [
	//      "h3_goods_id",
	//      "h2_goods_id"
	//    ]
	//  }
	// }
	// 打印参数结束 ====>
	// 查询到的结果总数为 2
	// 已经命中查询的数据为 ==> h2
	// {AppName:1 GoodsId:h2_goods_id SiteId:19 CheckStatus:4 CreatedAt:1660579860 UpdatedAt:1660579860}
	//
	// 已经命中查询的数据为 ==> h3
	// {AppName:3 GoodsId:h3_goods_id SiteId:20 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
}

func testRangeQuery() {
	// 范围查找 [where updated_at >= 0 and updated_at <= 1659695758]
	// Gt（大于）、Lt（小于）、Gte（大于等于）、Lte（小于等于）
	query3 := elastic.NewRangeQuery("updated_at").Gte(0).Lte(1659695758)
	querySearch(query3)

	// 开始打印参数 ====>
	// {
	//  "range": {
	//    "updated_at": {
	//      "from": 0,
	//      "include_lower": true,
	//      "include_upper": true,
	//      "to": 1659695758
	//    }
	//  }
	// }
	// 打印参数结束 ====>
	// 查询到的结果总数为 0
}

func testMatchAllQuery() {
	// match_all
	query4 := elastic.NewMatchAllQuery()
	querySearch(query4)

	// 开始打印参数 ====>
	// {
	//  "match_all": {}
	// }
	// 打印参数结束 ====>
	// 查询到的结果总数为 4
	// 已经命中查询的数据为 ==> 2_19_alex111
	// {AppName:2 GoodsId:alex111 SiteId:18 CheckStatus:23 CreatedAt:1660579517 UpdatedAt:1660579517}
	//
	// 已经命中查询的数据为 ==> h2
	// {AppName:1 GoodsId:h2_goods_id SiteId:19 CheckStatus:4 CreatedAt:1660579860 UpdatedAt:1660579860}
	//
	// 已经命中查询的数据为 ==> h1
	// {AppName:2 GoodsId:h1_goods_id SiteId:17 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
	//
	// 已经命中查询的数据为 ==> h3
	// {AppName:3 GoodsId:h3_goods_id SiteId:20 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
}

func testMatchQuery() {
	// match
	query5 := elastic.NewMatchQuery("goods_id", "h2_goods_id")
	querySearch(query5)

	// 开始打印参数 ====>
	// {
	//  "match": {
	//    "goods_id": {
	//      "query": "h2_goods_id"
	//    }
	//  }
	// }
	// 打印参数结束 ====>
	// 查询到的结果总数为 1
	// 已经命中查询的数据为 ==> h2
	// {AppName:1 GoodsId:h2_goods_id SiteId:19 CheckStatus:4 CreatedAt:1660579860 UpdatedAt:1660579860}
}

func testMatchPhraseQuery() {
	// match_phrase
	query6 := elastic.NewMatchPhraseQuery("goods_id", "h2_goods_id")
	querySearch(query6)

	// 开始打印参数 ====>
	// {
	//  "match_phrase": {
	//    "goods_id": {
	//      "query": "h2_goods_id"
	//    }
	//  }
	// }
	// 打印参数结束 ====>
	// 查询到的结果总数为 1
	// 已经命中查询的数据为 ==> h2
	// {AppName:1 GoodsId:h2_goods_id SiteId:19 CheckStatus:4 CreatedAt:1660579860 UpdatedAt:1660579860}
}

func testMatchPhrasePrefixQuery() {
	// match_phrase_prefix
	// 这里因为类型不支持前缀匹配，可能会直接报错
	query7 := elastic.NewMatchPhrasePrefixQuery("goods_id", "h2_")
	querySearch(query7)
}

func testRegexpQuery() {
	// regexp
	// 搜索 goods_id 字段对应的值以 `h` 开头的所有文档
	query8 := elastic.NewRegexpQuery("goods_id", "h.*")
	querySearch(query8)

	// 开始打印参数 ====>
	// {
	//  "regexp": {
	//    "goods_id": {
	//      "value": "h.*"
	//    }
	//  }
	// }
	// 打印参数结束 ====>
	// 查询到的结果总数为 3
	// 已经命中查询的数据为 ==> h2
	// {AppName:1 GoodsId:h2_goods_id SiteId:19 CheckStatus:4 CreatedAt:1660579860 UpdatedAt:1660579860}
	//
	// 已经命中查询的数据为 ==> h1
	// {AppName:2 GoodsId:h1_goods_id SiteId:17 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
	//
	// 已经命中查询的数据为 ==> h3
	// {AppName:3 GoodsId:h3_goods_id SiteId:20 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
}

func testBoolQuery() {
	// 组合查询
	boolQuery := elastic.NewBoolQuery()
	// must
	boolQuery.Must(elastic.NewTermQuery("check_status", 2))
	// should
	boolQuery.Should(elastic.NewTermQuery("app_name", 20))
	// must_not
	boolQuery.MustNot(elastic.NewTermQuery("site_id", 18))
	// filter
	boolQuery.Filter(elastic.NewRangeQuery("updated_at").Gte(0).Lte(1660579923))
	querySearch(boolQuery)

	// 开始打印参数 ====>
	// {
	//  "bool": {
	//    "filter": {
	//      "range": {
	//        "updated_at": {
	//          "from": 0,
	//          "include_lower": true,
	//          "include_upper": true,
	//          "to": 1660579923
	//        }
	//      }
	//    },
	//    "must": {
	//      "term": {
	//        "check_status": 2
	//      }
	//    },
	//    "must_not": {
	//      "term": {
	//        "site_id": 18
	//      }
	//    },
	//    "should": {
	//      "term": {
	//        "app_name": 20
	//      }
	//    }
	//  }
	// }
	// 打印参数结束 ====>
	// 查询到的结果总数为 2
	// 已经命中查询的数据为 ==> h1
	// {AppName:2 GoodsId:h1_goods_id SiteId:17 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
	//
	// 已经命中查询的数据为 ==> h3
	// {AppName:3 GoodsId:h3_goods_id SiteId:20 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
}

func testPageSort() {
	// 分页查询，并排序
	// from 为起始偏移量（offset）默认为 0，size 为每页显示数（limit）默认为 10
	// from 等于当前页码数减去一的商然后乘以每页显示数
	// Sort() 第二个参数，true 为升序、false 为降序
	pageRet, err := ESClient.Search().Index(RcpGoodsImgChecksESIndex).From(0).Size(20).Sort("updated_at", false).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range pageRet.Hits.Hits {
		var tmp RcpGoodsImgChecksES
		json.Unmarshal(v.Source, &tmp)
		fmt.Printf("分页查询，已经命中查询的数据为 ==> %+v \n %+v \n\n", v.Id, tmp)
	}

	// 分页查询，已经命中查询的数据为 ==> h1
	// {AppName:2 GoodsId:h1_goods_id SiteId:17 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
	//
	// 分页查询，已经命中查询的数据为 ==> h3
	// {AppName:3 GoodsId:h3_goods_id SiteId:20 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
	//
	// 分页查询，已经命中查询的数据为 ==> h2
	// {AppName:1 GoodsId:h2_goods_id SiteId:19 CheckStatus:4 CreatedAt:1660579860 UpdatedAt:1660579860}
	//
	// 分页查询，已经命中查询的数据为 ==> 2_19_alex111
	// {AppName:2 GoodsId:alex111 SiteId:18 CheckStatus:23 CreatedAt:1660579517 UpdatedAt:1660579517}
}

func testMultiFieldSort() {
	// 多字段排序
	sortsBuilders := []elastic.Sorter{
		elastic.NewFieldSort("check_status").Asc(), // 升序
		elastic.NewFieldSort("created_at").Desc(),  // 降序
	}
	sortRet, err := ESClient.Search().Index(RcpGoodsImgChecksESIndex).SortBy(sortsBuilders...).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range sortRet.Hits.Hits {
		var tmp RcpGoodsImgChecksES
		json.Unmarshal(v.Source, &tmp)
		fmt.Printf("多字段排序，已经命中查询的数据为 ==> %+v \n %+v \n\n", v.Id, tmp)
	}

	// 多字段排序，已经命中查询的数据为 ==> h1
	// {AppName:2 GoodsId:h1_goods_id SiteId:17 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
	//
	// 多字段排序，已经命中查询的数据为 ==> h3
	// {AppName:3 GoodsId:h3_goods_id SiteId:20 CheckStatus:2 CreatedAt:1660579860 UpdatedAt:1660579923}
	//
	// 多字段排序，已经命中查询的数据为 ==> h2
	// {AppName:1 GoodsId:h2_goods_id SiteId:19 CheckStatus:4 CreatedAt:1660579860 UpdatedAt:1660579860}
	//
	// 多字段排序，已经命中查询的数据为 ==> 2_19_alex111
	// {AppName:2 GoodsId:alex111 SiteId:18 CheckStatus:23 CreatedAt:1660579517 UpdatedAt:1660579517}
}

func testFetchSource() {
	// 返回指定字段
	includeFields := elastic.NewFetchSourceContext(true).Include([]string{"app_name", "goods_id"}...)
	includeRet, err := ESClient.Search().Index(RcpGoodsImgChecksESIndex).FetchSourceContext(includeFields).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range includeRet.Hits.Hits {
		var tmp RcpGoodsImgChecksES
		json.Unmarshal(v.Source, &tmp)
		fmt.Printf("返回指定字段，已经命中查询的数据为 ==> %+v \n %+v \n\n", v.Id, tmp)
	}

	// 返回指定字段，已经命中查询的数据为 ==> 2_19_alex111
	// {AppName:2 GoodsId:alex111 SiteId:0 CheckStatus:0 CreatedAt:0 UpdatedAt:0}
	//
	// 返回指定字段，已经命中查询的数据为 ==> h2
	// {AppName:1 GoodsId:h2_goods_id SiteId:0 CheckStatus:0 CreatedAt:0 UpdatedAt:0}
	//
	// 返回指定字段，已经命中查询的数据为 ==> h1
	// {AppName:2 GoodsId:h1_goods_id SiteId:0 CheckStatus:0 CreatedAt:0 UpdatedAt:0}
	//
	// 返回指定字段，已经命中查询的数据为 ==> h3
	// {AppName:3 GoodsId:h3_goods_id SiteId:0 CheckStatus:0 CreatedAt:0 UpdatedAt:0}
}

func testTotal() {
	// 查询总命中计数
	total, err := ESClient.Count().Index(RcpGoodsImgChecksESIndex).Do(context.Background())
	if err != nil {
		panic(err)
	}

	// 查询总命中计数，已经命中查询的数据为 ==> 2
	fmt.Printf("查询总命中计数，已经命中查询的数据为 ==> %+v \n", total)
}

func testDeleteBulkDoc() {
	// 通过文档 id 批量删除
	deleteBulkDocRet, err := DeleteBulkDoc(RcpGoodsImgChecksESIndex, []string{"h2", "h3_goods_id"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("DeleteBulkDoc ==> %+v \n\n", deleteBulkDocRet)

	// DeleteBulkDoc ==> &{Took:36 Errors:false Items:[map[delete:0xc0000ea080] map[delete:0xc0000ea100]]}
}

func testDeleteByQuery() {
	// 按照条件删除
	deleteDocByQuery, err := ESClient.DeleteByQuery(RcpGoodsImgChecksESIndex).
		Query(elastic.NewRangeQuery("updated_at").Gte(0).Lte(1660579923)).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("deleteDocByQuery ==> %+v \n\n", deleteDocByQuery)

	// deleteDocByQuery ==> &{Header:map[] Took:36 SliceId:<nil> TimedOut:false Total:3 Updated:0 Created:0 Deleted:3 Batches:1 VersionConflicts:0 Noops:0 Retries:{Bulk:0 Search:0} Throttled: ThrottledMillis:0 RequestsPerSecond:-1 Canceled: ThrottledUntil: ThrottledUntilMillis:0 Failures:[]}
}
