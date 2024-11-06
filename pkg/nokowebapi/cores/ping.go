package cores

import (
	"golang.org/x/net/http2"
	"net/http"
	"net/url"
	"time"
)

func TryFetchUrl(url *url.URL) bool {
	var err error
	var req *http.Request
	var res *http.Response
	KeepVoid(err, req, res)

	// create http transport
	transport := &http.Transport{
		DisableKeepAlives: true,
	}

	// create http client with binding http transport
	client := &http.Client{
		Transport: transport,
	}

	// configuration http2 transport
	if err = http2.ConfigureTransport(transport); err != nil {
		return false
	}

	// create new request
	if req, err = http.NewRequest("HEAD", url.String(), nil); err != nil {
		return false
	}

	// set http request prototype to http2
	req.Proto = "HTTP/2.0"

	// no keep alive
	req.Header.Del("Keep-Alive")
	req.Header.Set("Connection", "close")

	// dial http client
	if res, err = client.Do(req); err != nil {
		return false
	}

	// closing body request
	defer NoErr(res.Body.Close())

	// check status code available
	return 100 <= res.StatusCode && res.StatusCode < 500
}

func TryFetchUrlWaitForAlive(url *url.URL) {
	for {
		if TryFetchUrl(url) {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
