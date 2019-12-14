package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/MiLk/kingpin"
)

var (
	prefix string
	url    string
)

func getMetrics(url string, prefix string) {
	date := int32(time.Now().Unix())
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Couldn't get url: %s", url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// split into lines
	valid := regexp.MustCompile(`kube_metrics`)
	if !valid.MatchString(string(body)) {
		fmt.Println("couldn't find valid metrics!")
		return
	} else {
		lines := strings.Split(string(body), "\n")
		for x := range lines {
			lines[x] = strings.Replace(lines[x], "{", ",", -1)
			lines[x] = strings.Replace(lines[x], "}", "", -1)
			parts := strings.Split(lines[x], " ")
			if len(parts) >= 2 {
				fmt.Printf("%s%s value=%s %d\n", prefix, parts[0], parts[1], date)
			}
		}
		return
	}
}

func main() {
	kingpin.Flag("prefix", "prefix to use for metrics").Short('s').StringVar(&prefix)
	kingpin.Flag("url", "url to get metrics from").Short('u').StringVar(&url)
	kingpin.CommandLine.HelpFlag.Hidden()
	kingpin.Parse()

	getMetrics(url, prefix)
}
