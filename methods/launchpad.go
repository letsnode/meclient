package methods

import (
	"context"
	"encoding/json"

	"github.com/letsnode/meclient/types"
)

func (e *Executor) RawLaunchpadCollections(ctx context.Context, opts *LimitOffsetOpts) ([]byte, error) {
	path := "launchpad/collections"
	var query string
	if opts != nil {
		err := opts.CheckLimit(500)
		if err != nil {
			return nil, err
		}
		query = opts.String()
	}
	return e.callMethod(ctx, "GET", path, query)
}

func (e *Executor) LaunchpadCollections(ctx context.Context, opts *LimitOffsetOpts) (*types.LaunchpadCollections, error) {
	b, err := e.RawLaunchpadCollections(ctx, opts)
	if err != nil {
		return nil, err
	}
	v := new(types.LaunchpadCollections)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, err
}
