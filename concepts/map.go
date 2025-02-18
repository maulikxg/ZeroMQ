package main

import (
	"fmt"
	"sync"
)

func main(){

	var rw sync.RWMutex
	m := make(map[int]int)
	
	go func() {
		for i := 0; i < 1000; i++ {
			rw.Lock()
			m[i] = i
			rw.Unlock()
		}
	}()
	
	go func() {
		for i := 0; i < 1000; i++ {
			rw.RLock()
			fmt.Println(m[i])
			rw.RUnlock()
		}
	}()
	
}