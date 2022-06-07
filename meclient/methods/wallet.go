package methods

import (
	"context"
	"fmt"
	"net/http"

	"meclient/types"
)

func (m *Methods) Wallet(ctx context.Context, address string) (*types.Wallet, error) {
	path := fmt.Sprintf("wallets/%s", address)
	v := new(types.Wallet)
	err := m.callMethod(ctx, http.MethodGet, path, "", v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletTokens(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletTokens, error) {
	path := fmt.Sprintf("wallets/%s/tokens", address)
	err := opts.CheckLimit(500)
	if err != nil {
		return nil, err
	}
	query := opts.String()
	v := new(types.WalletTokens)
	err = m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletActivities(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletActivities, error) {
	path := fmt.Sprintf("wallets/%s/activities", address)
	err := opts.CheckLimit(500)
	if err != nil {
		return nil, err
	}
	query := opts.String()
	v := new(types.WalletActivities)
	err = m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletOffersMade(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletOffersMade, error) {
	path := fmt.Sprintf("wallets/%s/offers_made", address)
	err := opts.CheckLimit(500)
	if err != nil {
		return nil, err
	}
	query := opts.String()
	v := new(types.WalletOffersMade)
	err = m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletOffersReceived(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletOffersReceived, error) {
	path := fmt.Sprintf("wallets/%s/offers_received", address)
	err := opts.CheckLimit(500)
	if err != nil {
		return nil, err
	}
	query := opts.String()
	v := new(types.WalletOffersReceived)
	err = m.callMethod(ctx, http.MethodGet, path, query, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Methods) WalletEscrowBalance(ctx context.Context, address string) (*types.WalletEscrowBalance, error) {
	path := fmt.Sprintf("wallets/%s/escrow_balance", address)
	v := new(types.WalletEscrowBalance)
	err := m.callMethod(ctx, http.MethodGet, path, "", v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
