package main

import (
	"fmt"
	"sync"
	"time"
)

func Store(m map[int]int, mx *sync.Mutex, key int, value int) {
	mx.Lock()
	m[key] = value
	mx.Unlock()
}

func main() {
	m := make(map[int]int)
	var mx sync.Mutex

	for i := 0; i < 100; i++ {
		go Store(m, &mx, i, i)
	}

	for i, k := range m {
		fmt.Println("key - ", k, "value  - ", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Finish")
}
