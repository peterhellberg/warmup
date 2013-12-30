// Copyright 2013 Peter Hellberg. All rights reserved.
// Licensed under the MIT License found in the LICENSE file.

/*
HTTP cache warming.

Installation

Go 1.1+ is required since bufio.NewScanner is used.

		go get -u github.com/peterhellberg/warmup

Usage of warmup

Note: You probably want to warm the (soon to be) hottest routes on your site.

Command line flags:

		-base-url="http://0.0.0.0:7000": The base URL used for paths
		-color=true: If you want the output in color
		-delay=100: Delay (in ms) between requests
		-fatal-errors=false: Useful for automated testing
		-filename="urls.txt": List of URLs
		-limit=100: Limit of concurrent requests

*/
package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	filename    = flag.String("filename", "urls.txt", "List of URLs")
	showColor   = flag.Bool("color", true, "If you want the output in color")
	fatalErrors = flag.Bool("fatal-errors", false, "Useful for automated testing")
	limit       = flag.Int("limit", 100, "Limit of concurrent requests")
	delay       = flag.Int("delay", 100, "Delay (in ms) between requests")
	baseURL     = flag.String("base-url", "http://0.0.0.0:7000", "The base URL used for paths")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	urls, err := readURLs(filename)

	if err != nil {
		fatal(err.Error())
	}

	numURLs := len(urls)

	log.Print("Launching " +
		green(strconv.Itoa(numURLs)) +
		green(" GET") + " requests.")

	ch := make(chan string, *limit)

	for i, url := range urls {
		go get(url, i, &ch)
	}

	for i := 0; i < numURLs; i++ {
		select {
		case r := <-ch:
			log.Print(r)
		}
	}

	log.Print(gray("DONE."))
}

func fatal(err string) {
	log.Fatalln(fatalMessage(err))
}

func fatalMessage(err string) string {
	return red("ERR") + " " + err
}

func get(url string, i int, ch *chan string) {
	if *delay > 0 {
		time.Sleep(time.Duration(*delay*i) * time.Millisecond)
	}

	log.Print(gray("REQ " + url))

	resp, err := http.Get(url)

	if err != nil {
		if *fatalErrors {
			fatal(err.Error())
		} else {
			*ch <- red("ERR ") + err.Error()
		}
	} else {
		defer resp.Body.Close()

		if resp.StatusCode < 400 {
			*ch <- green(resp.Status) + " " + url
		} else {
			if *fatalErrors {
				fatal(resp.Status + " " + blue(url))
			} else {
				*ch <- red(resp.Status) + " " + blue(url)
			}
		}
	}
}

func readURLs(urls *string) ([]string, error) {
	file, err := os.Open(*urls)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "http") {
			lines = append(lines, scanner.Text())
		}

		if *baseURL != "" && strings.HasPrefix(scanner.Text(), "/") {
			lines = append(lines, *baseURL+scanner.Text())
		}
	}

	return lines, scanner.Err()
}

func color(c, s string) string {
	if *showColor {
		return "\033[" + c + "m" + s + "\033[0m"
	} else {
		return s
	}
}

func red(s string) string {
	return color("0;31", s)
}

func green(s string) string {
	return color("0;32", s)
}

func blue(s string) string {
	return color("0;34", s)
}

func gray(s string) string {
	return color("1;30", s)
}
