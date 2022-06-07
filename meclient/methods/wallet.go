package methods

import (
	"context"
	"demo/src/meclient/types"
	"fmt"
	"net/http"
)

func (m *Methods) Wallet(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.Wallet, error) {
	path := fmt.Sprintf("wallets/%s", address)
	query := opts.String()
	v := new(types.Wallet)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletTokens(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletTokens, error) {
	path := fmt.Sprintf("wallets/%s/tokens", address)
	query := opts.String()
	v := new(types.WalletTokens)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletActivities(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletActivities, error) {
	path := fmt.Sprintf("wallets/%s/activities", address)
	query := opts.String()
	v := new(types.WalletActivities)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletOffersMade(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletOffersMade, error) {
	path := fmt.Sprintf("wallets/%s/offers_made", address)
	query := opts.String()
	v := new(types.WalletOffersMade)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletOffersReceived(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletOffersReceived, error) {
	path := fmt.Sprintf("wallets/%s/offers_received", address)
	query := opts.String()
	v := new(types.WalletOffersReceived)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletEscrowBalance(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletEscrowBalance, error) {
	path := fmt.Sprintf("wallets/%s/escrow_balance", address)
	query := opts.String()
	v := new(types.WalletEscrowBalance)
	err := m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
