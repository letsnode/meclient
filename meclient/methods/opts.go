package methods

import (
	"net/url"
	"strconv"
)

type LimitOffsetOpts struct {
	Limit  int
	Offset int
}

func (q *LimitOffsetOpts) String() string {
	v := &url.Values{}
	if q.Limit != 0 {
		v.Add("limit", strconv.Itoa(q.Limit))
	}
	if q.Offset != 0 {
		v.Add("limit", strconv.Itoa(q.Limit))
	}
	return v.Encode()
}
