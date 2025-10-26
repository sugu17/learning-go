package exercises

import (
	"fmt"
	"sync"
)

func Exercise_12_1() {
	ch := make(chan int, 20)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range 10 {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			ch <- i
		}
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()
	for num := range ch {
		fmt.Println("Number:", num)
	}
}
