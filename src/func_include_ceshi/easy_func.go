package func_include_ceshi

import "strconv"

//一个简单的相加函数
func AddSomeThing(a,b int) int  {
	 return a+b
}

//一个简单的int转string函数
func Str2Int(t int) string{
	return strconv.Itoa(t)
}
