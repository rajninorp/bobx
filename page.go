package main

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

var pageLink []string

func TopPage(url string) (string, error) {
	reg := regexp.MustCompile("(.*)-[0-9]+-([0-9]+-[0-9]+.html)")
	group := reg.FindSubmatch([]byte(url))
	if len(group) < 3 {
		return "", errors.New(url + ": parse error")
	}
	return string(group[1]) + "-0-" + string(group[2]), nil
}

func pageCrawler(url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	doc.Find("link").Each(func(_ int, s *goquery.Selection) {
		rel, exist := s.Attr("rel")
		if exist && rel == "next" {
			link, _ := s.Attr("href")
			pageLink = append(pageLink, link)
			pageCrawler(link)
		}
	})
	return nil

}

func PageLink(url string) ([]string, error) {
	topPage, err := TopPage(url)
	if err != nil {
		panic(err)
	}
	pageLink = append(pageLink, topPage)
	err = pageCrawler(url)
	if err != nil {
		panic(err)
	}
	return pageLink, nil
}
