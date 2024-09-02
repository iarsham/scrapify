package scrapify

import (
	"compress/gzip"
	"compress/zlib"
	"github.com/andybalholm/brotli"
	"io"
	"net/http"
	"strings"
)

func ReadHTTPBody(resp *http.Response) (io.Reader, error) {
	encoding := strings.ToLower(resp.Header.Get("Content-Encoding"))
	switch encoding {
	case "gzip":
		zr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer zr.Close()
		return zr, err
	case "deflate":
		zr, err := zlib.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer zr.Close()
		return zr, err
	case "br":
		return brotli.NewReader(resp.Body), nil
	default:
		return resp.Body, nil
	}
}
