package methods

import (
	"context"
	"demo/src/meclient/types"
	"net/http"
)

func (m *Methods) LaunchpadCollections(ctx context.Context, opts *LimitOffsetOpts) (*types.LaunchpadCollections, error) {
	path := "launchpad/collections"
	query := opts.String()
	v := new(types.LaunchpadCollections)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
