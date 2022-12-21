package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/alexflint/go-arg"
)

type config struct {
	Filename      string `arg:"positional,required,help:The name of the text file to save the table to"`
	HumanReadable bool   `arg:"-H,help:Print the time in human-readable format"`
}

func main() {
	var c config
	arg.MustParse(&c)

	file, err := os.Create(c.Filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		currentTime := time.Now()
		var timeString string
		if c.HumanReadable {
			timeString = currentTime.Format("2006-01-02 15:04:05")
		} else {
			timeString = fmt.Sprintf("%d", currentTime.UnixNano()/int64(time.Millisecond))
		}
		_, err := fmt.Fprintf(w, "%s: %s\n", timeString, row)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("%s: %s\n", timeString, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
