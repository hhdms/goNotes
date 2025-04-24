package ch04_process_control

import (
	"fmt"
	"testing"
)

// 流程控制 if 判断语句
func TestControlIf(t *testing.T) {
	// 1. 基本写法
	var score1 = 65
	if score1 >= 90 {
		fmt.Println("A级")
	} else if score1 >= 75 {
		fmt.Println("B级")
	} else {
		fmt.Println("C级")
	}
	// 2. 特殊写法; score2 只在代码块中生效
	if score2 := 75; score2 >= 90 {
		fmt.Println("A级")
	} else if score2 >= 75 {
		fmt.Println("B级")
	} else {
		fmt.Println("C级")
	}
	// 3. 可以只使用 if , 不使用 else
	var score3 = 100
	if score3 >= 90 {
		fmt.Println("A+级")
	}
}
