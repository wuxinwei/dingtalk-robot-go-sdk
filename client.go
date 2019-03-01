package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"time"
	"errors"
)

type Client struct {
	httpCli *http.Client
}

// NewClient return dingtalk client with a independent network state http client
func NewClient() *Client {
	return &Client{
		httpCli: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:       30 * time.Second,
					KeepAlive:     30 * time.Second,
					FallbackDelay: time.Millisecond * 300,
				}).DialContext,
				MaxIdleConns:          10,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConnsPerHost:   -1,
				DisableKeepAlives:     true,
			},
		},
	}
}

// SendMessage is action to send dingtalk message
func (cli *Client) SendMessage(ctx context.Context, r *Request) error {
	body, err := r.GetBody()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", stringBuilder(address, r.GetAccessToken()), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")
	respInRaw, err := cli.httpCli.Do(req)
	if err != nil {
		return err
	}
	defer respInRaw.Body.Close()

	resp := &Response{}
	if err := json.NewDecoder(respInRaw.Body).Decode(resp); err != nil {
	   return err
	}

	if resp.Errmsg != "ok" {
	   return errors.New(stringBuilder("failed to send dingtalk message: ", resp.Errmsg))
	}
	return nil
}

func stringBuilder(strs ...string) string {
	var stringBuilder = strings.Builder{}
	for _, s := range strs {
		stringBuilder.WriteString(s)
	}
	return stringBuilder.String()
}
