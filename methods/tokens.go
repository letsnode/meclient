package methods

import (
	"context"
	"encoding/json"
	"fmt"
	
	"github.com/letsnode/meclient/types"
)

func (e *Executor) RawTokenMetadata(ctx context.Context, address string) ([]byte, error) {
	path := fmt.Sprintf("tokens/%s", address)
	return e.callMethod(ctx, "GET", path, "")
}

func (e *Executor) TokenMetadata(ctx context.Context, address string) (*types.TokenMetadata, error) {
	b, err := e.RawTokenMetadata(ctx, address)
	if err != nil {
		return nil, err
	}
	v := new(types.TokenMetadata)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (e *Executor) RawTokenListings(ctx context.Context, address string) ([]byte, error) {
	path := fmt.Sprintf("tokens/%s/listings", address)
	return e.callMethod(ctx, "GET", path, "")
}

func (e *Executor) TokenListings(ctx context.Context, address string) (*types.TokenListings, error) {
	b, err := e.RawTokenListings(ctx, address)
	if err != nil {
		return nil, err
	}
	v := new(types.TokenListings)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (e *Executor) RawTokenOfferReceived(ctx context.Context, address string, opts *LimitOffsetOpts) ([]byte, error) {
	path := fmt.Sprintf("tokens/%s/listings", address)
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

func (e *Executor) TokenOfferReceived(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.TokenOfferReceived, error) {
	b, err := e.RawTokenOfferReceived(ctx, address, opts)
	if err != nil {
		return nil, err
	}
	v := new(types.TokenOfferReceived)
	err = json.Unmarshal(b, v)
	return v, err
}

func (e *Executor) RawTokenActivities(ctx context.Context, address string, opts *LimitOffsetOpts) ([]byte, error) {
	path := fmt.Sprintf("tokens/%s/activities", address)
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

func (e *Executor) TokenActivities(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.TokenActivities, error) {
	b, err := e.RawTokenActivities(ctx, address, opts)
	if err != nil {
		return nil, err
	}
	v := new(types.TokenActivities)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
