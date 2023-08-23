package job

import (
	"context"
	"fmt"
	"time"
	"workerpoll/pkg/wp"
)

type ExecutionFn func(ctx context.Context, value int) (int, error)

type MainJob struct {
	descriptor wp.JobDescriptor
	value      int
	execFn     ExecutionFn
}

func New(id string, value int) MainJob {
	return MainJob{
		value: value,
		descriptor: wp.JobDescriptor{
			Id: id,
		},
		execFn: func(ctx context.Context, value int) (int, error) {
			time.Sleep(200 * time.Millisecond)
			return value * 2, nil
		},
	}
}

func (j MainJob) Execute(ctx context.Context) wp.Result {
	value, err := j.execFn(ctx, j.value)
	if err != nil {
		return wp.Result{
			Err:        err,
			Descriptor: j.descriptor,
		}
	}

	return wp.Result{
		Value:      value,
		Descriptor: j.descriptor,
	}
}

func GenerateJobs(jobsCount int) []wp.Job {
	jobs := make([]wp.Job, jobsCount)
	for i := 0; i < jobsCount; i++ {
		jobs[i] = New(fmt.Sprintf("id%d", i), i)
	}
	return jobs
}
