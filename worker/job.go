package worker

type Job struct {
	fn      func() error
	errchan chan error
}

func NewJob(fn func() error, errchan chan error) *Job {
	return &Job{fn, errchan}
}
func (j *Job) Perform() error {
	return j.fn()
}
