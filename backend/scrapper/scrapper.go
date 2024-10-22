package scrapper
import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"anitsayac/database"

	"github.com/gocolly/colly/v2"
)


var (
	baseUrl = "https://anitsayac.com"
	date_layout = "02/01/2006"
)

func ReplaceAll(s, old, new string, n int) string {
	re := regexp.MustCompile(old)
	return re.ReplaceAllString(s, new)
}

func getArticleContent(url string) database.Detail {
	detail := database.Detail{}
	c := colly.NewCollector(
		colly.AllowedDomains("anitsayac.com"),
		colly.CacheDir("./anitsayac_cache"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Detail: ", r.URL.String())
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		nameMatches := regexp.MustCompile(`(?i)<b>Ad Soyad:\s*</b>\s*(.+?)<br>`).FindStringSubmatch(string(e.Response.Body))
		if len(nameMatches) > 1 {
			detail.Name = nameMatches[1]
		}

		ageMatches := regexp.MustCompile(`(?i)<b>Maktülün yaşı:\s*</b>\s*(.+?)<br>`).FindStringSubmatch(string(e.Response.Body))
		if len(ageMatches) > 1 {
			detail.Age = ageMatches[1]
		}

		locationMatches := regexp.MustCompile(`(?i)<b>İl/ilçe:\s*</b>\s*(.+?)<br>`).FindStringSubmatch(string(e.Response.Body))
		if len(locationMatches) > 1 {
			detail.Location = locationMatches[1]
		}

		dateMatches := regexp.MustCompile(`(?i)<b>Tarih:\s*</b>\s*(.+?)<br>`).FindStringSubmatch(string(e.Response.Body))
		if len(dateMatches) > 1 {
			detail.Date = dateMatches[1]
		}

		reasonMatches := regexp.MustCompile(`(?i)<b>Neden öldürüldü:\s*</b>\s*(.+?)<br>`).FindStringSubmatch(string(e.Response.Body))
		if len(reasonMatches) > 1 {
			detail.Reason = reasonMatches[1]
		}

		byMatches := regexp.MustCompile(`(?i)<b>Kim tarafından öldürüldü:\s*</b>\s*(.+?)<br>`).FindStringSubmatch(string(e.Response.Body))
		if len(byMatches) > 1 {
			detail.By = byMatches[1]
		}

		protectionMatches := regexp.MustCompile(`(?i)<b>Korunma talebi:\s*</b>\s*(.+?)<br>`).FindStringSubmatch(string(e.Response.Body))
		if len(protectionMatches) > 1 {
			detail.Protection = protectionMatches[1]
		}

		methodMatches := regexp.MustCompile(`(?i)<b>Öldürülme şekli:\s*</b>\s*(.+?)<br>`).FindStringSubmatch(string(e.Response.Body))
		if len(methodMatches) > 1 {
			detail.Method = methodMatches[1]
		}

		statusMatches := regexp.MustCompile(`(?i)<b>Failin durumu:\s*</b>\s*(.+?)<br>`).FindStringSubmatch(string(e.Response.Body))
		if len(statusMatches) > 1 {
			detail.Status = statusMatches[1]
		}

		sources := e.ChildAttrs("a", "href")
		detail.Source = sources

		detail.Image = baseUrl + "/" + e.ChildAttr("img", "src")
	})

	c.Visit(url)

	return detail
}

func ScrapeData() {
	c := colly.NewCollector(
		colly.AllowedDomains("anitsayac.com"),
		colly.CacheDir("./anitsayac_cache"),
	)

	incidents := make([]database.Incident, 0, 20000)

	c.OnHTML("div#divcounter", func(e *colly.HTMLElement) {
		e.ForEach("span.xxy", func(i int, e *colly.HTMLElement) {
			detail := getArticleContent(baseUrl + "/" + e.ChildAttr("span.xxy > a", "href"))
			if(len(detail.Name) > 1 && len(detail.Date) > 1){

				year, _ := time.Parse(date_layout,detail.Date )
				fmt.Println(detail.Date)
				incident := database.Incident{
					Id: func() int {
						id, _ := strconv.Atoi(strings.Split(e.ChildAttr("span.xxy > a", "href"), "=")[1])
						return id
					}(),
					Name:       e.ChildText("span.xxy > a"),
					FullName:   detail.Name,
					Age:        detail.Age,
					Location:   detail.Location,
					Date:       detail.Date,
					Year: 		strconv.Itoa(year.Year()),
					Reason:     detail.Reason,
					By:         detail.By,
					Protection: detail.Protection,
					Method:     detail.Method,
					Status:     detail.Status,
					Source:     detail.Source,
					Image:      detail.Image,
					Url:        baseUrl + "/" + e.ChildAttr("span.xxy > a", "href"),
				}
				database.InsertIncident(incident)
				incidents = append(incidents, incident)
			}
		})
	})

	c.Visit(baseUrl + "/?year=2000")

}