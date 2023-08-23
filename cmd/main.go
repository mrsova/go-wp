package main

import (
	"context"
	"fmt"
	"math"
	"time"
	"workerpoll/internal/job"
	"workerpoll/pkg/wp"
)

func main() {
	ctx := context.Background()
	go runProcess(ctx, job.GenerateJobs(1000000))
	time.Sleep(10000 * time.Second)
}

func runProcess(ctx context.Context, jobs []wp.Job) {
	start := time.Now()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	jobsCount := len(jobs)
	workersCount := int(math.Round(float64(jobsCount / 2)))
	if workersCount == 0 {
		workersCount = 1
	}
	pool := wp.NewWorkerPool(workersCount, jobsCount)
	pool.AddJobs(jobs)
	go pool.RunWorkerPool(ctx)

	for {
		select {
		case _, ok := <-pool.Results():
			if !ok {
				elapsed := time.Since(start)
				fmt.Printf("DONE: Time over ===============> %s\n", elapsed)
				return
			}
			//fmt.Println(fmt.Sprintf("RESULT - %+v", value))
		default:
		}
	}
}
