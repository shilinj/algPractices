package concurrent


import (
	"fmt"
	"rand"
	"runtime"
)


func GenInt(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
		Label:
			for {
				select {
					case ch <- rand.Int():
						fmt.Println("Read an int into chan.")
					case <- done:
						break Label
				}
			}
		close(ch)
	}()
	return ch
}

func TestGenInt() {
	done := make(chan struct{})
	ch := GenInt(done)
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("NumGoroutine=", runtime.NumGoroutine())

}