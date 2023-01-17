package main

import (
	"fmt"
	"sync"
)

func main() {
	ManyGoWait()
}

func ManyGoWait() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}
func hello(i int) {
	println("hello 1 goroutine : " + fmt.Sprint(i))
}
