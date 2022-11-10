package confs

import (
	"net/http"
	"net/url"
	"os"
)

func SetupProxyConnection() *http.Client {
	pHost := os.Getenv("PROXY_HOST")
	pPort := os.Getenv("PROXY_PORT")

	pUrl, _ := url.Parse(pHost + ":" + pPort)
	//proxyUrl, _ := url.Parse("socks5://192.168.100.6:1080")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(pUrl)}}

	return client
}
