package scrapify

import "net/http"

var defaultHeaders = map[string]string{
	"Accept":                    "*/*",
	"Accept-Encoding":           "gzip, deflate, br",
	"Accept-Language":           "en-US,en;q=0.9",
	"Cache-Control":             "no-cache",
	"Connection":                "keep-alive",
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0",
	"Upgrade-Insecure-Requests": "1",
	"sec-ch-ua":                 `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        "Windows",
	"sec-fetch-user":            "?1",
	"sec-fetch-dest":            "empty",
	"sec-fetch-mode":            "cors",
	"sec-fetch-site":            "same-origin",
	"Priority":                  "u=0, i",
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
}
