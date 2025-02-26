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

func randomNumber(wg *sync.WaitGroup, num *atomic.Int32, stop chan struct{}, seed int64) {
	wg.Add(1)
	defer wg.Done()

	tick := time.NewTicker(time.Second * 2)
	defer tick.Stop()

	random := rand.New(rand.NewSource(seed))

	for {
		select {
		case <-tick.C:
			num.Add(int32(random.Intn(101)))
			fmt.Println("number:", num.Load())
		case <-stop:
			fmt.Println("stopping")
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
			return
		}
	}
}

func check(wg *sync.WaitGroup, num1, num2, sig *atomic.Int32, stop chan struct{}) {
	wg.Add(1)
	defer wg.Done()

	for {

		if sig.Load() == 1 {
			return
		}

		if num1.Load() != 0 && num2.Load() != 0 {
			sum := num1.Load() + num2.Load()
			num1.Store(0)
			num2.Store(0)
			if sum == 100 {
				break
			}

			if sum > 100 {

				fmt.Println(sum)

			}
		}
	}
	stop <- struct{}{}
	fmt.Println("the process is completed because the amount is 100")
}

func shutdown(wg *sync.WaitGroup, sig *atomic.Int32, stop chan struct{}) {

	wg.Add(1)
	defer wg.Done()

	var stopSignal = make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {

	case <-stopSignal:
		stop <- struct{}{}
		sig.Add(1)
		fmt.Println("the process is completed because the completion signal has been received")
		return

	case <-stop:
		fmt.Println("stopShutdown")
		return
	}
}

func main() {
	var (
		num1, num2, sig atomic.Int32
		wg              sync.WaitGroup
		stop            = make(chan struct{})
		seed1           = time.Now().UnixNano()
		seed2           = time.Now().UnixNano() * time.Now().UnixNano()
	)

	go randomNumber(&wg, &num1, stop, seed1)
	go randomNumber(&wg, &num2, stop, seed2)
	go check(&wg, &num1, &num2, &sig, stop)
	go shutdown(&wg, &sig, stop)

	wg.Wait()
	close(stop)
}
