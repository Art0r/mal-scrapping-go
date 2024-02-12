package handlers

import (
	"strings"

	dt "github.com/Art0r/mal-scrapping/data_structures"

	"github.com/gocolly/colly/v2"

	wop "github.com/Art0r/mal-scrapping/worker_pools"
)

func GetWorkDemographic(
	c *colly.Collector,
	worksList *dt.LinkedList,
	demographicRelation *dt.DemographicRelation,
	demographicXpath string) {

	defer c.Wait()
	callDemographicScraper(c, demographicRelation, demographicXpath)

	current := worksList.Head

	for current != nil {
		if visited, _ := c.HasVisited(current.Data); !visited {
			c.Visit(current.Data)
		}

		current = current.Next
	}
}

func GetWorkDemographicConcurrently(c *colly.Collector,
	worksList *dt.LinkedList,
	demographicRelation *dt.DemographicRelation,
	demographicXpath string) {

	defer c.Wait()
	callDemographicScraper(c, demographicRelation, demographicXpath)

	jobs := setJobs(c, worksList)

	workerPools := wop.WorkerPools{}

	wpParams := wop.NewWorkerPoolsParams{Jobs: *jobs, NumberOfWorkers: 4}

	workerPools.NewWorkerPools(wpParams).Start()

}

func setJobs(c *colly.Collector, worksList *dt.LinkedList) *[]wop.Job {
	current := worksList.Head

	var jobs []wop.Job
	var i int

	visitAddress := func() interface{} {
		for current != nil {
			if visited, _ := c.HasVisited(current.Data); !visited {
				c.Visit(current.Data)
			}

			i = i + 1
			current = current.Next
		}

		return 0
	}

	jobs = append(jobs, wop.Job{ID: i, Excute: visitAddress})

	return &jobs
}

func callDemographicScraper(
	c *colly.Collector,
	demographicRelation *dt.DemographicRelation,
	demographicXpath string,
) {

	c.OnXML(demographicXpath,
		func(e *colly.XMLElement) {

			category := strings.TrimSpace(e.Text)

			switch strings.ToLower(category) {

			case strings.ToLower(string(dt.Josei)):
				demographicRelation.Josei++
				break
			case strings.ToLower(string(dt.Seinen)):
				demographicRelation.Seinen++
				break
			case strings.ToLower(string(dt.Shonen)):
				demographicRelation.Shonen++
				break
			case strings.ToLower(string(dt.Shojo)):
				demographicRelation.Shojo++
				break
			default:
				break
			}

		})

}
