package main

import (
	"strings"
	"os"
)

func addTailSlash(s string) string {
	return strings.TrimRight(s, "/") + "/"
}

func isExist(s string) (bool, error) {
	_, err := os.Stat(s)
	return !os.IsNotExist(err), err
}
