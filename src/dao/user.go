/*
* @File : user
* @Describe : 用户类型的结构体
* @Author: Jerry Yang
* @Date : 2022/3/19 9:58
* @Software: GoLand
 */

package dao

//注意此处的标签tag中虽然是db开头，但是后边声明的值如果和数据库相同，xorm也是可以读取的，不必把db：换为xorm:
type User struct {
	Id         int    `db:"id"`
	Age        int    `db:"age"`
	Name       string `db:"name"`
	Gender     int    `db:"gender"`
	SubmitTime string `db:"submitTime"`
}
