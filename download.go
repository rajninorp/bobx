package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func Download(dir string, url string) error {
	ref, fname := path.Split(url)
	fname = dir + fname
	exist, err := isExist(fname)
	if err != nil {
		return err
	} else if exist {
		return errors.New(fname + ": is exist")
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Referer", ref)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(body)
	return nil
}
