package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		exitWithError(errors.New("no website provided"))
	}

	if len(os.Args) > 4 {
		exitWithError(errors.New("too many arguments provided"))
	}

	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		exitWithError(err)
	}

	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		exitWithError(err)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		exitWithError(fmt.Errorf("error - configure %w", err))
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizeURL, count := range cfg.pages {
		fmt.Printf("%d - %s \n", count, normalizeURL)
	}

}
