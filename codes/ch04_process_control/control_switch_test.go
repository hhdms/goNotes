package ch04_process_control

import (
	"fmt"
	"testing"
)

func TestControlSwitch(t *testing.T) {
	// 1. 基本用法
	var grade = "B"
	switch grade {
	case "A":
		fmt.Println("优秀")
		//break  // 默认使用 break
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// 2. 一个 case 多个值
	var day = 2
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("无效的星期")
	}

	// 3. case 中使用表达式
	var score = 90
	switch {
	case score >= 90:
		fmt.Println("A级")
	case score >= 80:
		fmt.Println("B级")
	case score >= 70:
		fmt.Println("C级")
	case score >= 60:
		fmt.Println("D级")
	default:
		fmt.Println("不及格")
	}

	// 4. fallthrough 关键字(继续执行下一个case)
	var num = 5
	switch num {
	case 5:
		fmt.Println("数字5")
		fallthrough
	case 10:
		fmt.Println("数字10")
	default:
		fmt.Println("其他数字")
	}
}
