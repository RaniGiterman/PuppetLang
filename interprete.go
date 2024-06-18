package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

type Line struct {
	ctx     context.Context // chromedp context to execute chromedp commands
	command []string        // string array of the written command (line split by spaces)
	num     int             // line number in the script
}

// recieves a line and interpretes it according to instructions.
func interprete(line Line) {
	switch line.command[0] {
	case "url":
		line.url()
	case "screenshot":
		line.screenshot()
	case "click":
		line.click()
	case "write":
		line.write()
	}
}

// goes to requested URL
// ex: url "https://rg-playground.vercel.app/"
func (line Line) url() {
	if err := safeLine(line, 1); err != nil {
		log.Fatal(err)
		return
	}

	address := strings.Trim(line.command[1], "\"")
	if err := chromedp.Run(line.ctx, chromedp.Navigate(address)); err != nil {
		log.Fatal(err)
		return
	}

	time.Sleep(time.Second) // wait a second for page to fully load
}

// takes screenshot and saves it to requested path
// ex: screenshot "a.png"
func (line Line) screenshot() {
	if err := safeLine(line, 1); err != nil {
		log.Fatal(err)
		return
	}

	var buf []byte
	path := strings.Trim(line.command[1], "\"")
	if err := chromedp.Run(line.ctx, chromedp.CaptureScreenshot(&buf)); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(path, buf, 0o644); err != nil {
		log.Fatal(err)
	}

}

// clicks an elements based on the selector recieved
// ex: click "#run_code"
func (line Line) click() {
	if err := safeLine(line, 1); err != nil {
		log.Fatal(err)
		return
	}

	selector := strings.Trim(line.command[1], "\"")
	if err := chromedp.Run(line.ctx, chromedp.Click(selector, chromedp.ByQuery)); err != nil {
		log.Fatal(err)
	}
}

// types the text into requested input/textarea field.
// ex: type "#editor" "הדפסה(1)"
func (line Line) write() {
	if err := safeLine(line, 2); err != nil {
		log.Fatal(err)
		return
	}

	selector := strings.Trim(line.command[1], "\"")
	value := strings.Trim(line.command[2], "\"")
	if err := chromedp.Run(line.ctx, chromedp.SetValue(selector, value, chromedp.ByQuery)); err != nil {
		log.Fatal(err)
	}
}
