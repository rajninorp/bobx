package main

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

var imageLink []string

func ImageLink(url string) ([]string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		var empty []string
		return empty, err
	}
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		link, exist := s.Attr("src")
		if exist {
			if m, _ := regexp.MatchString(`/thumbnail.*-preview-.*`, link); m {
				link = strings.Replace(link, "/thumbnail", "http://www.bobx.com", 1)
				link = strings.Replace(link, "preview-", "", 1)
				imageLink = append(imageLink, link)
			}
		}
	})
	return imageLink, nil

}
