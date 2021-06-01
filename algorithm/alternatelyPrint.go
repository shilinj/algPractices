package algorithm

import (
	"fmt"
	"sync"
)

/*
	使用2个goroutine交替打印数字
	goroutine1 打印 1 3 5 7 9
	goroutine2 打印 2 4 6 8 10
*/

func AlternatelyPrintV1() {
	numChan := make(chan int)
	exitChan := make(chan struct{})

	go func() {
		for i := 2; i <= 10; i = i + 2 {
			if num, ok := <-numChan; ok && num < 10 {
				fmt.Println("goroutine1:", num)
			}
			numChan <- i
		}
	}()

	go func() {
		defer close(exitChan)
		for j := 1; j < 10; j = j + 2 {
			numChan <- j
			if num, ok := <-numChan; ok && num <= 10 {
				fmt.Println("goroutine2:", num)
			}
		}
	}()

	<-exitChan
}

func AlternatelyPrintV2() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= 9; i += 2 {
			fmt.Println("goroutine11:", i)
			ch <- 1
			<-ch
		}
	}(ch, &wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-ch
			fmt.Println("goroutine22:", i)
			ch <- 1
		}
	}(ch, &wg)

	wg.Wait()
}
