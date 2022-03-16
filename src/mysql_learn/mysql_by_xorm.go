/*
* @File : mysql_by_xorm
* @Describe : xorm关系映射包访问mysql，文档见http://xorm.topgoer.com/
* 拉取使用 github.com/go-xorm/cmd/xorm  带cmd字符串
* @Author: Jerry Yang
* @Date : 2022/3/14 23:14
* @Software: GoLand
 */

package mysql_learn

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
)

//
//  InitMySqlXorm
//  @Description: 初始化xorm支持的mysql,一般情况下如果只操作一个数据库，只需要创建一个engine即可。engine是GoRoutine安全的。
//  @param dsn
//  @return *xorm.Engine
//
func InitMySqlXorm(dsn string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err.Error())
		return nil
	}

	engine.ShowSQL(true)                     //则会在控制台打印出生成的SQL语句；
	engine.Logger().SetLevel(core.LOG_DEBUG) //则会在控制台打印调试及以上的信息；

	//如果需要保存日志文件：
	f, errOs := os.Create("sql.log")
	if errOs != nil {
		println(errOs.Error())
		return nil
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))

	engine.SetMaxOpenConns(100) //设置最大打开连接数
	engine.SetMaxIdleConns(20)  //设置连接池的空闲数大小

	//创建完成engine之后，并没有立即连接数据库，此时可以通过engine.Ping()来进行数据库的连接测试是否可以连接到数据库
	errPing := engine.Ping()
	if errPing != nil {
		println(errPing.Error())
		return nil
	}

	return engine
}

//
//  SimpleGet
//  @Description: 简单的get查询
//  @return *User
//  @return error
//
func SimpleGet() (*User, error) {
	engine := InitMySqlXorm(dsn)
	user := new(User)
	has, err := engine.Where("name=?", "张曼玉").Get(user) //方法体如果传参是结构体对象则必须是指针类型
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	if has {
		return user, err
	} else {
		return nil, err
	}
}

//
//  ShowDbInfo
//  @Description: 现实数据库信息和指定表名称的表信息
//
func ShowDbInfo() {
	engine := InitMySqlXorm(dsn)
	dbInfos, err := engine.DBMetas()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range dbInfos {
		fmt.Println("表名：" + v.Name + "\n")
	}

	tablesInfo := engine.TableInfo("user")
	fmt.Println("表名：" + tablesInfo.Name)
	fmt.Println("表注释：" + tablesInfo.Comment)
}
