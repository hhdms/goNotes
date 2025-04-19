package ch02_const

import (
	"fmt"
	"testing"
)

// 声明常量
const pi = 3.1415

//const e = 2.7182

// 批量声明
const (
	a1 = 10
	a2
	a3
)

// 常量
func TestConstant(t *testing.T) {
	//fmt.Println(pi, e)
	fmt.Println(a1, a2, a3)
}
