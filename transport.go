package scrapify

import (
	"crypto/tls"
	"net/http"
)

func NewTransport(browser Browser) http.RoundTripper {
	tlsConfig := &tls.Config{
		MinVersion:         tls.VersionTLS12,
		MaxVersion:         tls.VersionTLS13,
		CipherSuites:       MapBrowserToCipher[browser],
		InsecureSkipVerify: true,
	}
	return &http.Transport{
		TLSClientConfig: tlsConfig,
	}
}
