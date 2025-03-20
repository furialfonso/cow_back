package restful

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type IRestClient interface {
	EncodeFormData(data map[string]string) string
	Get(ctx context.Context, url string, timeOut string, headers map[string]string) ([]byte, error)
	Post(ctx context.Context, url string, timeOut string, headers map[string]string, data string) ([]byte, error)
}

type restClient struct{}

func NewRestClient() IRestClient {
	return &restClient{}
}

func (rs *restClient) Get(ctx context.Context, url string, timeOut string, headers map[string]string) ([]byte, error) {
	to, _ := time.ParseDuration(timeOut)
	ctx, cancel := context.WithTimeout(ctx, to)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	setHeaders(req, headers)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (rs *restClient) Post(ctx context.Context, url string, timeOut string, headers map[string]string, data string) ([]byte, error) {
	to, _ := time.ParseDuration(timeOut)
	ctx, cancel := context.WithTimeout(ctx, to)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(data))
	setHeaders(req, headers)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (rs *restClient) EncodeFormData(data map[string]string) string {
	form := url.Values{}
	for key, value := range data {
		form.Set(key, value)
	}
	return form.Encode()
}

func setHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}
