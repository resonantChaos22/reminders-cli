package client

import "net/http"

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		client:     &http.Client{},
		BackendURI: uri,
	}
}

type HTTPClient struct {
	client     *http.Client
	BackendURI string
}
