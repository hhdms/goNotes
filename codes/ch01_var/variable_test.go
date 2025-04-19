package ch01_var

import (
	"fmt"
	"testing"
)

// 全局变量
var p = 100

func foo() (int, string) {
	return 10, "墨水"
}

// 变量相关的内容
func TestVariable(t *testing.T) {
	// 变量声明格式: var 变量名 变量类型
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)

	// 批量声明
	var (
		a string
		b int
		c bool
		d float64
	)
	fmt.Println(a, b, c, d)

	// 变量初始化: var 变量名 变量类型 = 表达式
	var name1 string = "Tom"
	var age1 int = 18
	var height1 float64 = 170.50
	fmt.Printf("姓名：%s 年龄：%d 身高：%.1f\n", name1, age1, height1)

	// 声明多个变量
	var name2, age2, height2 = "墨水", 20, 170.51
	// name1, age1, height1 := "墨水", 20, 170.51 // 短变量形式
	fmt.Printf("姓名：%s 年龄：%d 身高：%.1f\n", name2, age2, height2)

	// 短变量声明, 只能用在函数内, 函数外无法使用
	// 局部变量
	m := 10
	n := 20
	fmt.Printf("m = %d n = %d\n", m, n)

	// 类型推导, 编辑器根据初始值自动推到类型
	var name3 = "墨水"
	var age3 = 20
	var height3 = 170.50
	fmt.Printf("姓名：%s 年龄：%d 身高：%.1f\n", name3, age3, height3)

	// 匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Printf("x = %d y = %s\n", x, y)
}
