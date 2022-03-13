/*
* @File : mysql_by_go_test
* @Describe :go调用mysql练习单元测试
* @Author: Jerry Yang
* @Date : 2022/3/3 23:45
* @Software: GoLand
 */

package mysql_learn

import "testing"

//测试单条查询
func TestQueryRowDemo(t *testing.T) {
	QueryRowDemo()
}

//测试多条查询
func TestQueryMultiRowDemo(t *testing.T) {
	QueryMultiRowDemo()
}

//测试插入操作
func TestInsertRowDemo(t *testing.T) {
	InsertRowDemo()
}

//测试删除数据
func TestUpdateRowDemo(t *testing.T) {
	UpdateRowDemo()
}

func TestDeleteRowDemo(t *testing.T) {
	DeleteRowDemo()
}

//测试预处理查询
func TestPrepareQueryDemo(t *testing.T) {
	PrepareQueryDemo()
}

//测试事务
func TestTransactionDemo(t *testing.T) {
	TransactionDemo()
}
