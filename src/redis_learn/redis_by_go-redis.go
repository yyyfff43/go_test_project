/*
* @File : redis_by_go-redis
* @Describe : go-redis包支持的go语言redis存取访问
* @Author: Jerry Yang
* @Date : 2022/3/13 0:09
* @Software: GoLand
 */

package redis_learn

//最新版本的go-redis库的相关命令都需要传递context.Context参数
import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

var (
	rdb *redis.Client
)

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

//初始化小测试
func V8Example() {
	//v8版本带ctx，如果是正式版，不带ctx
	ctx := context.Background()
	if err := initClient(); err != nil {
		fmt.Println(errors.Wrap(err, "redis初始化"))
		return
	}
	defer rdb.Close()

	err := rdb.Set(ctx, "key1", "value1", 600*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// Output: key value
	// key2 does not exist
}

//
//  stringOperation
//  @Description: String 操作
//  @param client
//
func StringOperation() {

	ctx := context.Background()
	if err := initClient(); err != nil {
		fmt.Println(errors.Wrap(err, "redis初始化"))
		return
	}

	// 第三个参数是过期时间, 如果是 0, 则表示没有过期时间.
	err := rdb.SetEX(ctx, "markName", "xys", 600*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "markName").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("markName", val)

	// 测试一个自增自减的例子
	err = rdb.Set(ctx, "age", "20", 600*time.Second).Err()
	if err != nil {
		panic(err)
	}

	rdb.Incr(ctx, "age") // 自增
	rdb.Incr(ctx, "age") // 自增
	rdb.Decr(ctx, "age") // 自减

	val, err = rdb.Get(ctx, "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("age", val) // age 的值为21

}

//
//  ListOperation
//  @Description: List操作
//  @param client
//
func ListOperation() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		fmt.Println(errors.Wrap(err, "redis初始化"))
		return
	}

	rdb.RPush(ctx, "fruit", "apple")               // 在名称为 fruit 的list尾添加一个值为value的元素
	rdb.LPush(ctx, "fruit", "banana")              // 在名称为 fruit 的list头添加一个值为value的 元素
	rdb.LPush(ctx, "fruit", "kiwi")                // 在名称为 fruit 的list头添加一个值为value的 元素
	length, err := rdb.LLen(ctx, "fruit").Result() // 返回名称为 fruit 的list的长度
	if err != nil {
		panic(err)
	}
	rdb.Expire(ctx, "fruit", 600*time.Second)
	fmt.Println("length: ", length) // 长度为2

	//value, err := rdb.LPop(ctx, "fruit").Result() //返回并删除名称为 fruit 的list中的首元素
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("fruit: ", value)
	//
	//value, err = rdb.RPop(ctx, "fruit").Result() // 返回并删除名称为 fruit 的list中的尾元素
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("fruit: ", value)

	//遍历一个列表
	res := rdb.LRange(ctx, "fruit", 0, -1)
	resstr, errList := res.Result()
	if errList != nil {
		fmt.Println(errList.Error())
	} else {
		fmt.Println(resstr)
	}

}
