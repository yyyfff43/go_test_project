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

//
//  SetOperation
//  @Description: set 操作去重集合
//
func SetOperation() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		fmt.Println(errors.Wrap(err, "redis初始化"))
		return
	}

	rdb.SAdd(ctx, "blacklist", "Obama")     // 向 blacklist 中添加元素
	rdb.SAdd(ctx, "blacklist", "Hillary")   // 再次添加
	rdb.SAdd(ctx, "blacklist", "the Elder") // 添加新元素

	rdb.SAdd(ctx, "whitelist", "the Elder") // 向 whitelist 添加元素

	// 判断元素是否在集合中
	isMember, err := rdb.SIsMember(ctx, "blacklist", "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is Bush in blacklist: ", isMember)

	// 求交集, 即既在黑名单中, 又在白名单中的元素
	names, err := rdb.SInter(ctx, "blacklist", "whitelist").Result()
	if err != nil {
		panic(err)
	}
	// 获取到的元素是 "the Elder"
	fmt.Println("Inter result: ", names)

	// 获取指定集合的所有元素
	all, err := rdb.SMembers(ctx, "blacklist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All member: ", all)
	//设置过期时间
	rdb.Expire(ctx, "blacklist", 600*time.Second)
	rdb.Expire(ctx, "whitelist", 600*time.Second)
}

//
//  hashOperation
//  @Description:  hash 操作
//
func HashOperation() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		fmt.Println(errors.Wrap(err, "redis初始化"))
		return
	}

	rdb.HSet(ctx, "user_xys", "name", "xys") // 向名称为 user_xys 的 hash 中添加元素 name
	rdb.HSet(ctx, "user_xys", "age", "18")   // 向名称为 user_xys 的 hash 中添加元素 age

	// 批量地向名称为 user_test 的 hash 中添加元素 name 和 age
	rdb.HMSet(ctx, "user_test", map[string]string{"name": "test", "age": "20"})
	// 批量获取名为 user_test 的 hash 中的指定字段的值.
	fields, err := rdb.HMGet(ctx, "user_test", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fields in user_test: ", fields)

	// 获取名为 user_xys 的 hash 中的字段个数
	length, err := rdb.HLen(ctx, "user_xys").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("field count in user_xys: ", length) // 字段个数为2

	// 删除名为 user_test 的 age 字段
	rdb.HDel(ctx, "user_test", "age")
	age, err := rdb.HGet(ctx, "user_test", "age").Result()
	if err != nil {
		fmt.Printf("Get user_test age error: %v\n", err)
	} else {
		fmt.Println("user_test age is: ", age)
	}

	//设置过期时间
	rdb.Expire(ctx, "user_xys", 600*time.Second)
	rdb.Expire(ctx, "user_test", 600*time.Second)
}
