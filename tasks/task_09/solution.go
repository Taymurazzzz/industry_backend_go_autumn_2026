package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if len(in) == 0 {
		return []Result[R]{}, nil
	}
	results := make([]Result[R], len(in))
	jobs := make(chan int)
	workerCount := min(workers, len(in))
	var wg sync.WaitGroup
	wg.Add(workerCount)
	work := func() {
		defer wg.Done()
		for i := range jobs {
			if ctx.Err() != nil {
				return
			}
			res, err := fn(ctx, in[i])
			if err != nil {
				results[i].Err = err
			} else {
				results[i].Value = res
			}
		}
	}
	for i := 0; i < workerCount; i++ {
		go work()
	}
	for i := range in {
		if ctx.Err() != nil {
			break
		}
		select {
		case jobs <- i:
		case <-ctx.Done():
			break
		}
	}
	close(jobs)
	wg.Wait()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return results, nil
}
