package scrapify

import (
	"compress/gzip"
	"compress/zlib"
	"github.com/andybalholm/brotli"
	"io"
	"net/http"
	"strings"
)

func ReadAll(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	encoding := strings.ToLower(resp.Header.Get("Content-Encoding"))
	switch encoding {
	case "gzip":
		zr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer zr.Close()
		return io.ReadAll(zr)
	case "deflate":
		zb, err := zlib.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer zb.Close()
		return io.ReadAll(zb)
	case "br":
		return io.ReadAll(brotli.NewReader(resp.Body))
	default:
		return io.ReadAll(resp.Body)
	}
}
