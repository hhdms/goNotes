package ch03_operator

import (
	"fmt"
	"testing"
)

// 运算符
func TestOperator(t *testing.T) {
	// 1. 算数运算符
	a := 985
	b := 211
	var c int
	c = a + b
	fmt.Printf("相加：c = %d\n", c)
	c = a - b
	fmt.Printf("相减：c = %d\n", c)
	c = a * b
	fmt.Printf("相乘：c = %d\n", c)
	c = a / b
	fmt.Printf("相除：c = %d\n", c)
	c = a % b
	fmt.Printf("取余：c = %d\n", c)

	// 2. 关系运算符
	fmt.Println("10 == 10", 10 == 10) // true
	fmt.Println("10 != 10", 10 != 10) // false
	fmt.Println("10 > 10", 10 > 10)   // false
	fmt.Println("10 >= 10", 10 >= 10) // true
	fmt.Println("10 < 10", 10 < 10)   // false
	fmt.Println("10 <= 10", 10 <= 10) // true

	// 3. 逻辑运算符
	fmt.Println(10 == 10 && 20 > 10) // true
	fmt.Println(10 == 10 || 20 > 10) // true
	fmt.Println(!(10 == 10))         // false

	// 4. 位运算符
	m := 1              // 001
	n := 5              // 101
	fmt.Println(m & n)  // 001
	fmt.Println(m | n)  // 101
	fmt.Println(m ^ n)  // 100
	fmt.Println(m << 2) // 100
	fmt.Println(m >> 2) // 001

	// 5. 赋值运算符
	var a1 int
	a1 = 10
	a1 += 10 // a1 = a1 + 10
	fmt.Println(a1)
}
