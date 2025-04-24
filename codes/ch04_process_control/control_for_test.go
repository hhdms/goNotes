package ch04_process_control

import (
	"fmt"
	"testing"
)

// 流程控制 for 循环语言
func TestControlFor(t *testing.T) {
	// 1. for 循环基本写法
	fmt.Println("第一个 for 循环")
	for i1 := 0; i1 < 5; i1++ {
		fmt.Println(i1)
	}
	// 2. 省略初始语句, 但必须保留 `;`
	fmt.Println("第二个 for 循环")
	var i2 = 0
	for ; i2 < 5; i2++ {
		fmt.Println(i2)
	}
	// 3. 省略初始语句和结束语句 类似其他语言的 while 循环
	fmt.Println("第三个 for 循环")
	var i3 = 0
	for i3 < 5 {
		fmt.Println(i3)
		i3++ // 不写会变成死循环
	}
	// 4. 无限循环
	//for {
	//	fmt.Println("hello world")
	//}
	// 5. break 跳出循环
	fmt.Println("第四个 for 循环")
	for i4 := 0; i4 < 5; i4++ {
		if i4 == 3 {
			break
		}
		fmt.Println(i4)
	}
	// 6. continue 跳过本次循环, 继续下一次循环
	fmt.Println("第五个 for 循环")
	for i5 := 0; i5 < 5; i5++ {
		if i5 == 3 {
			continue
		}
		fmt.Println(i5)
	}
}
