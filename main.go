package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

func randomNumber(wg *sync.WaitGroup, sum *atomic.Int32, stop chan struct{}) {
	wg.Add(1)
	defer wg.Done()

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	tick := time.NewTicker(time.Second * 2)
	defer tick.Stop()

	for {
		select {

		case <-tick.C:
			sum.Add(int32(random.Intn(101)))

		case <-stop:
			return
		}
	}
}

func check(wg *sync.WaitGroup, sum, sig *atomic.Int32, stop chan struct{}) {
	wg.Add(1)
	defer wg.Done()

	for sum.Load() != 100 {
		time.Sleep(time.Second * 2)

		if sig.Load() == 1 {
			wg.Done()
			return
		}
		if sum.Load() > 100 {

			fmt.Println(sum.Load())
		}
		sum.Store(0)
	}
	stop <- struct{}{}
	fmt.Println("the process is completed because the amount is 100")
}

func shutdown(wg *sync.WaitGroup, sig *atomic.Int32, stop chan struct{}) {

	wg.Add(1)
	defer wg.Done()

	var stopSignal = make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGKILL)

	select {

	case <-stopSignal:
		stop <- struct{}{}
		sig.Add(1)
		fmt.Println("the process is completed because the completion signal has been received")
		return

	case <-stop:
		return
	}
}
func main() {
	var (
		sum, sig atomic.Int32
		wg       sync.WaitGroup
		stop     = make(chan struct{})
	)

	go randomNumber(&wg, &sum, stop)
	go randomNumber(&wg, &sum, stop)
	go check(&wg, &sum, &sig, stop)
	go shutdown(&wg, &sig, stop)

	wg.Wait()
	close(stop)
}
