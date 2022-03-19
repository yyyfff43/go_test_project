/*
* @File : book
* @Describe : 书籍结构体
* @Author: Jerry Yang
* @Date : 2022/3/19 10:00
* @Software: GoLand
 */

package dao

//定义结构体(xorm支持双向映射)
type Book struct {
	Id         int    `xorm:"id pk autoincr"` //指定主键并自增，执行insert后传引用的结构体对象会获得id作为last_insert_id,否则id是0
	BookName   string `xorm:"book_name"`
	Desc       string `xorm:"desc"`
	Pic        string `xorm:"pic"`
	Category   int    `xorm:"category"`
	UpdateTime int64  `xorm:"update_time updated"` //修改后自动更新时间，Unix时间戳格式
	CreateTime int64  `xorm:"create_time created"` //创建时间，如果不指定则自己按当前时间创建，Unix时间戳格式
	//Version string `xorm:"version"` //乐观锁
}
