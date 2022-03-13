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
	"time"

	"fmt"
)

var dsn string = "root:123456@tcp(127.0.0.1:3306)/simple_test?charset=utf8mb4&parseTime=True"

type User struct {
	Id         int    `db:"id"`
	Age        int    `db:"age"`
	Name       string `db:"name"`
	Gender     int    `db:"gender"`
	SubmitTime string `db:"submitTime"`
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
	//defer db.Close()
	//如果需要在一个完整的初始换到关闭的过程中，defer db.close()必须在err判断为空后边调用，否则db对象有
	//可能报错返回一个db为nil的对象，这时db调用close会报panic
	//initDb返回出去db对象要在调用函数或包用完时关闭，如果怕忘记获得可使用defer提前定义

	// 尝试与数据库建立连接（校验dsn是否正确）
	errPing := db.Ping()
	if err != nil {
		fmt.Println(errPing.Error())
	}

	//设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，
	//会将最大闲置连接数减小到匹配最大开启连接数的限制。 如果n<=0，不会限制最大开启连接数，默认为0（无限制）。
	db.SetMaxOpenConns(100)

	//SetMaxIdleConns设置连接池中的最大闲置连接数。 如果n大于最大开启连接数，
	//则新的最大闲置连接数会减小到匹配最大开启连接数的限制。 如果n<=0，不会保留闲置连接。
	db.SetMaxIdleConns(20)

	return db, err
}

//
//  QueryRowDemo
//  @Description: 查询单条数据示例
//
func QueryRowDemo() {
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
//  queryMultiRowDemo
//  @Description: 查询多条数据示例
//
func QueryMultiRowDemo() {
	db, _ := initDb(dsn)
	defer db.Close() //数据库必须要释放连接

	sqlStr := "select id, name, age from user where id > ? order by id desc"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
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
//  InsertRowDemo
//  @Description: 插入数据
//
func InsertRowDemo() {
	db, _ := initDb(dsn)
	defer db.Close() //数据库必须要释放连接

	sqlStr := "insert into user(name, age,gender,submit_time) values (?,?,?,?)"
	ret, err := db.Exec(sqlStr, "张博智", 31, 2, time.Now())
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func UpdateRowDemo() {
	db, _ := initDb(dsn)
	defer db.Close() //数据库必须要释放连接

	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 19, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

//
//  DeleteRowDemo
//  @Description: 删除数据
//
func DeleteRowDemo() {
	db, _ := initDb(dsn)
	defer db.Close() //数据库必须要释放连接

	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 7)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
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
