package err_learn

import "testing"
//单元测试
func TestTryRecover(t *testing.T) {
	TryRecover()
	//if actual != expected {
	//	t.Errorf("TryRecover返回的是%s; 期望的是 %s", actual, expected)
	//}
}

//性能测试
func BenchmarkTryRecover(b *testing.B) {
	for i := 0; i < b.N; i++ { //b.N表示Test自动会选择执行测试的次数
		TryRecover()
	}
}
