package panda

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (p Handler) post(uri string, query, values url.Values) (*http.Response, error) {
	req, _ := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return p.client.Do(req)
}

func (p Handler) get(uri string, query url.Values) (*http.Response, error) {
	req, _ := http.NewRequest("GET", uri, new(bytes.Buffer))
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	req.Header.Add("User-Agent", userAgent)

	return p.client.Do(req)
}

func (p Handler) Download(path string, url string) error {
	resp, err := p.get(url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
