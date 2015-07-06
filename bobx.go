package main

import (
	"fmt"
	"github.com/rajninorp/bobx"
)

func main() {
	url := "http://www.bobx.com/av-idol/syoko-akiyama/series-syoko-akiyama-40-4-10.html"
	links, err := bobx.PageLink(url)
	if err != nil {
		panic(err)
	}
	for _, link := range links {
		fmt.Println(link)
	}
}
