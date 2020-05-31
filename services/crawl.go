package services

import (
	"ASHWIN_20200529/config"
	"ASHWIN_20200529/util"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"os"
	"strings"
)

type CrawlService interface {
	Crawl(string) []recipeInformation
	SaveAllTheRecipeInformation([]recipeInformation)
}

func NewCrawlService(configuration *config.Configuration) CrawlService {
	return &crawlService{
		configuration: configuration,
	}

}

type (
	crawlService struct {
		configuration *config.Configuration
	}
)

type recipeInformation struct {
	Name             string
	Description      string
	IngredientsList  []string
	PreparationSteps []string
	Categories       []string
	Tags             []string
}

func (instance *crawlService) Crawl(urlToCrawl string) []recipeInformation {
	response, err := http.Get(urlToCrawl)
	checkError(err)
	defer response.Body.Close()

	tokenizer := html.NewTokenizer(response.Body)

	rpInformationList := make([]recipeInformation, 12)

	index := 0
	recipeChannelDetails := make(chan recipeInformation, 12)

	for {
		nextToken := tokenizer.Next()
		switch nextToken {
		case html.StartTagToken:
			value := tokenizer.Token()
			if value.Data == "a" {
				for _, a := range value.Attr {
					if a.Key == "href" && !util.Contains(a.Val) {
						go crawlRecipeDetails(a.Val, recipeChannelDetails)
						break
					}
				}
			}

		case html.ErrorToken:
			for {
				result := <-recipeChannelDetails
				rpInformationList[index] = result
				index++
				if len(rpInformationList) == index {
					break
				}
			}
			close(recipeChannelDetails)
			return rpInformationList
		}

	}

}
func crawlRecipeDetails(url string, rpDetails chan recipeInformation) {
	response, err := http.Get(url)

	checkError(err)
	defer response.Body.Close()

	tokenizer := html.NewTokenizer(response.Body)
	countMap := make(map[string]int)
	countMap["P"] = 0
	values := map[string][]string{}
	values["Ingredients"] = []string{}
	values["Steps"] = []string{}
	values["Categories"] = []string{}
	values["Tags"] = []string{}
	var rpInfo recipeInformation
	for {
		nextToken := tokenizer.Next()
		switch nextToken {
		case html.StartTagToken:
			value := tokenizer.Token()
			switch value.DataAtom {
			case atom.H3:
				tokenizer.Next()
				data := tokenizer.Token().Data
				if value.Attr == nil {
					rpInfo.Name = data
				}
				if data == "Steps" {
					rpInfo.IngredientsList = util.RemoveValuFromIndex(values["Ingredients"], 0)
				}

			case atom.P:
				if countMap["P"] == 0 {
					countMap["P"]++
					tokenizer.Next()
					rpInfo.Description = tokenizer.Token().Data
				}

			case atom.Li:
				tokenizer.Next()
				if len(rpInfo.IngredientsList) == 0 {
					appendedValues := append(values["Ingredients"], tokenizer.Token().Data)
					values["Ingredients"] = appendedValues
				} else if len(rpInfo.PreparationSteps) == 0 {
					appendedValues := append(values["Steps"], tokenizer.Token().Data)
					values["Steps"] = appendedValues
				} else if len(rpInfo.Categories) == 0 {
					tokenizer.Next()
					appendedValues := append(values["Categories"], tokenizer.Token().Data)
					values["Categories"] = appendedValues
				}
			case atom.A:
				tokenizer.Next()
				if len(rpInfo.Tags) == 0 {
					appendedValues := append(values["Tags"], tokenizer.Token().Data)
					values["Tags"] = appendedValues
				}

			case atom.Strong:
				tokenizer.Next()
				data := tokenizer.Token().Data
				if data == "Categories" {
					rpInfo.PreparationSteps = values["Steps"]
				} else if data == "Tags" {
					rpInfo.Categories = values["Categories"]
				}

			}
		case html.ErrorToken:
			rpInfo.Tags = util.RemoveValuFromIndex(values["Tags"], 0)
			rpDetails <- rpInfo
			return
		}
	}

}

func (instance *crawlService) SaveAllTheRecipeInformation(information []recipeInformation) {
	database := instance.configuration.DatabaseInfo
	for _, value := range information {
		stmt, err := database.Prepare("INSERT INTO recipe_information(recipe_name,desc,ingredients,preparation_steps,categories,tags) values (?,?,?,?,?,?)")
		checkError(err)
		result, err := stmt.Exec(value.Name, value.Description, strings.Join(value.IngredientsList, ","), strings.Join(value.PreparationSteps, ","), strings.Join(value.Categories, ","), strings.Join(value.Tags, ","))
		checkError(err)
		id, err := result.LastInsertId()
		fmt.Println("id=", id)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
