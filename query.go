package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pejman-hkh/gdp/gdp"
)

type Weblink struct {
	title string
	description string
	link string
}

func (l Weblink) Title() string { return l.title; }
func (l Weblink) Description() string { return l.description; }
func (l Weblink) FilterValue() string { return l.title + " " + l.description; }
func (l Weblink) Link() string { return l.link; }

func Query(query string) ([]Weblink, error) {
	query = url.QueryEscape(query)
	request, err := http.NewRequest(http.MethodGet, "https://html.duckduckgo.com/html?q=" + query, nil)
	if err != nil {
		return nil, fmt.Errorf("while creating request: %+v", err)
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:139.0) Gecko/20100101 Firefox/139.0")

	response, err := http.Get("https://html.duckduckgo.com/html?q=" + query)
	if err != nil {
		return nil, fmt.Errorf("while sending duckduckgo request: %+v", err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("while reading the response from duckduckgo: %+v", err)
	}

	document := gdp.Default(string(content))

	titles := document.Find(".result__a")
	descriptions := document.Find(".result__snippet")
	urls := document.Find(".result__url")

	links := make([]Weblink, 0, titles.Length())
	titles.Each(func(index int, tag *gdp.Tag) {
		links = append(links, Weblink {
			title: strings.TrimSpace(tag.Text()),
			description: strings.TrimSpace(descriptions.Eq(index).Text()),
			link: strings.TrimSpace(urls.Eq(index).Text()),
		})
	})

	if len(links) == 0 {
		fmt.Println(string(content))
	}

	return links, nil
}
