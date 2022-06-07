package meclient

//
//import (
//	"context"
//	"demo/src/meclient/types"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//)
//
//const (
//	Scheme = "https"
//	MeHost = "api-mainnet.magiceden.dev"
//)
//
//func (m *MeClient) doRequest(ctx context.Context, path, query string) (*http.Response, error) {
//	req := &http.Request{
//		Method: http.MethodGet,
//		URL: &url.URL{
//			Scheme:   Scheme,
//			Host:     MeHost,
//			Path:     fmt.Sprintf("/v2/%s", path),
//			RawQuery: query,
//		},
//	}
//	req.WithContext(ctx)
//	err := m.rateLimiter.Wait(ctx)
//	if err != nil {
//		return nil, err
//	}
//	cl := &http.Client{}
//	return cl.Do(req)
//}
//
//func (m *MeClient) callMethod(ctx context.Context, path, query string, v any) error {
//	resp, err := m.doRequest(ctx, path, query)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return err
//	}
//	err = json.Unmarshal(body, v)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (m *MeClient) TokenMetadata(ctx context.Context, address string) (*types.TokenMetadata, error) {
//	path := fmt.Sprintf("tokens/%s", address)
//	v := new(types.TokenMetadata)
//	err := m.callMethod(ctx, path, "", v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) TokenListings(ctx context.Context, address string) (*types.TokenListings, error) {
//	path := fmt.Sprintf("tokens/%s/listings", address)
//	v := new(types.TokenListings)
//	err := m.callMethod(ctx, path, "", v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) TokenOfferReceived(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.TokenOfferReceived, error) {
//	path := fmt.Sprintf("tokens/%s/listings", address)
//	query := opts.String()
//	v := new(types.TokenOfferReceived)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) TokenActivities(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.TokenActivities, error) {
//	path := fmt.Sprintf("tokens/%s/activities", address)
//	query := opts.String()
//	v := new(types.TokenActivities)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) WalletTokens(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletTokens, error) {
//	path := fmt.Sprintf("wallets/%s/tokens", address)
//	query := opts.String()
//	v := new(types.WalletTokens)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) WalletActivities(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletActivities, error) {
//	path := fmt.Sprintf("wallets/%s/activities", address)
//	query := opts.String()
//	v := new(types.WalletActivities)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) WalletOffersMade(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletOffersMade, error) {
//	path := fmt.Sprintf("wallets/%s/offers_made", address)
//	query := opts.String()
//	v := new(types.WalletOffersMade)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) WalletOffersReceived(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletOffersReceived, error) {
//	path := fmt.Sprintf("wallets/%s/offers_received", address)
//	query := opts.String()
//	v := new(types.WalletOffersReceived)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) WalletEscrowBalance(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletEscrowBalance, error) {
//	path := fmt.Sprintf("wallets/%s/escrow_balance", address)
//	query := opts.String()
//	v := new(types.WalletEscrowBalance)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) Wallet(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.Wallet, error) {
//	path := fmt.Sprintf("wallets/%s", address)
//	query := opts.String()
//	v := new(types.Wallet)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) Collections(ctx context.Context, opts *LimitOffsetOpts) (*types.Collections, error) {
//	path := "collections"
//	query := opts.String()
//	v := new(types.Collections)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) CollectionsListings(ctx context.Context, symbol string, opts *LimitOffsetOpts) (*types.CollectionsListings, error) {
//	path := fmt.Sprintf("collections/%s/listings", symbol)
//	query := opts.String()
//	v := new(types.CollectionsListings)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) CollectionActivities(ctx context.Context, symbol string, opts *LimitOffsetOpts) (*types.CollectionActivities, error) {
//	path := fmt.Sprintf("collections/%s/activities", symbol)
//	query := opts.String()
//	v := new(types.CollectionActivities)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) CollectionStats(ctx context.Context, symbol string, opts *LimitOffsetOpts) (*types.CollectionStats, error) {
//	path := fmt.Sprintf("collections/%s/stats", symbol)
//	query := opts.String()
//	v := new(types.CollectionStats)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (m *MeClient) LaunchpadCollections(ctx context.Context, opts *LimitOffsetOpts) (*types.LaunchpadCollections, error) {
//	path := "launchpad/collections"
//	query := opts.String()
//	v := new(types.LaunchpadCollections)
//	err := m.callMethod(ctx, path, query, v)
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
