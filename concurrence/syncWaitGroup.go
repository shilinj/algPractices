package concurrence

import (
	"fmt"
	"sync"
	"time"
)

func cocurSync(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("Doing something ... ...", i)
	time.Sleep(time.Second * 1)
	fmt.Println("Done something ... ...", i)
}

func TestSyncWaitGroup() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go cocurSync(wg, i)
	}
	wg.Wait()
}
