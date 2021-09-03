package main

import (
	"fmt"
	"sync"	
)

var wg sync.WaitGroup

func main() {
	// c := make(chan struct{}, 0)
	wg.Add(1)
	fmt.Println("Hello, WebAssembly!")
	web.RegisterCallbacks()
	wg.Wait()
}
