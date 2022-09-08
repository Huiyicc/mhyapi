package request

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func HttpGet(uri RequestStruct, header http.Header) ([]byte, error) {
	urls := uri.Url
	if uri.Query != "" {
		urls += "?" + uri.Query
	}
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)
	if req, err = http.NewRequest("GET", urls, nil); err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	cli := http.Client{}
	if ProxyUrl != "" {
		httpProxy(&cli)
	}
	if resp, err = cli.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func HttpPost(uri RequestStruct, header http.Header) ([]byte, error) {
	urls := uri.Url
	if uri.Query != "" {
		urls += "?" + uri.Query
	}
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	if req, err = http.NewRequest("POST", urls, strings.NewReader(uri.Body.Get())); err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	cli := http.Client{}
	if ProxyUrl != "" {
		httpProxy(&cli)
	}
	if resp, err = cli.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// ProxyUrl 为代理链接,修改此值用以启用代理,非线程安全
var ProxyUrl string

func httpProxy(client *http.Client) {
	proxyStr := ProxyUrl
	urls, _ := url.Parse(proxyStr)
	p := http.ProxyURL(urls)
	client.Transport = &http.Transport{
		Proxy: p,
	}
}
