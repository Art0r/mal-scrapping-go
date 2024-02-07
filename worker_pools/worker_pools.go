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

type WorkerPools struct {
	NumberOfJobs int
	JobsQueue    chan Job
	ResultsQueue chan Result
	jobs         []Job
}

func (wp *WorkerPools) NewWorkerPools(jobs []Job) *WorkerPools {
	return &WorkerPools{
		NumberOfJobs: len(jobs) - 1,
		JobsQueue:    make(chan Job, len(jobs)-1),
		ResultsQueue: make(chan Result, len(jobs)-1),
		jobs:         jobs,
	}
}

func (wp *WorkerPools) Start() {
	for w := 0; w <= 3; w++ {
		worker := Worker{ID: w, JobsQueue: &wp.JobsQueue, ResultsQueue: &wp.ResultsQueue}
		go worker.Work()
	}

	for j := 0; j <= wp.NumberOfJobs; j++ {
		wp.JobsQueue <- wp.jobs[j]
	}
	close(wp.JobsQueue)

	for a := 0; a <= wp.NumberOfJobs; a++ {
		<-wp.ResultsQueue
	}
}

func (w *Worker) Work() {
	for j := range *w.JobsQueue {
		fmt.Println("worker", w.ID, "started  job", j)
		j.Excute()
		fmt.Println("worker", w.ID, "finished job", j)
		*w.ResultsQueue <- Result{jobID: j.ID}
	}
}
