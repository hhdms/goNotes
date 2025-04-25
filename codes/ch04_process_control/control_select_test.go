package ch04_process_control

import (
	"fmt"
	"testing"
	"time"
)

func TestControlSelect(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "来自通道1的信息"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自通道2的信息"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}
