package main

import (
	"fmt"
	"time"

	"github.com/Art0r/mal-scrapping/scraper"

	wop "github.com/Art0r/mal-scrapping/worker_pools"
)

func main() {

	start := time.Now()

	jobs := []wop.Job{
		{ID: 1, Excute: func() interface{} { return scraper.Scraper(false) }},
		{ID: 2, Excute: func() interface{} { return scraper.Scraper(true) }},
	}

	workerPools := wop.WorkerPools{}

	wpParams := wop.NewWorkerPoolsParams{Jobs: jobs, NumberOfWorkers: len(jobs)}

	workerPools.NewWorkerPools(wpParams).Start()

	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
