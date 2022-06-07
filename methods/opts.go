package methods

import (
	"net/url"
	"strconv"

	"github.com/letsnode/meclient"
)

type LimitOffsetOpts struct {
	Limit  int
	Offset int
}

func (q *LimitOffsetOpts) CheckLimit(maxLimit int) error {
	if q.Limit <= 0 || q.Limit > maxLimit {
		return meclient.ErrInvalidQueryParam
	}
	return nil
}

func (q *LimitOffsetOpts) String() string {
	v := &url.Values{}
	v.Add("limit", strconv.Itoa(q.Limit))
	if q.Offset != 0 {
		v.Add("limit", strconv.Itoa(q.Limit))
	}
	return v.Encode()
}
