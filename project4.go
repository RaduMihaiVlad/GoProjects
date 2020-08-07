package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	return rand.Int()
}

func generatePrime(done <- chan struct{}, random <- chan int, N int) chan int {

		output := make(chan int)
		go func() {
			defer close(output)
			for i := 0; i < N; i++ {
				select {
				case <- done:
					return
				case x := <- random:
					found := false
					for i := 2; i * i <= x; i++ {
						if x % i == 0 {
							found = true
							break
						}
					}
					if !found {
						output <- x
					}
				}
			}
		}()
		
		return output
	}

func generateNumbers(done <-chan struct{}, values ...int) chan int {
		
		valueChan := make(chan int)
		
		go func() {

			defer close(valueChan)
			for {
				for _, value := range values {
					select {
					case <- done:
						return
					case valueChan <- value:
					}
				}
			}
		}()

		return valueChan
	}

func generateNNumbers(done <-chan struct{}, N int, values ...int) chan int {

		valueChan := make(chan int)
		
		go func() {

			defer close(valueChan)

			for i := 0; i < N; i++ {
				for _, value := range values {
					select {
					case <- done:
						return
					case valueChan <- value:
					}
				}
			}
		}()

		return valueChan
}

func fannIn(done chan struct{}, channels ...chan int) chan int {
		fannOut := make(chan int)
		numChanClosed := 0
		totalChan := len(channels)
		for _, c := range channels {
			go func(c1 chan int) {
				for value := range c1 {
					select {
					case <- done:
						return
					default:
						fannOut <- value
					}
				}
				numChanClosed = numChanClosed + 1
				if numChanClosed == totalChan {
					close(fannOut)
				}
			}(c)
		}
		
		return fannOut

	}


func main() {

	/*
		I compared the execution time between a function that writes data in a channel
		and use that data vs 4 channels that uses the same amount of data and a channel
		that listen on the 4 channels.
		The first execution time is ~ 3.1 ms
		The second execution time is ~2.0 ms

	*/

	start := time.Now()
	done := make(chan struct{})
	defer close(done)

	outChan1 := generatePrime(done, generateNNumbers(done, 2050, generateRandomNumber(), 
									generateRandomNumber(), generateRandomNumber()), 2050)
	for x := range outChan1 {
		fmt.Println(x)
	}

	fmt.Println(time.Since(start))

	start = time.Now()

	var chans [4]chan int
	for i := range chans {
	    chans[i] = generatePrime(done, generateNNumbers(done, 500, generateRandomNumber(), 
	   				generateRandomNumber(), generateRandomNumber()), 500)
	}

	fannOutChan := fannIn(done, chans[0], chans[1], chans[2], chans[3])

	for x := range fannOutChan {
		fmt.Println(x)
	}
	fmt.Println(time.Since(start))


}