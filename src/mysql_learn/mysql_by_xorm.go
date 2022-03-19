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
	"github.com/go-xorm/builder"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"go_test_project/src/dao"
	"time"
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

	//如果某个文件不存在，那么使用os.Lstat就一定会返回error，只要判断error是否代表文件不存在即可
	//_, errF := os.Lstat("sql.log")
	//var f = new(os.File)
	//var errOs error
	//if !os.IsNotExist(errF) {
	//	f, errOs = os.Create("sql.log")
	//	if errOs != nil {
	//		println(errOs.Error())
	//		return nil
	//	}
	//} else {
	//f, errOs = os.Open("sql.log")
	//if errOs != nil {
	//	println(errOs.Error())
	//	return nil
	//}
	//}

	//	engine.SetLogger(xorm.NewSimpleLogger(f)) // TODO:待解决**************每次执行的日志都会被覆盖

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
func SimpleGet() (*dao.User, error) {
	engine := InitMySqlXorm(dsn)
	user := new(dao.User)
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

	//xorm支持获取表结构信息，通过调用engine.DBMetas()可以获取到数据库中所有的表，字段，索引的信息。
	dbInfos, err := engine.DBMetas()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range dbInfos {
		fmt.Println("涉及到的表名：" + v.Name + "\n")
	}

	//根据传入的结构体指针及其对应的Tag，提取出模型对应的表结构信息。这里不是数据库当前的表结构信息，
	//而是我们通过struct建模时希望数据库的表的结构信息
	var user = new(dao.User) //此处可以直接使用传统mysql定义的结构体，因为数据库声明tag一致
	isUserExist, errEx := engine.IsTableExist(user)
	if errEx != nil {
		fmt.Println(errors.Wrap(errEx, "判断user表IsTableExist出错"))
		return
	}

	if isUserExist {
		tablesInfo := engine.TableInfo(user)
		fmt.Println("表名：" + tablesInfo.Name)
		for _, v := range tablesInfo.Columns() {
			fmt.Println("字段名： " + v.Name)
		}
	} else {
		fmt.Println(errors.New("user表不存在"))
	}
}

//
//  insertData
//  @Description: 插入数据
//  @param user
//  @return bool
//
func InsertData(user *dao.User) int64 {
	if user != nil {
		engine := InitMySqlXorm(dsn)
		//默认使用格林尼治时间，需改变xorm的时区，否则差8小时
		engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
		insertRowSum, err := engine.Insert(user) //如果只插入一个对象也可使用InsertOne
		if err != nil {
			fmt.Println(errors.Wrap(err, "插入数据库表user失败"))
			return 0
		}
		if insertRowSum > 0 {
			return insertRowSum
		} else {
			return 0
		}
	} else {
		return 0
	}

}

//
//  InsertDatas
//  @Description: 插入多条记录
//  @param users
//  @return int64
//
func InsertDatas(users []*dao.User) int64 {
	if len(users) > 0 {
		engine := InitMySqlXorm(dsn)
		//默认使用格林尼治时间，需改变xorm的时区，否则差8小时
		engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
		insertRowSum, err := engine.Insert(users) //插入多条数据直接传入这个结构体的切片
		if err != nil {
			fmt.Println(errors.Wrap(err, "插入数据库表user失败"))
			return 0
		}
		if insertRowSum > 0 {
			return insertRowSum
		} else {
			return 0
		}
	} else {
		return 0
	}
	return 0
}

//
//  InsertBookAndGetId
//  @Description: 插入一本新书并获取这个书的last_insert_id
//  @param book
//  @return int64
//
func InsertBookAndGetId(book *dao.Book) int {
	if book != nil {
		engine := InitMySqlXorm(dsn)
		insertRowSum, err := engine.Insert(book) //插入多条数据直接传入这个结构体的切片
		if err != nil {
			fmt.Println(errors.Wrap(err, "插入数据库表book失败"))
			return 0
		}
		if insertRowSum > 0 {
			return book.Id
		} else {
			return 0
		}
	}
	return 0
}

//
//  DoQuery
//  @Description: 基本查询练习
//
func DoQuery() {
	engine := InitMySqlXorm(dsn)
	//默认使用格林尼治时间，需改变xorm的时区，否则差8小时
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	var user = new(dao.User)
	//给表定义一个别名
	res, err := engine.Alias("u").Where("u.name = ?", "张曼玉").Get(user)
	if err != nil {
		fmt.Println(errors.Wrap(err, "别名查询练习时出错"))
	}
	if res {
		fmt.Println(user)
	}

	//使用and并排序，排序可以多个条件，越靠前优先级越高
	var user2 = new(dao.User)
	res2, err2 := engine.Alias("u").Where("u.name = ?", "刘得滑").And("u.age = ?", "59").
		Desc("submit_time").Desc("id").Get(user2)
	if err2 != nil {
		fmt.Println(errors.Wrap(err2, "And查询练习时出错"))
	}
	if res2 {
		fmt.Println(user2)
	}

	//按主键查询一条数据
	var book = new(dao.Book)
	resBool, err3 := engine.ID(1).Get(book)
	if err2 != nil {
		fmt.Println(errors.Wrap(err3, "按id查出一本书"))
	}
	if resBool {
		fmt.Println(book)
	}

}

//
//  DoQueryBySql
//  @Description: sql语句直接通过xorm访问数据库
//  @param sql
//
func DoQueryBySql(sql string, id, startPage, pageSize int) {
	engine := InitMySqlXorm(dsn)
	var beans = make([]*dao.Book, 0)
	err := engine.SQL(sql, id, startPage, pageSize).Find(&beans)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Sql批量查书"))
	} else {
		fmt.Println(beans)
	}
}

//
//  DoWhereQuery
//  @Description: where条件查询
//
func DoWhereQuery() {
	engine := InitMySqlXorm(dsn)
	var beans = make([]*dao.Book, 0)
	err := engine.Where("book_name = ? AND category = ?", "鲁菜大全", 0).Find(&beans)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Sql,Where条件批量查书"))
	} else {
		fmt.Println(beans)
	}

	var beanS2 = make([]*dao.Book, 0)
	where := builder.Eq{"book_name": "图解Http", "category": 1}
	whereGte := builder.Gte{"update_time": "1647654000"}
	//多个where条件组合，使用链式访问
	err2 := engine.Where(where).Where(whereGte).Limit(0, 100).Find(&beanS2)
	if err2 != nil {
		fmt.Println(errors.Wrap(err, "Sql,Where使用builder构建条件批量查书"))
	} else {
		fmt.Println(beans)
	}
}
