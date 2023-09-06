package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(j int, wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
			// fmt.Printf("i => %v counter => %v\n", j, counter)
			wg.Done()
		}(i, &wg)

	}

	wg.Wait()
}
