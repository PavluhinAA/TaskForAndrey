package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func randomNumber(wg *sync.WaitGroup, sum *atomic.Int32, stop chan struct{}) {
	wg.Add(1)
	defer wg.Done()
	tick := time.NewTicker(time.Second * 2)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			sum.Add(rand.Int31n(101))
		case <-stop:
			return
		}
	}
}

func check(wg *sync.WaitGroup, sum *atomic.Int32, stop chan struct{}) {
	wg.Add(1)
	defer wg.Done()

	for sum.Load() != 100 {

		if sum.Load() > 100 {
			fmt.Println(sum.Load())
			sum.Store(0)
		}
	}
	stop <- struct{}{}
}
func shutdown(wg *sync.WaitGroup, sum *atomic.Int32, stop chan struct{}) {

}
func main() {
	var (
		sum  atomic.Int32
		wg   sync.WaitGroup
		stop = make(chan struct{})
	)

	go randomNumber(&wg, &sum, stop)
	go randomNumber(&wg, &sum, stop)
	go check(&wg, &sum, stop)
	go shutdown(&wg, &sum, stop)

	wg.Wait()
	close(stop)
}
