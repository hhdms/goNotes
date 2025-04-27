package ch04_process_control

import (
	"fmt"
	"testing"
)

func TestControlGoto(t *testing.T) {
	// 定义局部变量
	var a = 10
loop:
	for a < 20 {
		if a == 15 {
			// 跳过迭代
			a = a + 1
			goto loop
		}
		fmt.Printf("a的值为:%d\n", a)
		a++
	}

	// 跳出多层循环
	for i := 0; i < 10; i++ {
		// 内层循环
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 跳转到标签
				goto end
			}
		}
	}
	// 手动返回, 避免进入标签
	return
end:
	fmt.Println("done")
}
