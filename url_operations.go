package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func normalizeURL(originalURL string) (string, error) {
	url, err := url.Parse(originalURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := url.Host + url.Path

	return strings.TrimSuffix(strings.ToLower(fullPath), "/"), nil
}

func traverseNodesAndCollectURLs(node *html.Node) []string {
	collectedURLs := []string{}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				collectedURLs = append(collectedURLs, attr.Val)
			}
		}
	}

	for childNode := range node.ChildNodes() {
		collectedURLs = append(collectedURLs, traverseNodesAndCollectURLs(childNode)...)
	}

	return collectedURLs
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	urls := make([]string, 0)

	reader := strings.NewReader(htmlBody)

	htmlNode, err := html.Parse(reader)
	if err != nil {
		return urls, err
	}

	urls = append(urls, traverseNodesAndCollectURLs(htmlNode)...)
	// check for relative and absolute urls
	// TODO:

	for i := range urls {
		if strings.HasPrefix(urls[i], "/") {
			fullPath, err := url.JoinPath(rawBaseURL, urls[i])
			if err != nil {
				return urls, err
			}

			urls[i] = fullPath
		}
	}
	return urls, nil

}
