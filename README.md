# Golang Web Crawler

## Project Description

This project is a CLI application written in Go that acts as a web crawler to generate an "internal links" report for any website. Internal linking is an important aspect of Search Engine Optimization (SEO), as it helps websites rank better in search engines like Google. By crawling each page of a site, the tool identifies and reports all internal links, providing valuable insights for improving website structure.

## Features

- Crawl Websites: Traverse and analyze internal links within any given website.

- Generate Reports: Output a report containing all internal links found during the crawl.

- Customizable Parameters: Specify depth and timeout for the crawling process.

## Installation

Ensure you have Go installed on your system. If not, download and install it from Go's official website.

- Clone this repository:

```
git clone https://github.com/orton009/golang-web-crawler.git
cd golang-web-crawler
```

- Build the CLI application:

```go build -o webcrawler```

## Usage

Run the CLI tool using the following command:

```
./webcrawler <website_url> <max_concurrent_threads> <max_pages_to_fetch>

./webcrawler "https://crawler-test.com/" 3 25
```

## Output

The tool generates a report listing all the internal links discovered on the crawled website. The output is displayed in the terminal and can be redirected to a file if needed:

```./webcrawler https://crawler-test.com/ 3 25 > internal_links_report.txt```

## Development

To run the application without building an executable:

```go run main.go <website_url> <max_depth> <timeout>```

## Contributions

Contributions are welcome! Feel free to fork the repository, submit pull requests, or file issues for bugs or feature requests.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

Inspired by the need to enhance website SEO through better internal linking.

Thanks to https://crawler-test.com/ for providing a testable domain for this project.

