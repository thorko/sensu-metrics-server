package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
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
		namespaceCPU := make(map[string]int)
		namespaceMem := make(map[string]int)
		comment := regexp.MustCompile(`^#`)
		lines := strings.Split(string(body), "\n")
		for x := range lines {
			if !comment.MatchString(lines[x]) {
				lines[x] = strings.Replace(lines[x], "{", ",", -1)
				lines[x] = strings.Replace(lines[x], "}", "", -1)
				parts := strings.Split(lines[x], " ")
				if len(parts) >= 2 {
					re1 := regexp.MustCompile("pod_namespace=\"(.*)\"")
					nn := re1.FindStringSubmatch(parts[0])
					re2 := regexp.MustCompile("kube_metrics_server_pods_(cpu|mem)")
					typ := re2.FindStringSubmatch(parts[0])

					f, _ := strconv.ParseFloat(parts[1], 64)

					if len(nn) > 0 && len(typ) > 0 {

						if typ[1] == "cpu" {
							namespaceCPU[nn[1]] += int(f)
						}
						if typ[1] == "mem" {
							namespaceMem[nn[1]] += int(f)
						}
					}

					fmt.Printf("%s%s value=%d %d\n", prefix, parts[0], int(f), date)
				}
			}
		}

		for x := range namespaceCPU {
			fmt.Printf("%skube_state_metrics_namespace_cpu,namespace=%s value=%d %d\n", prefix, x, namespaceCPU[x], date)
		}
		for x := range namespaceMem {
			fmt.Printf("%skube_state_metrics_namespace_mem,namespace=%s value=%d %d\n", prefix, x, namespaceMem[x], date)
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
