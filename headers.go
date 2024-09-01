package scrapify

import (
	"github.com/gocolly/colly/v2"
	"net/http"
)

var defaultHeaders = map[string]string{
	"accept":                    "*/*",
	"accept-encoding":           "gzip, deflate, br",
	"accept-language":           "en-US,en;q=0.9",
	"cache-control":             "no-cache",
	"connection":                "keep-alive",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0",
	"upgrade-insecure-requests": "1",
	"sec-ch-ua":                 `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        "Windows",
	"sec-fetch-user":            "?1",
	"sec-fetch-dest":            "empty",
	"sec-fetch-mode":            "cors",
	"sec-fetch-site":            "same-origin",
	"priority":                  "u=0, i",
	"referer":                   "https://www.google.com/",
}

func SetHeaders(req *http.Request, headers M) {
	for k, v := range defaultHeaders {
		req.Header.Set(k, v)
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	req.Header.Set("Origin", req.URL.String())
	req.Header.Set("Host", req.URL.Host)
}

func SetCollyHeaders(req *colly.Request, headers M) {
	for k, v := range defaultHeaders {
		req.Headers.Set(k, v)
	}
	if headers != nil {
		for k, v := range headers {
			req.Headers.Set(k, v)
		}
	}
	req.Headers.Set("Origin", req.URL.String())
	req.Headers.Set("Host", req.URL.Host)
}
