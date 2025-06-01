package scrape

import (
	"database/sql"
	"embeddings/util"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	
)

func cleanHtml(html string) []string {
extractedData := []string{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("error loading html", err)
	}

		doc.Find("body").Each(func(i int, p *goquery.Selection) {

			doc.Find(".paywall").Each(func(i int, p *goquery.Selection) {
				title := p.Filter("h2").Text()
				description := p.Filter("p").First().Text()
				extractedData = append(extractedData, fmt.Sprintf("Title: %s, Description: %s,", title, description))
			})
		})
		fmt.Println(extractedData)
	return extractedData
}

func fetchHtml(url string)(string, error) {
	resp, err:= http.Get(url)
	if err != nil {
		return " ", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(" failed to fetch %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return " ", err
	}
	// fmt.Println(string(body))
	return string(body), nil
}
// decprecated this function scrapes the title and the descr

func scrapeTitles(html string) []string {
	extractedData := []string{}
	
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			log.Println("error loading html", err)
		}
	
			doc.Find("body").Each(func(i int, p *goquery.Selection) {
	
				doc.Find(".paywall").Each(func(i int, p *goquery.Selection) {

					title := p.Filter("h2").Text()
					if title != "" {
						extractedData = append(extractedData, title)
					}
				})
			})
		return extractedData
	}



func ScrapeNetflix(db *sql.DB) {
	env := util.LoadEnviroment()

	URL := env[5]
	html, err := fetchHtml(URL)
	if err != nil {
		log.Fatal("Failed to fetch HTML:", err)
	}
	
	scrapedTitles := scrapeTitles(html)
	contentArr, err := ScrapeToContent(scrapedTitles, "Netflix")
	if err != nil {		
		log.Fatal("error in getting content", err)
	}
	util.WriteJSONToFile("scrape",contentArr)
	err = contentToTurso(db,contentArr)
	if err != nil {
		log.Fatal("err in placing in db", err)
	}
}


func ScrapeHbo(db *sql.DB){
	env := util.LoadEnviroment()
	URL := env[4]
	html, err := fetchHtml(URL)
	if err != nil {
		log.Fatal("Failed to fetch HTML:", err)
	}
	
	scrapedTitles := scrapeTitles(html)
	contentArr, err := ScrapeToContent(scrapedTitles, "HBO Max")
	if err != nil {		
		log.Fatal("error in getting content", err)
	}
	util.WriteJSONToFile("scrape",contentArr)
	err = contentToTurso(db,contentArr)
	if err != nil {
		log.Fatal("err in placing in db", err)
	}
}
func ScrapeAppleTV(db *sql.DB){
	env := util.LoadEnviroment()
	URL := env[7]
	html, err := fetchHtml(URL)
	if err != nil {
		log.Fatal("Failed to fetch HTML:", err)
	}
	
	scrapedTitles := scrapeTitles(html)
	contentArr, err := ScrapeToContent(scrapedTitles, "Apple TV+")
	if err != nil {		
		log.Fatal("error in getting content", err)
	}
	util.WriteJSONToFile("scrape",contentArr)
	err = contentToTurso(db,contentArr)
	if err != nil {
		log.Fatal("err in placing in db", err)
	}
}

func ScrapeAmazon(db *sql.DB){
	env := util.LoadEnviroment()
	URL := env[8]
	html, err := fetchHtml(URL)
	if err != nil {
		log.Fatal("Failed to fetch HTML:", err)
	}
	
	scrapedTitles := scrapeTitles(html)
	contentArr, err := ScrapeToContent(scrapedTitles, "Amazon Prime")
	if err != nil {		
		log.Fatal("error in getting content", err)
	}
	util.WriteJSONToFile("scrape",contentArr)
	err = contentToTurso(db,contentArr)
	if err != nil {
		log.Fatal("err in placing in db", err)
	}
}
func ScrapeDisneyPlus(db *sql.DB){
	env := util.LoadEnviroment()
	URL := env[6]
	html, err := fetchHtml(URL)
	if err != nil {
		log.Fatal("Failed to fetch HTML:", err)
	}
	
	scrapedTitles := scrapeTitles(html)
	contentArr, err := ScrapeToContent(scrapedTitles, "Disney+")
	if err != nil {		
		log.Fatal("error in getting content", err)
	}
	util.WriteJSONToFile("scrape",contentArr)
	err = contentToTurso(db,contentArr)
	if err != nil {
		log.Fatal("err in placing in db", err)
	}
}

func ScrapeAll(db *sql.DB){
	// the apple tv link is for movies!
	ScrapeAppleTV(db)
	// ScrapeAmazon(db)
	// ScrapeHbo(db)
	// ScrapeNetflix(db)
	// ScrapeDisneyPlus(db)
}