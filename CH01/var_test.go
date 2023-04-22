package CH01

import (
	"testing"
)

// 全局变量不能用简明声明
var (
	age    = 22
	lesson = "go"
)

func TestVariable(t *testing.T) {
	//go是静态语言，静态语言和动态语言相比变量差异很大
	//1．变量必须先定义后使用2．变量必须有类型3．类型定下来后不能改变/定义变量的方式
	var name int
	name = 1
	//变量定义并赋值
	var name2 string = "rdf"
	//简明变量赋值
	name3 := "rdf2"
	//多变量定义
	name4, name5, name6 := "rdf", "rdf", "rdf"
	t.Log(name)
	t.Log(name2)
	t.Logf("name is %s", name3)
	t.Log(name4)
	t.Log(name5)
	t.Log(name6)
	t.Logf("age is %d,lesson is %s", age, lesson)
}
