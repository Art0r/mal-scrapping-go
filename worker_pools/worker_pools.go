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

type NewWorkerPoolsParams struct {
	Jobs            []Job
	NumberOfWorkers int
}

type WorkerPools struct {
	NumberOfWorkers int
	numberOfJobs    int
	jobsQueue       chan Job
	resultsQueue    chan Result
	jobs            []Job
}

func (wp *WorkerPools) NewWorkerPools(newWorkerPoolsParams NewWorkerPoolsParams) *WorkerPools {
	return &WorkerPools{
		numberOfJobs:    len(newWorkerPoolsParams.Jobs) - 1,
		jobsQueue:       make(chan Job, len(newWorkerPoolsParams.Jobs)-1),
		resultsQueue:    make(chan Result, len(newWorkerPoolsParams.Jobs)-1),
		jobs:            newWorkerPoolsParams.Jobs,
		NumberOfWorkers: newWorkerPoolsParams.NumberOfWorkers,
	}
}

func (wp *WorkerPools) Start() {
	for w := 0; w <= wp.NumberOfWorkers; w++ {
		worker := Worker{ID: w, JobsQueue: &wp.jobsQueue, ResultsQueue: &wp.resultsQueue}
		go worker.Work()
	}

	for j := 0; j <= wp.numberOfJobs; j++ {
		wp.jobsQueue <- wp.jobs[j]
	}
	close(wp.jobsQueue)

	for a := 0; a <= wp.numberOfJobs; a++ {
		<-wp.resultsQueue
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
