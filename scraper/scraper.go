package scraper

import (
	"fmt"
	dt "github.com/Art0r/mal-scrapping/data_structures"
	handlers "github.com/Art0r/mal-scrapping/scraper/handlers"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

func Scraper(is_manga bool) bool {

	c := colly.NewCollector(
		colly.AllowedDomains("anilist.co", "myanimelist.net"),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	var demographicRelation dt.DemographicRelation
	var demographicXpath string
	var itemXpath string

	demographicXpath = handlers.ANIME_DEMOGRAPHIC_XPATH
	itemXpath = handlers.ANIME_ITEMS_XPATH

	if is_manga {
		demographicXpath = handlers.MANGA_DEMOGRAPHIC_XPATH
		itemXpath = handlers.MANGA_ITEMS_XPATH

	}

	worksList := handlers.GetTotalMangaList(c, itemXpath, is_manga)

	handlers.GetWorkDemographic(c,
		&worksList, &demographicRelation, demographicXpath)

	totalWorks := calculateTotalWorks(&demographicRelation)

	data := setOutput(&demographicRelation, totalWorks, is_manga)

	writeOutput(data)

	return true

}

func setOutput(demographicRelation *dt.DemographicRelation, totalWorks float32, is_manga bool) string {
	shonenRatio := demographicRelation.Shonen / totalWorks
	shojoRatio := demographicRelation.Shojo / totalWorks
	seinenRatio := demographicRelation.Seinen / totalWorks
	joseiRatio := demographicRelation.Josei / totalWorks
	originalRatio := demographicRelation.NovelWebcomicOriginal / totalWorks

	currentTime := time.Now()
	formattedCurrentTime := currentTime.Format("02-01-2006 15:04:05")

	work_type := "ANIMES"
	platform := "Original Work"

	if is_manga {
		work_type = "MANGAS"
		platform = "Novel or Webcomic"
	}

	data := fmt.Sprintf(
		`Time: %s
--------------------------------------
Top %d %s

Shonen:
	Total %d;
	Ratio %.2f;

Shojo:
	Total %d;
	Ratio %.2f;

Seinen:
	Total %d;
	Ratio %.2f;

Josei:
	Total %d;
	Ratio %.2f;

%s:
		Total %d;
		Ratio %.2f;
--------------------------------------`,
		formattedCurrentTime,
		uint8(totalWorks),
		work_type,

		uint8(demographicRelation.Shonen),
		shonenRatio,

		uint8(demographicRelation.Shojo),
		shojoRatio,

		uint8(demographicRelation.Seinen),
		seinenRatio,

		uint8(demographicRelation.Josei),
		joseiRatio,

		platform,
		uint8(demographicRelation.NovelWebcomicOriginal),
		originalRatio)

	return data
}

func writeOutput(data string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(data + "\n"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Output appended in log.txt")
	fmt.Println(data)
}

func calculateTotalWorks(demographicRelation *dt.DemographicRelation) float32 {
	sum := demographicRelation.Josei +
		demographicRelation.Seinen +
		demographicRelation.Shojo +
		demographicRelation.Shonen

	demographicRelation.NovelWebcomicOriginal = (float32(handlers.LIMITS) * 50) - sum

	sum += demographicRelation.NovelWebcomicOriginal

	return sum
}
