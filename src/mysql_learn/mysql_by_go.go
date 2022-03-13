/*
* @File : mysql_by_go
* @Describe : go调用mysql练习,使用database/sql标准包
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

//
//  initDb
//  @Description: 初始化mysql数据库
//  @param dsn Data Source 数据源信息，包括主机ip账号密码端口等
//  @return *sql.DB
//  @return error
//
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

//
//  QueryRowDemo
//  @Description: 查询单条数据示例
//
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

//
//  prepareQueryDemo
//  @Description: 预处理查询示例,为什么使用预处理：1优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，
// 一次编译多次执行，节省后续编译的成本。
//  2避免SQL注入问题。
//
func PrepareQueryDemo() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/simple_test?charset=utf8mb4&parseTime=True"
	db, _ := initDb(dsn)
	defer db.Close() //数据库必须要释放连接
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
}

//
//  TransactionDemo
//  @Description: 事务操作，该事物操作能够确保两次更新操作要么同时成功要么同时失败，不会存在中间状态。
//
func TransactionDemo() {
	//事务的原子性：一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。
	//事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。

	dsn := "root:123456@tcp(127.0.0.1:3306)/simple_test?charset=utf8mb4&parseTime=True"
	db, _ := initDb(dsn)
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}
