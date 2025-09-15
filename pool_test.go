package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "default"
		},
	}
	pool.Put("Vito")
	pool.Put("Jihad")
	pool.Put("Abdilah")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			time.Sleep(1 * time.Second)
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Done")
}
