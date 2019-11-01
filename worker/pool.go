package worker

type Pool struct {
	workerCap   int
	jobQueueCap int
	JobQueue    chan *Job
	WorkerQueue chan chan *Job
}

func NewPool(wc int, jqc int) *Pool {
	return &Pool{wc, jqc, make(chan *Job, 100), make(chan chan *Job, 5)}
}

func (p *Pool) Start() {
	for i := 0; i < cap(p.WorkerQueue); i++ {
		//fmt.Printf("Starting Worker (%v)\n", i)
		worker := NewWorker(p) // NewWorker(i+1, WorkerQueue)
		worker.Start()
	}
	go func() {
		for {
			select {
			case job := <-p.JobQueue:
				//fmt.Println("Received Job")
				go func() {
					worker := <-p.WorkerQueue
					//fmt.Println("Dispatching Job")
					worker <- job
				}()
			}
		}
	}()
}

func (p *Pool) ScheduleJob(fn func() error) error {
	errchan := make(chan error, 1)
	job := NewJob(fn)
	p.JobQueue <- job
	return <-errchan
}
