/*
* @File : Jerry Yang
* @Describe :
* @Author: 248037973@qq.com
* @Date : 2022/3/6 22:21
* @Software: GoLand
 */
package jwt_learn

import (
	"fmt"
	"testing"
)

//测试生成令牌
func TestGenToken(t *testing.T) {
	var name = "yangfan"

	tokenStr, err := GenToken(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("生成的令牌为：%v\n", tokenStr)
	}
}

//测试解析令牌
func TestParseToken(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InlhbmdmYW4iLCJleHAiOjE2NDY1Nzc0NjEsImlzcyI6Im15LXByb2plY3QifQ.opsq6UmA6t9ckiQJpuW81pL49wZf-RBCcgzCplfk0C0"
	mc, err := ParseToken(token)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("令牌验证通过，获得用户名为：%s\n", mc.Username)
	}

}
