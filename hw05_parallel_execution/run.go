package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	taskChannel := make(chan Task) // из этого канала воркеры будут получать задачу
	var errCounter int32
	var wg sync.WaitGroup

	defer close(taskChannel)

	for i := 0; i < n; i++ { // запускаем n-воркеров
		go func() {
			var err error
			for task := range taskChannel {
				if err = task(); err != nil {
					atomic.AddInt32(&errCounter, 1)
				}
				wg.Done()
			}
		}()
	}

	for i := range tasks {
		wg.Add(1)
		if atomic.LoadInt32(&errCounter) >= int32(m) {
			return ErrErrorsLimitExceeded
		}
		taskChannel <- tasks[i]
	}

	wg.Wait()

	return nil
}
