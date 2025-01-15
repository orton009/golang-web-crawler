package main

import (
	"reflect"
	"testing"
)

func TestCollectURLs(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			urls, err := getURLsFromHTML(test.inputBody, test.inputURL)
			if err != nil {
				t.Errorf("test: %v unexpected error %v", test.name, err)
			}

			for idx := range urls {
				if !reflect.DeepEqual(urls[idx], test.expected[idx]) {
					t.Errorf("test: %v failed matching url, expected: %v actual: %v", test.name, test.expected[idx], urls[idx])
				}
			}
		})

	}

}
