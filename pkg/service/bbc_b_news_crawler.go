package service

import (
	"bbc-bcrawl/pkg/utils"
	"log"

	"github.com/gocolly/colly"
)

func CrawlMain() {
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML(`main div section[aria-labelledby="Special-reports-1"] ul`, func(h *colly.HTMLElement) {
		// h.DOM.
		h.ForEach("li div h3 a", func(i int, h *colly.HTMLElement) {
			// dc.Visit(h.Request.AbsoluteURL(h.Attr("href")))
			h.Request.Visit(h.Attr("href"))

		})
	})

	if err := c.Visit("https://www.bbc.com/burmese"); err != nil {
		log.Println(err)
	}
	c.Wait()
}

func CrawlInternationalReports(url string) {

	c := colly.NewCollector(
		colly.Async(true),
	)

	// count := 0

	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// in detail page
	c.OnHTML("main[role=main] div", func(h *colly.HTMLElement) {
		log.Println(h.Text)
	})

	c.OnHTML("main div ul[data-testid=topic-promos] li a", func(h *colly.HTMLElement) {
		// log.Println(h.Attr("href"))
		// if count == 2 {
		// 	return
		// }
		c.Visit(h.Attr("href"))

		// if count == 22 || count == 23 {
		// 	c.Visit(h.Attr("href"))
		// }
		// count++
	})

	if err := c.Visit(url); err != nil {
		log.Println(err)
	}
	c.Wait()

}

func CrawlSpecificPageOfInternationalReport(url string) {

	var data []string

	var imgs []string

	imgStart := 0

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML(`main div`, func(h *colly.HTMLElement) {

		h.ForEach("h1", func(i int, h *colly.HTMLElement) {
			data = append(data, h.Text)
		})

		h.ForEach("h2", func(i int, h *colly.HTMLElement) {
			data = append(data, h.Text)
		})

		h.ForEach("img", func(i int, h *colly.HTMLElement) {
			// data = append(data, h.Attr("src"))
			imgs = append(imgs, h.Attr("src"))
		})
		h.ForEach("noscript", func(i int, h *colly.HTMLElement) {
			link, err := utils.GetSrcFromImg(h.Text)
			if err != nil {
				imgs = append(imgs, h.Text)
			} else {
				imgs = append(imgs, link)
			}
		})

		h.ForEach("div > figcaption > p", func(i int, h *colly.HTMLElement) {

			data = append(data, imgs[imgStart])
			// log.Println(imgs[imgStart])
			// log.Println(h.Text)
			imgStart++
			data = append(data, h.Text)
		})

		h.ForEach("div > p", func(i int, h *colly.HTMLElement) {
			data = append(data, h.Text)
		})

	})

	// if err := c.Visit("https://www.bbc.com/burmese/articles/c2lz01w9qgvo"); err != nil {
	if err := c.Visit(url); err != nil {
		log.Println(err)
	}
	c.Wait()

	for k, v := range data {

		// log.Println(k, v)
		_ = k
		_ = v

		// Do Database operation or somthing
	}

}
