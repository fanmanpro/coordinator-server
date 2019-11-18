package worker

type Job struct {
	fn      func(chan error)
	errChan chan error
}

func NewJob(fn func(chan error)) *Job {
	return &Job{fn, make(chan error, 1)}
}
func (j *Job) Perform() {
	j.fn(j.errChan)
}
