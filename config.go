package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	baseURL            *url.URL
	pages              map[string]int
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func configure(rawBaseURL string, maxConcurrency int, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}

func (c *config) addPageVisit(normalizedURL string) (isFirst bool) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if _, visited := c.pages[normalizedURL]; visited {
		c.pages[normalizedURL]++
		return false
	}

	c.pages[normalizedURL] = 1
	return true
}

func (c *config) checkPageLimitExceeded() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.pages) >= c.maxPages {
		return true
	}
	return false
}

func (c *config) crawlPage(rawCurrentURL string) {
	c.concurrencyControl <- struct{}{}
	defer func() {
		<-c.concurrencyControl
		c.wg.Done()
	}()

	if c.checkPageLimitExceeded() {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error - crawlpage: couldn't parse url %s \n %v", rawCurrentURL, err)
	}

	if c.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error - normalizedURL : %v ", err)
		return
	}

	isFirst := c.addPageVisit(normalizedURL)
	if !isFirst {
		// refrain from calling same URL twice
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := fetchURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error - getHTML: %v", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, rawCurrentURL)
	if err != nil {
		fmt.Printf("error - getURLSFromHTML: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		c.wg.Add(1)
		go c.crawlPage(nextURL)
	}

}
