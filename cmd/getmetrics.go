package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MiLk/kingpin"
)

var (
	prefix string
	url    string
)

func getMetrics(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Couldn't get url: %s", url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// split into lines
	lines := strings.Split(string(body), "\n")
	for x := range lines {
		fmt.Printf("line: %s\n", lines[x])
	}
}

func main() {
	kingpin.Flag("prefix", "prefix to use for metrics").Short('s').StringVar(&prefix)
	kingpin.Flag("url", "url to get metrics from").Short('u').StringVar(&url)
	kingpin.CommandLine.HelpFlag.Hidden()
	kingpin.Parse()

	getMetrics(url)
}
