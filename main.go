package main

import (
	"os"
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

func main() {
	dat, _ := os.ReadFile("src/from_file_if_exists.sh")
	r := strings.NewReader(string(dat))
	f, err := syntax.NewParser().Parse(r, "")
	if err != nil {
		return
	}
	syntax.NewPrinter().Print(os.Stdout, f)
}