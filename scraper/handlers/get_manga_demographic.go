package handlers

import (
	dt "github.com/Art0r/mal-scrapping/data_structures"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetWorkDemographic(
	c *colly.Collector,
	worksList *dt.LinkedList,
	demographicRelation *dt.DemographicRelation,
	demographicXpath string) {

	callDemographicScraper(c, demographicRelation, demographicXpath)

	current := worksList.Head

	for current != nil {
		if visited, _ := c.HasVisited(current.Data); !visited {
			c.Visit(current.Data)
		}

		current = current.Next
	}

	c.Wait()

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
