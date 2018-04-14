package main

import (
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		println("Hello World1")
		wg.Done()
	}()

	go func() {
		println("Hello World2")
		wg.Done()
	}()

	wg.Wait()

	println("all done")

}
