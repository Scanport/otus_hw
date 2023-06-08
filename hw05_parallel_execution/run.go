package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	taskCh := make(chan Task) // из этого канала воркеры будут получать задачу
	resCh := make(chan error, n)

	var errCounter int
	var wg sync.WaitGroup

	for i := 0; i < n; i++ { // запускаем n-воркеров
		go func() {
			for task := range taskCh {
				if task == nil {
					return
				}
				resCh <- task()
				wg.Done()
			}
		}()
	}
	i := 0
LOOP:
	for {
		select {
		case err := <-resCh:
			if err != nil {
				errCounter++
			}
			if errCounter == m {
				break LOOP
			}
		default:
			if i > len(tasks)-1 {
				break LOOP
			}
			wg.Add(1)
			taskCh <- tasks[i]
			i++
		}
	}
	wg.Wait()
	close(taskCh)
	close(resCh)
	if errCounter >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
