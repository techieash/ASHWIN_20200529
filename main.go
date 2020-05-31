package main

import (
	"ASHWIN_20200529/config"
	"ASHWIN_20200529/services"
)

func main() {

	configuration := config.New("recipe.db")
	service := services.NewCrawlService(configuration)
	const URLToCrawl = "http://tech-challenge-golang.s3-website.eu-central-1.amazonaws.com/post"

	crawledInformation := service.Crawl(URLToCrawl)
	service.SaveAllTheRecipeInformation(crawledInformation)

}
