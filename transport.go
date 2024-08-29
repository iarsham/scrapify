package scrapify

import (
	"crypto/tls"
	"net/http"
)

func NewTransport(browser Browser, debug bool) http.RoundTripper {
	tlsConfig := &tls.Config{
		MinVersion:         tls.VersionTLS12,
		MaxVersion:         tls.VersionTLS13,
		CipherSuites:       MapBrowserToCipher[browser],
		InsecureSkipVerify: debug,
	}
	return &http.Transport{
		TLSClientConfig: tlsConfig,
	}
}
