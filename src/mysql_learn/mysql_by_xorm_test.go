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
	"github.com/stretchr/testify/assert"
	"testing"
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
