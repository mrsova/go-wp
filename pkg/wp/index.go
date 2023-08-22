package wp

import (
	"context"
	"sync"
)

type WorkerPool struct {
	workersCount int
	jobs         chan Job
	results      chan Result
}

func NewWorkerPool(count int, jobCount int) WorkerPool {
	return WorkerPool{
		workersCount: count,
		jobs:         make(chan Job, jobCount),
		results:      make(chan Result, jobCount),
	}
}

func (wp WorkerPool) RunWorkerPool(ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		go workerProcessing(ctx, i, &wg, wp.jobs, wp.results)
	}
	wg.Wait()
	close(wp.results)
}

func (wp WorkerPool) Results() <-chan Result {
	return wp.results
}

func (wp WorkerPool) AddJobs(jobs []Job) {
	for i, _ := range jobs {
		wp.jobs <- jobs[i]
	}
	close(wp.jobs)
}

func workerProcessing(ctx context.Context, numWorker int, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()
	for job := range jobs {
		results <- job.Execute(ctx)
	}
}
