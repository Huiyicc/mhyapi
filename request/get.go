package request

import (
	"io"
	"net/http"
	"strings"
)

func HttpGet(uri RequestStruct, header http.Header) ([]byte, error) {
	url := uri.Url
	if uri.Query != "" {
		url += "?" + uri.Query
	}
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	cli := http.Client{}
	if resp, err = cli.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func HttpPost(uri RequestStruct, header http.Header) ([]byte, error) {
	url := uri.Url
	if uri.Query != "" {
		url += "?" + uri.Query
	}
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)
	if req, err = http.NewRequest("POST", url, strings.NewReader(uri.Body.Get())); err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	cli := http.Client{}
	if resp, err = cli.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
