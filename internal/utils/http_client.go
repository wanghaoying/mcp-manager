package utils

import (
	"context"
	"io"
	"net/http"
	"time"
)

// HTTPClient 封装 http 访问能力，便于 mock 和扩展
type HTTPClient interface {
	DoRequest(ctx context.Context, method, url string, body io.Reader) (string, error)
}

// HTTPClientOption 用于自定义 http client 配置
type HTTPClientOption func(*http.Client)

// DefaultHTTPClient 支持自定义超时和 Transport
type DefaultHTTPClient struct {
	client *http.Client
}

// NewHTTPClient 支持自定义超时和 Transport
func NewHTTPClient(opts ...HTTPClientOption) HTTPClient {
	c := &http.Client{}
	for _, opt := range opts {
		opt(c)
	}
	return &DefaultHTTPClient{client: c}
}

// WithTimeout 设置超时时间
func WithTimeout(timeoutSec int) HTTPClientOption {
	return func(c *http.Client) {
		c.Timeout = time.Duration(timeoutSec) * time.Second
	}
}

// WithTransport 设置自定义 Transport
func WithTransport(transport http.RoundTripper) HTTPClientOption {
	return func(c *http.Client) {
		c.Transport = transport
	}
}

func (c *DefaultHTTPClient) DoRequest(ctx context.Context, method, url string, body io.Reader) (string, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return "", err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBytes), nil
}

// DefaultTransport 返回带有合理默认配置的 http.Transport
func DefaultTransport() http.RoundTripper {
	return &http.Transport{
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
