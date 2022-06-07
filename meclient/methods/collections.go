package methods

import (
	"context"
	"demo/src/meclient/types"
	"fmt"
	"net/http"
)

func (m *Methods) Collections(ctx context.Context, opts *LimitOffsetOpts) (*types.Collections, error) {
	path := "collections"
	query := opts.String()
	v := new(types.Collections)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) CollectionsListings(ctx context.Context, symbol string, opts *LimitOffsetOpts) (*types.CollectionsListings, error) {
	path := fmt.Sprintf("collections/%s/listings", symbol)
	query := opts.String()
	v := new(types.CollectionsListings)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) CollectionActivities(ctx context.Context, symbol string, opts *LimitOffsetOpts) (*types.CollectionActivities, error) {
	path := fmt.Sprintf("collections/%s/activities", symbol)
	query := opts.String()
	v := new(types.CollectionActivities)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) CollectionStats(ctx context.Context, symbol string, opts *LimitOffsetOpts) (*types.CollectionStats, error) {
	path := fmt.Sprintf("collections/%s/stats", symbol)
	query := opts.String()
	v := new(types.CollectionStats)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
