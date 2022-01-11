
package func_include_ceshi

import (
	"fmt"
	"testing"
)

//一般单元测试方法
func TestAddSomeThing(t *testing.T) {
	var (
		a,b = 1,2
		expected = 3
	)

	actual := AddSomeThing(a,b)
	if actual != expected {
		t.Errorf("AddSomeThing方法(%d，%d) = %d; 期望的是 %d", a,b, actual, expected)
	}
}

//表格驱动式测试方法
func TestStr2Int(t *testing.T) {
	tests := []struct{
		a int
		s string
	}{
		{3,"3"},
		{88,"88"},
		{123,"123"},
		{1234423245,"1234423245"},
		{34584534579345893,"34584534579345893"},
	}
	for _, tt := range tests {
		if actual := Str2Int(tt.a); actual != tt.s {
			t.Errorf("Str2Int(%d); "+
				"程序结果为 %s; 期望结果是 %s",
				tt.a, actual, tt.s)
		}
	}
}

//go test -coverprofile=c.out  生成本测试用例的c.out文件用来查看单元测试的覆盖情况
//go tool cover -html=c.out 用来以网页可视化查看单元测试覆盖情况


//Benchmark性能测试，如果通过命令行在当前目录则执行go test -bench . （前提是所有这个test单元的测试都pass了）
func BenchmarkAddSomeThing(b *testing.B) {
	var t1,t2 = 1,2
	ans := 3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {//b.N表示Test自动会选择执行测试的次数
		actual := AddSomeThing(t1,t2)
		if actual != ans {
			b.Errorf("got %d for input %d and %d; "+
				"expected %d",
				actual, t1,t2, ans)
		}
	}
}

func ExampleAddSomeThing() {
    res := AddSomeThing(1,2)
    fmt.Println(res)
	//Output:
	//5
}

//go test -bench . -cpuprofile=cpu.out
//go tool pprof cpu.out (图里框框越大，花的时间越久)
//执行web命令查看svg图，quit退出，之前需要安装graphviz插件，地址：http://www.graphviz.org/download/


//go doc 通过命令行查看文档介绍和注释
//godoc命令比如执行godoc -http://localhost:6060 可用网页查看注释

//使用Goanno插件在函数上方生成注释模板

//Example来执行函数的示例,如果不和期望结果相等，这个用例不会通过