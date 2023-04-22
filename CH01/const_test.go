package CH01

import "testing"

func TestConst(t *testing.T) {
	//常量，定义的时候就指定的值，不能修改  常量尽量全部大写
	const PI = 3.1141516  //显示定义类型
	const PI2 = 3.1141516 //隐式定义

	const (
		UNKNOWN = 1
		FEMALE  = 2
		MALE    = 3
	)
	const (
		X int = 16
		Y
		S = "ABC"
		Z
	)

	t.Log(X, Y, S, Z)

	/*常量类型只可以定义bool、数值(整数、浮点数和复数)和字符串
	不曾使用的常量，没有强制使用的要求
	显示指定类型的时候，必须确保常量左右值类型一致*/

}
