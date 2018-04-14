package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	callbackF("Bill", func(name string) {
		println("Hello " + name)
		wg.Done()
	})
	callbackF("Joe ", func(name string) {
		println("Hello " + name)
		wg.Done()
	})
	wg.Wait()

	println("all done")
}

func callbackF(name1 string, f func(name string)) {
	go f(name1)
}
