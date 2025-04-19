package ch02_const

import (
	"fmt"
	"testing"
)

const (
	January = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// 跳过某些值
//const (
//	n1 = iota
//	n2
//	_
//	n4
//)

// 中间插队
const (
	n1 = iota
	n2 = 100
	n3 = iota
	n4
)
const n5 = iota

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 多个 iota 在同一行
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

// iota 用法
func TestIota(t *testing.T) {
	//fmt.Println(December)
	//fmt.Println(n1, n2, n3, n4)
	//fmt.Println(n5)

	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
}
