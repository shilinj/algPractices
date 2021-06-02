package concurrent

import (
	"fmt"
	"time"
)

func read(ch1 chan int, ch2 chan bool) {
	defer fmt.Println("exit read")
	for {
		fmt.Println(time.Now())
		v, ok := <-ch1
		if ok {
			fmt.Printf("read a int is %d\n", v)
		} else {
			ch2 <- true
			break
		}
	}
}

func write(ch chan int) {
	defer fmt.Println("exit write")
	time.Sleep(time.Second * 1)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

}

func TestSyncCocurrent() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	go read(ch1, ch2)
	go write(ch1)
	<-ch2
}
