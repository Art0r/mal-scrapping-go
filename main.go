package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Art0r/mal-scrapping/scraper"
)

func main() {
	starTime := time.Now()

	numberOfJobs := 2
	var wg sync.WaitGroup

	jobs := make(chan bool, 2)
	results := make(chan bool, 2)
	
	args := []bool{true, false}

	for i := 0; i < numberOfJobs; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for i := 0; i < numberOfJobs; i++ {
		jobs <- args[i]
	}

	close(jobs)
	wg.Wait()
	close(results)

	elapsed := time.Since(starTime)

	fmt.Println("Elapsed time: ", elapsed)
}

func worker(id int, jobs <-chan bool, results chan<- bool, wg *sync.WaitGroup)  {
	defer wg.Done()
	
	for job := range jobs {
		fmt.Printf("FUNCTION %d STARTED\n", id + 1)
		results <- scraper.Scraper(job)
		fmt.Printf("FUNCTION %d FINISHED\n", id + 1)
	}
}