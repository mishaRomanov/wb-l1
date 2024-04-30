// 18. Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
package main

import (
	"fmt"
	"sync"
)

// Concurrent Counter структура
type CCounter struct {
	mu      *sync.Mutex
	counter int64
}

// инкремент счетчика
func (c *CCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// выводим значение
func (c *CCounter) Display() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println(c.counter)
}

// конструктор
func NewCCounter() *CCounter {
	return &CCounter{
		mu:      &sync.Mutex{},
		counter: 0,
	}
}

func main() {
	newcc := NewCCounter()

	for range 100 {
		go newcc.Increment()
	}

	newcc.Display()
}
