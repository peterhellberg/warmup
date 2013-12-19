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
)

var (
	filename  = flag.String("filename", "urls.txt", "List of URLs")
	showColor = flag.Bool("color", true, "If you want the output in color")
	limit     = flag.Int("limit", 100, "Limit of concurrent requests")
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

	for _, url := range urls {
		go get(url, &ch)
	}

	for i := 0; i < numURLs; i++ {
		select {
		case r := <-ch:
			log.Print(r)
		}
	}
}

func fatal(err string) {
	log.Fatalln(fatalMessage(err))
}

func fatalMessage(err string) string {
	return red("ERROR:") + " " + err
}

func get(url string, ch *chan string) {
	log.Print(gray("REQ " + url))

	resp, err := http.Get(url)

	if err != nil {
		*ch <- red("ERR ") + err.Error()
	}

	defer resp.Body.Close()

	if resp.StatusCode < 400 {
		*ch <- green(resp.Status) + " " + url
	} else {
		*ch <- red(resp.Status) + " " + blue(url)
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
		if strings.Contains(scanner.Text(), "http") {
			lines = append(lines, scanner.Text())
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
