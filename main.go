package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func randomNumber(wg *sync.WaitGroup, numChan chan int, stop chan struct{}, seed int64) {

	wg.Add(1)
	defer wg.Done()

	tick := time.NewTicker(time.Second * 2)
	defer tick.Stop()

	random := rand.New(rand.NewSource(seed))

	for {
		select {
		case <-tick.C:
			numChan <- random.Intn(101)
		case <-stop:
			return
		}
	}
}

func check(wg *sync.WaitGroup, num1Chan, num2Chan chan int, stop chan struct{}) {

	wg.Add(1)
	defer wg.Done()

	var (
		num1, num2 int
	)

	for {
		select {
		case num1 = <-num1Chan:
			select {
			case num2 = <-num2Chan:
				sum := num1 + num2
				log.Println("num1:", num1, "|", "num2:", num2, "|", "sum:", sum) //Проверка логики
				if sum == 100 {
					for i := 0; i < 3; i++ {
						stop <- struct{}{}
					}
					fmt.Println("the process is completed because the amount is 100")
					return
				}
				if sum >= 200 {
					fmt.Println("Error in logic")
				}
				if sum > 100 {
					fmt.Println(sum)

				}
			case <-stop:
				return
			}
		case <-stop:
			return
		}
	}
}

func shutdown(wg *sync.WaitGroup, stop chan struct{}) {

	wg.Add(1)
	defer wg.Done()

	var stopSignal = make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {
	case <-stopSignal:
		for i := 0; i < 3; i++ {
			stop <- struct{}{}
		}
		fmt.Println("the process is completed because the completion signal has been received")
		return

	case <-stop:
		return
	}
}

func main() {
	var (
		num1Chan, num2Chan = make(chan int), make(chan int)
		wg                 sync.WaitGroup
		stop               = make(chan struct{}, 1)
		seed1              = time.Now().UnixNano()
		seed2              = time.Now().UnixNano() * time.Now().UnixNano()
	)

	go randomNumber(&wg, num1Chan, stop, seed1)
	go randomNumber(&wg, num2Chan, stop, seed2)
	go check(&wg, num1Chan, num2Chan, stop)
	go shutdown(&wg, stop)

	wg.Wait()
	close(stop)
}
