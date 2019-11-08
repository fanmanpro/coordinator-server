package worker

var id int

type Worker struct {
	pool     *Pool
	jobQueue chan *Job
	id       int
}

func NewWorker(p *Pool) *Worker {
	id++
	return &Worker{
		p,
		make(chan *Job, 1),
		id,
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.pool.WorkerQueue <- w.jobQueue

			select {
			case work := <-w.jobQueue:
				// Receive a work request.
				//fmt.Printf("worker%d: Received work request:%v\n", w.id, work.Name)
				err := work.Perform()
				work.errchan <- err

				//time.Sleep(work.Delay)
				//fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)

				//case <-w.QuitChan:
				//	// We have been asked to stop.
				//	fmt.Printf("worker%d stopping\n", w.ID)
				//	return
				//}
			}
		}
	}()
}
