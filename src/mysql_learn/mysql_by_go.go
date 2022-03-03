/*
* @File : mysql_by_go
* @Describe : go调用mysql练习
* @Author: Jerry Yang
* @Date : 2022/3/3 23:22
* @Software: GoLand
 */

package mysql_learn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

type User struct {
	Id         int
	Age        int
	Name       string
	Gender     int
	SubmitTime string
}

func initDb(dsn string) (*sql.DB, error) {
	// DSN:Data Source Name as param
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	errPing := db.Ping()
	if err != nil {
		fmt.Println(errPing.Error())
	}
	return db, err
}

// 查询单条数据示例
func QueryRowDemo() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/simple_test?charset=utf8mb4&parseTime=True"
	db, _ := initDb(dsn)
	defer db.Close() //数据库必须要释放连接
	sqlStr := "select id, name, age from user where id=?"
	var u User
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}
