package routine

import (
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func TestCloseChToBroadcast1() {
	done := make(chan struct{})
	ch := GenerateInt(done)
	fmt.Println("read data = ", <-ch)
	fmt.Println("read data = ", <-ch)
	close(done)
	fmt.Println(<-ch)
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
}
