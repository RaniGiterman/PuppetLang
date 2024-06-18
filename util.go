package main

import (
	"fmt"
	"strconv"
	"time"
)

// returns error if the line is invalid
// should recieve the Line, and expected number of arguments
func safeLine(line Line, argueCount int) error {
	// command's length should be {argueCount} + 1 for the actual command word (write, click..)
	if len(line.command) != argueCount+1 {
		return fmt.Errorf("line: %d, expected 1 argument to `%s` command but found: %d", line.num, line.command[0], len(line.command)-1)
	}

	for i := 1; i < argueCount; i++ {
		argument := line.command[i]
		// argument should start and end with quotes, eg: "https://google.com"
		if argument[0] != '"' || argument[len(argument)-1] != '"' {
			return fmt.Errorf(fmt.Sprintf("line: %d, syntax error: argument should be wrapped with quotations", line.num))
		}
	}

	return nil
}

// recieves a string describing duration in seconds, minutes or hours, and returns that duration in time.Duration
// eg: 30s, 45m, 3h
func strToTime(str string) (time.Duration, error) {
	t := str[len(str)-1]
	num, err := strconv.Atoi(str[:len(str)-1])
	if err != nil {
		return time.Duration(0), err // invalid string
	}

	var d time.Duration
	switch t {
	case 's':
		d = time.Second * time.Duration(num)
	case 'm':
		d = time.Minute * time.Duration(num)
	case 'h':
		d = time.Hour * time.Duration(num)
	}

	return d, nil
}
