package main

import (
	"fmt"
	"time"

	"github.com/Art0r/mal-scrapping/scraper"

	wop "github.com/Art0r/mal-scrapping/worker_pools"
)

func main() {

	// Worker Pools 49.089201851s
	// Default 1m38.904498968s

	start := time.Now()

	jobs := []wop.Job{
		{ID: 1, Excute: func() interface{} { return scraper.Scraper(false) }},
		{ID: 2, Excute: func() interface{} { return scraper.Scraper(true) }},
		{ID: 3, Excute: func() interface{} { return scraper.Scraper(true) }},
		{ID: 4, Excute: func() interface{} { return scraper.Scraper(false) }},
	}

	workerPools := wop.WorkerPools{}

	wpParams := wop.NewWorkerPoolsParams{Jobs: jobs, NumberOfWorkers: 4}

	workerPools.NewWorkerPools(wpParams).Start()

	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
