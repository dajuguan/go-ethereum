package state

import (
	"sync"
)

// Task represents a function that can be executed by the thread pool.
type Task func()

// ThreadPool manages a pool of goroutines to execute tasks concurrently.
type ThreadPool struct {
	workers   int
	taskQueue chan Task
	wg        sync.WaitGroup
}

// NewThreadPool creates a new thread pool with the specified number of workers.
func NewThreadPool(workers int) *ThreadPool {
	return &ThreadPool{
		workers:   workers,
		taskQueue: make(chan Task),
	}
}

// Start initializes the thread pool and starts the worker goroutines.
func (tp *ThreadPool) Start() {
	for i := 0; i < tp.workers; i++ {
		tp.wg.Add(1)
		go tp.worker()
	}
}

// Stop stops the thread pool and waits for all tasks to complete.
func (tp *ThreadPool) Stop() {
	close(tp.taskQueue)
	tp.wg.Wait()
}

// AddTask adds a new task to the task queue.
func (tp *ThreadPool) AddTask(task Task) {
	tp.taskQueue <- task
}

// worker represents a single worker goroutine that executes tasks from the queue.
func (tp *ThreadPool) worker() {
	defer tp.wg.Done()
	for task := range tp.taskQueue {
		task()
	}
}
