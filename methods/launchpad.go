package methods

import (
	"context"
	"net/http"

	"github.com/letsnode/meclient/types"
)

func (m *Methods) LaunchpadCollections(ctx context.Context, opts *LimitOffsetOpts) (*types.LaunchpadCollections, error) {
	path := "launchpad/collections"
	var query string
	if opts != nil {
		err := opts.CheckLimit(500)
		if err != nil {
			return nil, err
		}
		query = opts.String()
	}
	v := new(types.LaunchpadCollections)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
