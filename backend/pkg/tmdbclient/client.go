package tmdbclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// Client TMDB API 客户端
type Client struct {
	apiKey     string
	baseURL    string
	language   string
	httpClient *http.Client

	// 简单令牌桶限流
	rateLimiter chan struct{}
	once        sync.Once
}

// NewClient 创建 TMDB 客户端
func NewClient(apiKey, baseURL, defaultLanguage string, rateLimit int) *Client {
	c := &Client{
		apiKey:   apiKey,
		baseURL:  baseURL,
		language: defaultLanguage,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		rateLimiter: make(chan struct{}, rateLimit),
	}

	// 填充令牌桶
	for i := 0; i < rateLimit; i++ {
		c.rateLimiter <- struct{}{}
	}

	// 定时补充令牌
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rateLimit))
		defer ticker.Stop()
		for range ticker.C {
			select {
			case c.rateLimiter <- struct{}{}:
			default:
			}
		}
	}()

	return c
}

// RequestOption 请求选项
type RequestOption struct {
	Language         string
	Page             int
	Region           string
	AppendToResponse string
	ExtraParams      map[string]string
}

// Get 发送 GET 请求到 TMDB API
func (c *Client) Get(path string, opts *RequestOption) (json.RawMessage, error) {
	// 限流
	<-c.rateLimiter

	reqURL, err := c.buildURL(path, opts)
	if err != nil {
		return nil, fmt.Errorf("构建请求 URL 失败: %w", err)
	}

	logx.Debugf("TMDB 请求: %s", reqURL)

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("TMDB 请求失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB 返回错误状态码 %d: %s", resp.StatusCode, string(body))
	}

	return json.RawMessage(body), nil
}

// Request 通用请求方法（Get 的别名），供中间件透传路径使用
func (c *Client) Request(path string, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(path, opts)
}

// buildURL 构建完整的 TMDB API URL
func (c *Client) buildURL(path string, opts *RequestOption) (string, error) {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("api_key", c.apiKey)

	// 语言参数
	lang := c.language
	if opts != nil && opts.Language != "" {
		lang = opts.Language
	}
	q.Set("language", lang)

	if opts != nil {
		if opts.Page > 0 {
			q.Set("page", fmt.Sprintf("%d", opts.Page))
		}
		if opts.Region != "" {
			q.Set("region", opts.Region)
		}
		if opts.AppendToResponse != "" {
			q.Set("append_to_response", opts.AppendToResponse)
		}
		for k, v := range opts.ExtraParams {
			q.Set(k, v)
		}
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
