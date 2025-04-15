package state

import (
	"fmt"
	"testing"
	"time"
)

func TestThreadPool(t *testing.T) {
	// Create a thread pool with 5 workers
	pool := NewThreadPool(5)
	pool.Start()

	// Add some tasks to the pool
	for i := 0; i < 10; i++ {
		i := i
		pool.AddTask(func() {
			fmt.Printf("Task %d is running\n", i)
			time.Sleep(time.Second) // Simulate some work
			fmt.Printf("Task %d is done\n", i)
		})
	}

	// Stop the pool and wait for all tasks to complete
	pool.Stop()
	fmt.Println("All tasks completed")
}
