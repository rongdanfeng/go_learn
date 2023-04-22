package CH01

import "testing"

func TestIota(t *testing.T) {
	//iota，特殊常量，可以认为是一个可以被编译器修改的常量
	const (
		EER1 = iota
		EER2
		EER3
		EER4
	)
	t.Log(EER1, EER2, EER3, EER4)
	// 0,1,2,3

	/*如果中断了iota那么必须显式的恢复，后续会自动递增自增类型默认是int类型
	iota能简化const类型的定义*/

}
