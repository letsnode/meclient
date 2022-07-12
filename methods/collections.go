package methods

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/letsnode/meclient/types"
)

func (e *Executor) RawCollections(ctx context.Context, opts *LimitOffsetOpts) ([]byte, error) {
	path := "collections"
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

func (e *Executor) Collections(ctx context.Context, opts *LimitOffsetOpts) (*types.Collections, error) {
	b, err := e.RawCollections(ctx, opts)
	if err != nil {
		return nil, err
	}1111
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (e *Executor) RawCollectionsListings(ctx context.Context, symbol string, opts *LimitOffsetOpts) ([]byte, error) {
	path := fmt.Sprintf("collections/%s/listings", symbol)
	var query string
	if opts != nil {
		err := opts.CheckLimit(20)
		if err != nil {
			return nil, err
		}
		query = opts.String()
	}
	return e.callMethod(ctx, "GET", path, query)
}

func (e *Executor) CollectionsListings(ctx context.Context, symbol string, opts *LimitOffsetOpts) (*types.CollectionsListings, error) {
	b, err := e.RawCollectionsListings(ctx, symbol, opts)
	if err != nil {
		return nil, err
	}
	v := new(types.CollectionsListings)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (e *Executor) RawCollectionActivities(ctx context.Context, symbol string, opts *LimitOffsetOpts) ([]byte, error) {
	path := fmt.Sprintf("collections/%s/activities", symbol)
	var query string
	if opts != nil {
		err := opts.CheckLimit(1_000)
		if err != nil {
			return nil, err
		}
		query = opts.String()
	}
	return e.callMethod(ctx, "GET", path, query)
}

func (e *Executor) CollectionActivities(ctx context.Context, symbol string, opts *LimitOffsetOpts) (*types.CollectionActivities, error) {
	b, err := e.RawCollectionActivities(ctx, symbol, opts)
	if err != nil {
		return nil, err
	}
	v := new(types.CollectionActivities)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (e *Executor) RawCollectionStats(ctx context.Context, symbol string) ([]byte, error) {
	path := fmt.Sprintf("collections/%s/stats", symbol)
	return e.callMethod(ctx, "GET", path, "")
}

func (e *Executor) CollectionStats(ctx context.Context, symbol string) (*types.CollectionStats, error) {
	b, err := e.RawCollectionStats(ctx, symbol)
	if err != nil {
		return nil, err
	}
	v := new(types.CollectionStats)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, err
}
