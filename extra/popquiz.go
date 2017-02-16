package main

import "sync"

func Close(c chan int) {
	select {
	case <-c:
	default:
		close(c)
	}
}

func loop() {
	const N = 2000
	start := make(chan int)
	c := make(chan int)
	var done sync.WaitGroup
	done.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer done.Done()
			<-start
			Close(c)
		}()
	}
	close(start)
	done.Wait()
}

func main() {
	for {
		loop()
	}
}
