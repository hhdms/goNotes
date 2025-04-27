package ch04_process_control

import (
	"fmt"
	"testing"
)

func TestControlRange(t *testing.T) {
	// 1. 遍历切片
	var nums = []int{1, 2, 3, 4, 5}
	for _, n := range nums {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	// 2. 遍历字符串
	var str = "hello"
	for i, char := range str {
		fmt.Printf("索引:%d, 值:%c \n", i, char)
	}
	// 3. 遍历 map 映射
	var scores = map[string]int{"张三": 90, "李四": 80, "王五": 70}
	for k, v := range scores {
		fmt.Printf("姓名: %s, 成绩: %d\n", k, v)
	}
	// 4. 只取值
	for i := range nums {
		fmt.Printf("索引:%d \n", i)
	}
	// 5. 只取值
	for _, n := range nums {
		fmt.Printf("值:%d \n", n)
	}
}
