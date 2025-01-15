package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return "", errors.New("expected html content")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
