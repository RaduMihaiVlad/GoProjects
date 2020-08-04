package main

import (
	"fmt"
	"sync"
	"time"
)

func Get(c chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for x := range c {
		fmt.Println(x)
	}

}

func Put(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(c)
	for i := 0; i < 5; i++ {
		// c <- i
	}
}

func main() {

	chanOwner := func() <-chan int {
		results := make(chan int, 5) 
		go func() {
			defer close(results)
			for i := 0; i <= 20; i++ {
				results <- i
				fmt.Printf("Added value %d\n", i)
			}
		}()
		return results
	}
	
	consumer := func(results <-chan int) { 
		for result := range results {
			fmt.Printf("Received: %d\n", result)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Done receiving!")
	}
	
	results := chanOwner()        
	consumer(results)
}