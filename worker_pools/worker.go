package workerpools

import "fmt"

type Job struct {
	ID     int
	Excute func() interface{}
}

type Result struct {
	jobID int
}

type Worker struct {
	ID           int
	JobsQueue    *chan Job
	ResultsQueue *chan Result
}

func (w *Worker) Work() {
	for j := range *w.JobsQueue {
		fmt.Println("worker", w.ID, "started  job", j)
		j.Excute()
		fmt.Println("worker", w.ID, "finished job", j)
		*w.ResultsQueue <- Result{jobID: j.ID}
	}
}
