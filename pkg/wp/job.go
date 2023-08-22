package wp

import "context"

type JobDescriptor struct {
	Id string
}

type Result struct {
	Descriptor JobDescriptor
	Value      interface{}
	Err        error
}

type Job interface {
	Execute(ctx context.Context) Result
}
