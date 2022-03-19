/*
* @File : mysql_by_xorm_test
* @Describe :
* @Author: Jerry Yang
* @Date : 2022/3/14 23:55
* @Software: GoLand
 */

package mysql_learn

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go_test_project/src/dao"
	"testing"
	"time"
)

//测试初始化
func TestInitMySqlXorm(t *testing.T) {

	engine := InitMySqlXorm(dsn)
	//	assert.Nil(t,err)
	fmt.Println(engine)
}

//测试按条件取一条记录
func TestSimpleGet(t *testing.T) {
	res, err := SimpleGet()
	assert.Nil(t, err)
	fmt.Println(res)
}

//测试获取数据库和表的信息
func TestShowDbInfo(t *testing.T) {
	ShowDbInfo()
}

//测试插入一条数据
func TestInsertData(t *testing.T) {

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(errors.Wrap(err, "获取正确时区时出错"))
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, _ := time.ParseInLocation("2006/01/02 15:04:05", "2022/01/05 11:25:20", loc)

	var user = &dao.User{
		Name:       "谭勇林",
		Age:        58,
		Gender:     1,
		SubmitTime: timeObj.Format("2006/01/02 15:04:05"),
	}
	res := InsertData(user)
	//	assert.Nil(t, err)
	fmt.Println(res)
}

//测试插入多条数据
func TestInsertDatas(t *testing.T) {
	now := time.Now()
	// 格式化的模板为 2006-01-02 15:04:05

	var users = make([]*dao.User, 0, 3)

	var user = &dao.User{
		Name:       "王圆",
		Age:        20,
		Gender:     1,
		SubmitTime: now.Format("2006/01/02 15:04:05"),
	}

	var user2 = &dao.User{
		Name:       "赵今麦",
		Age:        19,
		Gender:     1,
		SubmitTime: now.Format("2006/01/02 15:04:05"),
	}

	users = append(users, user, user2)
	res := InsertDatas(users)

	//	assert.Nil(t, err)
	fmt.Println(res)
}

//测试插入书籍并返回新的id
func TestInsertBookAndGetId(t *testing.T) {
	var book = &dao.Book{
		BookName: "鲁菜大全",
		Desc:     "山东菜的由来",
		Pic:      "https://bkimg.cdn.bcebos.com/pic/cc11728b4710b912ea015e75cafdfc039345227c?x-bce-process=image/resize,m_lfit,w_440,limit_1/format,f_auto",
	}
	res := InsertBookAndGetId(book)
	fmt.Println(res)
}

//测试查询
func TestDoQuery(t *testing.T) {
	DoQuery()
}

//测试sql语句批量查询
func TestDoQueryBySql(t *testing.T) {
	sqlStr := "SELECT id, book_name, book.desc,pic,category,create_time,update_time FROM `book` where id > ? order by update_time desc LIMIT ?, ?"
	DoQueryBySql(sqlStr, 1, 0, 20)
}
