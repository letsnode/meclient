package methods

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/letsnode/meclient/consts"
	"golang.org/x/time/rate"
)

type Executor struct {
	Client      *http.Client
	Request     *http.Request
	RateLimiter *rate.Limiter
}

func NewMethodExecutor(host string, rateTime time.Duration, burst int) *Executor {
	return &Executor{
		Client:      &http.Client{Timeout: 10},
		Request:     &http.Request{URL: &url.URL{Scheme: consts.Scheme, Host: host}},
		RateLimiter: rate.NewLimiter(rate.Every(rateTime), burst),
	}
}

func (e *Executor) doRequest(ctx context.Context, method, path, query string) (*http.Response, error) {
	req := e.Request.Clone(ctx)
	req.Method = method
	req.URL.Path = fmt.Sprintf("/%s/%s", consts.Version, path)
	req.URL.RawQuery = query
	err := e.RateLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	return e.Client.Do(req.WithContext(ctx))
}

func (e *Executor) callMethod(ctx context.Context, method, path, query string) ([]byte, error) {
	resp, err := e.doRequest(ctx, method, path, query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
