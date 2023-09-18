package main

//Реализовать конкурентную запись данных в map.
import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu   sync.Mutex // также можно использовать RWMutex
	data map[int]int
}

func NewConcurrentMap() ConcurrentMap {
	return ConcurrentMap{
		mu:   sync.Mutex{},
		data: make(map[int]int)}
}

func (cm *ConcurrentMap) Set(id, key int) {
	cm.mu.Lock()
	cm.data[id] = key
	cm.mu.Unlock()
}

func (cm *ConcurrentMap) Get(id int) int {
	cm.mu.Lock()
	key := cm.data[id]
	cm.mu.Unlock()
	return key
}

func main() {
	cm := NewConcurrentMap()
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			cm.Set(i, i*2)
			fmt.Printf("Set %d key for %d id\n", i*2, i)

		}(i)
		go func(i int) {
			defer wg.Done()
			x := cm.Get(i)
			fmt.Printf("Got %d key for %d id\n", x, i)
		}(i)
	}
	wg.Wait()
}
