/*
package main

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
		wg := sync.WaitGroup{}

		wg.Add(n)
		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()
				var (
					input1, input2 int
				)
				wg1 := sync.WaitGroup{}

				select {
				case input1 = <-in1:
				case <-time.After(time.Second * 10):
					fmt.Println("select1")
				}
				select {
				case input2 = <-in2:
				case <-time.After(time.Second * 10):
					fmt.Println("select2")
				}

				wg1.Add(2)
				go func() {
					defer wg1.Done()
					input1 = fn(input1)

				}()
				go func() {
					defer wg1.Done()
					input2 = fn(input2)

				}()
				wg1.Wait()

				out <- input1 + input2

			}()
		}
		wg.Wait()
		close(out)
	}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

func main() {
	var (
		res int32
	)

	inputString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputString = strings.Trim(inputString, "\n")
	inputSlice := strings.Split(inputString, ",")

	number1, _ := strconv.Atoi(inputSlice[0])
	number2, _ := strconv.Atoi(inputSlice[1])

	finalSlice := make([]int, number2)

	for i := 0; i <= number2; i++ {
		finalSlice[i] = number1 + i
	}

	for i := 0; i < len(finalSlice); i++ {

		for j := 1; j < len(finalSlice); j++ {

			if finalSlice[i]/i+2 == finalSlice[i+1]/i+2 {
				atomic.AddInt32(&res, 1)
			}
		}
	}
	fmt.Println(res)
}
