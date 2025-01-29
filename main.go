package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		sum int32
	)

	stop := make(chan struct{})
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		defer wg.Done()
		tick := time.NewTicker(time.Second)
		defer tick.Stop()

		for sum != 100 {
			select {
			case <-tick.C:
				atomic.AddInt32(&sum, rand.Int31n(101))
			case <-stop:
				return
			}
		}

	}()

	go func() {
		defer wg.Done()
		tick := time.NewTicker(time.Second * 5)
		defer tick.Stop()

		for sum != 100 {
			select {
			case <-tick.C:
				atomic.AddInt32(&sum, rand.Int31n(11))
			case <-stop:
				return
			}
		}

	}()

	for sum != 100 {

		if sum > 100 {
			fmt.Println(sum)
			sum = 0
		}
	}
	wg.Wait()
}
