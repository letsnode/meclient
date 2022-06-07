package methods

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

const (
	Scheme = "https"
	MeHost = "api-mainnet.magiceden.dev"
)

type Methods struct {
	rateLimiter *rate.Limiter
}

func NewMethods(rateTime time.Duration, burst int) *Methods {
	return &Methods{
		rateLimiter: rate.NewLimiter(rate.Every(rateTime), burst),
	}
}

func (m *Methods) doRequest(ctx context.Context, method, path, query string) (*http.Response, error) {
	req := &http.Request{
		Method: method,
		URL: &url.URL{
			Scheme:   Scheme,
			Host:     MeHost,
			Path:     fmt.Sprintf("/v2/%s", path),
			RawQuery: query,
		},
	}
	req.WithContext(ctx)
	err := m.rateLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	cl := &http.Client{}
	return cl.Do(req)
}

func (m *Methods) callMethod(ctx context.Context, method, path, query string, v any) error {
	resp, err := m.doRequest(ctx, method, path, query)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}
