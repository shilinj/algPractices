package concurrence

import (
	"fmt"
	"math/rand"
)

//done 接收通知退出信号
func GenerateintA(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

// done 接收通知退出信号
func GenerateintB(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

//通过 select 执行扇入（Fan in）操作
func Generateint(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Lable:
		for {
			select {
			case ch <- <-GenerateintA(send):
			case ch <- <-GenerateintB(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func TestCloseChToBroadcast2() {
	done := make(chan struct{}) // 创建一个作为接收退出信号的chan
	ch := Generateint(done)     // 启动生成器
	for i := 0; i < 10; i++ {   // 获取生成器资源
		fmt.Println(<-ch)
	}
	done <- struct{}{} // 通知生产者停止生产
	fmt.Println("stop gerηarate")
}
