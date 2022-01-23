//家里电脑使用的es版本是7.16.2，所以olivere的go包用v7的
package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"reflect"
)

type people struct {
	Name string
	Country string
	Age int
	date string
}

//
//  BasicOperator
//  @Description: ES基本操作，包括index创建，文档添加，文档查找，文档删除，文档修改
//
func BasicOperator() {

	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 创建client
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),

		// 设置错误日志输出(会在控制台输出)
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出(会在控制台输出)
//		elastic.SetInfoLog(log.New(os.Stdout, "info:", log.LstdFlags)),
		// 设置trace日志输出(会在控制台输出，包括查询的dsl语句，跟踪过程等)
//		elastic.SetTraceLog(log.New(os.Stdout, "trace", log.LstdFlags)),
	)
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	
	// 创建match查询条件
	matchQuery := elastic.NewMatchQuery("name", "爸爸")
//    mapStr,_ := matchQuery.Source()//可以以map的形式访问查询条件
//    mapJson,_ := json.Marshal(mapStr)//转json格式
//    fmt.Println(mapStr)
	searchResult, err := client.Search().
		Index("people").        // 设置索引名
		Query(matchQuery).     // 设置查询条件
//		FetchSourceContext(elastic.NewFetchSourceContext(true).Include("id", "name")).//只返回的指定字段
		Sort("date", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(0).               // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(10).              // 设置分页参数 - 每页大小
		Do(ctx)                // 执行请求
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 people
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成Article对象
			if t, ok := item.(people); ok {
				fmt.Println(t.Name)
			}
		}
	}


}
