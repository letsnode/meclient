package methods

import (
	"context"
	"demo/src/meclient/types"
	"fmt"
	"net/http"
)

func (m *Methods) TokenMetadata(ctx context.Context, address string) (*types.TokenMetadata, error) {
	path := fmt.Sprintf("tokens/%s", address)
	v := new(types.TokenMetadata)
	err := m.callMethod(ctx, http.MethodGet, path, "", v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) TokenListings(ctx context.Context, address string) (*types.TokenListings, error) {
	path := fmt.Sprintf("tokens/%s/listings", address)
	v := new(types.TokenListings)
	err := m.callMethod(ctx, http.MethodGet, path, "", v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) TokenOfferReceived(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.TokenOfferReceived, error) {
	path := fmt.Sprintf("tokens/%s/listings", address)
	query := opts.String()
	v := new(types.TokenOfferReceived)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) TokenActivities(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.TokenActivities, error) {
	path := fmt.Sprintf("tokens/%s/activities", address)
	query := opts.String()
	v := new(types.TokenActivities)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
