# Scrapify🚀: A Go Library for Web Scraping

This library provides tools for building web scrapers in Go. It allows you to create custom HTTP requests with advanced features for bypassing basic anti-scraping measures.

## Features

- **Customizable Headers**: Set various headers like User-Agent and sec-ch-ua to mimic a real browser.
- **TLS Configuration**: Customize the Transport Layer Security (TLS) configuration for secure connections.
- **Browser Emulation**: Specify different browsers (Chrome, Firefox, Edge) to influence the cipher suites offered.
- **Default Headers**: Includes a set of common headers to improve compatibility.

## Installation
```bash
  go get -u github.com/iarsham/scrapify@latest
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/iarsham/scrapify"
	"net/http"
)

func main() {
	client := &http.Client{
		Transport: scrapify.NewTransport(scrapify.Chrome),
	}
	req, err := http.NewRequest(http.MethodGet, "https://cloudflare.com", nil)
	if err != nil {
		panic(err)
	}
	scrapify.SetHeaders(req, nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
```

```go
package main

import (
	"fmt"
	"github.com/iarsham/scrapify"
)

func main() {
	c := colly.NewCollector()
	c.WithTransport(scrapify.NewTransport(scrapify.Chrome))

	c.OnRequest(func(r *colly.Request) {
		scrapify.SetCollyHeaders(r, nil)
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	if err := c.Visit("https://chatgpt.com"); err != nil {
		panic(err)
	}
}

```
## Contributing

We welcome contributions to this library. Please feel free to submit pull requests with improvements and bug fixes.

## License

This library is licensed under the MIT License. See the LICENSE file for details.