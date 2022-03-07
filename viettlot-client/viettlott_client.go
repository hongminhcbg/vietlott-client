package viettlot_client

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const KenoLiveUrl = "https://vietlott.tv/truc-tiep-keno.html"

type ResultKeno struct {
	Name string
	Date string
	Time string
	LineFirst []string
	LineSecond []string
}

type Service struct {

}

func NewVietlottClient() *Service {
	return &Service{}
}

func (s *Service) KenoLive(ctx context.Context) ([]*ResultKeno, error) {
	resp, err := http.Get(KenoLiveUrl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code fail, %d", resp.StatusCode)
	}

	result := make([]*ResultKeno, 0)
	var element *ResultKeno
	lineFirst := make([]string, 0, 14)
	lineSecond := make([]string, 0, 14)


	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".w50pt #kq").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tbody tr").Each(func(i int, selection *goquery.Selection) {
				if i % 2 == 0 {
					//init
					element = &ResultKeno{
						Name:       selection.Text(),
						LineFirst: nil,
						LineSecond: nil,
					}
					lineFirst = []string{}
					lineSecond = []string{}
					selection.Find("td").Each(func(i int, selection *goquery.Selection) {
						switch i {
						case 0:
							element.Name = selection.Text()
						case 1:
							selection.Find("div").Each(func(i int, selection *goquery.Selection) {
								if i == 0 {
									element.Date = selection.Text()
									return
								}

								element.Time = selection.Text()
							})
						default:
							lineFirst = append(lineFirst, selection.Text())
						}
					})

					return
				}

				selection.Find("td").Each(func(i int, selection *goquery.Selection) {
					lineSecond = append(lineSecond, selection.Text())
				})

				element.LineFirst = append(element.LineFirst, lineFirst...)
				element.LineSecond = append(element.LineSecond, lineSecond...)
				result = append(result, element)
			})
	})

	return result, nil
}


