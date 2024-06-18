package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	path := flag.String("f", "script.pu", "file path to execute")                                                           // f for file
	timeout := flag.String("t", "30s", "time until script times out, s - seconds, m - minutes, h - hour, default is never") // t for timeout
	flag.Parse()

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var d time.Duration
	if *timeout == "" {
		d = time.Hour * time.Duration(24*100000) // ? 273.. years. If user didn't enter a timeout, the script will never timeout.
	} else {
		d, err = strToTime(*timeout)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	// create chromedp context
	ctx, cancel := chromedp.NewContext(
		ctxTimeout,
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	scanner := bufio.NewScanner(file)
	num := 1
	for scanner.Scan() {
		interprete(Line{ctx: ctx, command: strings.Split(scanner.Text(), " "), num: num})
		num++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
