package exercises

import "fmt"

func Exercise_12_2() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
	}()

	for {
		select {
		case val, ok := <-ch1:
			if !ok {
				ch1 = nil
				fmt.Println("Channel 1 closed")
			} else {
				fmt.Printf("Goroutine 1: %d\n", val)
			}
		case val, ok := <-ch2:
			if !ok {
				ch2 = nil
				fmt.Println("Channel 2 closed")
			} else {
				fmt.Printf("Goroutine 2: %d\n", val)
			}
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}

	fmt.Println("All values read, exiting")
}
