package services

import (
	"ASHWIN_20200529/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

var configurationTest = config.New("recipe_test.db")

func TestCrawlService_Crawl(t *testing.T) {

	crawlServiceTest := NewCrawlService(configurationTest)
	results := crawlServiceTest.Crawl("http://tech-challenge-golang.s3-website.eu-central-1.amazonaws.com/post")
	var buf = new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.Encode(results)

	f, err := os.Create("recipe_test.json")
	if err != nil {
		log.Fatal(f)
	}
	defer f.Close()
	io.Copy(f, buf)
}
func TestCrawlService_SaveAllTheRecipeInformation(t *testing.T) {
	crawlServiceTest := NewCrawlService(configurationTest)
	var recipeDetails []recipeInformation
	f, err := os.Open("recipe_test.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(f)
	decoder.Decode(&recipeDetails)

	crawlServiceTest.SaveAllTheRecipeInformation(recipeDetails)
}

func TestCrawlService_RetrieveRecipeInformation(t *testing.T) {
	db := configurationTest.DatabaseInfo
	rs, err := db.Query("SELECT  * FROM recipe_information")
	if err != nil {
		log.Fatal(err)
	}
	var id int
	var desc string
	var ingredeientList string
	var preparationSteps string
	var categories string
	var tags string

	for rs.Next() {
		rs.Scan(&id, &desc, &ingredeientList, &preparationSteps, &categories, &tags)
		break
	}
	fmt.Println(id, desc, ingredeientList, preparationSteps, categories, tags)

}
