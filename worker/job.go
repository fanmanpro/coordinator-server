package worker

type Job struct {
	fn func() error
}

func NewJob(fn func() error) *Job {
	return &Job{fn}
}
func (j *Job) Perform() error {
	return j.fn()
}
