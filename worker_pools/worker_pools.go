package workerpools

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
	for w := 1; w <= wp.NumberOfWorkers; w++ {
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