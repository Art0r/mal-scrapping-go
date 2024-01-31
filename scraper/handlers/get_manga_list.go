package handlers

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly/v2"

	dt "github.com/Art0r/mal-scrapping/data_structures"
)

func GetTotalMangaList(c *colly.Collector, itemsXpath string, is_manga bool) dt.LinkedList {
	var mutex sync.Mutex

	mangaList := dt.LinkedList{}

	getMangaList(c, &mangaList, &mutex, itemsXpath)

	for i := 0; i < LIMITS; i++ {

		url := fmt.Sprintf("https://myanimelist.net/topanime.php?limit=%d", i*50)

		if is_manga {
			url = fmt.Sprintf("https://myanimelist.net/topmanga.php?limit=%d", i*50)
		}

		c.Visit(url)

		c.Wait()
	}

	return mangaList
}

func getMangaList(c *colly.Collector, mangaList *dt.LinkedList, mutex *sync.Mutex, itemsXpath string) {

	c.OnXML(itemsXpath,
		func(e *colly.XMLElement) {
			href := e.Attr("href")

			mutex.Lock()
			defer mutex.Unlock()

			mangaList.Append(href)

		})
}
