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

func V8Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		fmt.Println(err.Error())
		return
	}

	err := rdb.Set(ctx, "key1", "value1", 300).Err()
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
