package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(errors.New("argument is not enough"))
		os.Exit(1)
	}
	var dlDir string
	var randomDelay bool
	var createDir bool
	var delay int
	var configFile string
	flag.StringVar(&dlDir, "dir", ".", "")
	flag.BoolVar(&randomDelay, "random-delay", false, "")
	flag.BoolVar(&createDir, "create-dir", false, "")
	flag.IntVar(&delay, "delay", 0, "Second")
	flag.StringVar(&configFile, "config", "", "Config file")
	flag.Parse()
	if len(strings.TrimSpace(configFile)) > 0 {
		fmt.Println("Read config")
	}
	dlDir = addTailSlash(dlDir)
	if randomDelay {
		rand.Seed(time.Now().UnixNano())
	}
	if exist, _ := isExist(dlDir); createDir && !exist {
		os.Mkdir(dlDir, 0775)
	}
	url := os.Args[len(os.Args)-1]
	if m, err := regexp.MatchString("https?://", url); !m {
		if err == nil {
			err = errors.New(url + ": is not url")
		}
		panic(err)
	}
	pageLink, err := PageLink(url)
	if err != nil {
		panic(err)
	}
	for _, link := range pageLink {
		fmt.Println(link)
		imageLink, err := ImageLink(link)
		if err != nil {
			panic(err)
		}
		for _, link := range imageLink {
			err := Download(dlDir, link)
			if err != nil {
				fmt.Println(err)
			} else {
				if randomDelay && delay > 0 {
					time.Sleep(time.Duration(rand.Intn(delay)) * time.Second)
				}
			}
		}
	}
}
