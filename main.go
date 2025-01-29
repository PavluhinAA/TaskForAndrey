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

/*package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func merge2Channels(fn func(int) int, in1 chan int, in2 <-chan int, out chan<- int, n int) {

	FX1 := func(x1 int, chan1 chan<- int, wg *sync.WaitGroup) {
		defer close(chan1)
		defer wg.Done()
		chan1 <- fn(x1)
	}

	FX2 := func(x2 int, chan2 chan<- int, wg *sync.WaitGroup) {
		defer close(chan2)
		defer wg.Done()
		chan2 <- fn(x2)
	}

	funC := func() {
		wg := new(sync.WaitGroup)
		defer close(out)

		for i := 0; i < n; i++ {
			chan1 := make(chan int)
			chan2 := make(chan int)
			x1, x2 := <-in1, <-in2
			wg.Add(2)
			go FX1(x1, chan1, wg)
			go FX2(x2, chan2, wg)

			y1, y2 := <-chan1, <-chan2
			res := y1 + y2
			out <- res

		}
		wg.Wait()
	}
	go funC()
}*/
