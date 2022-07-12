package methods

import (
	"context"
	"encoding/json"
	"fmt"
	
	"github.com/letsnode/meclient/types"
)

func (e *Executor) RawWallet(ctx context.Context, address string) ([]byte, error) {
	path := fmt.Sprintf("wallets/%s", address)
	return e.callMethod(ctx, "GET", path, "")
}

func (e *Executor) Wallet(ctx context.Context, address string) (*types.Wallet, error) {
	b, err := e.RawWallet(ctx, address)
	if err != nil {
		return nil, err
	}
	v := new(types.Wallet)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (e *Executor) RawWalletTokens(ctx context.Context, address string, opts *LimitOffsetOpts) ([]byte, error) {
	path := fmt.Sprintf("wallets/%s/tokens", address)
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

func (e *Executor) WalletTokens(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletTokens, error) {
	b, err := e.RawWalletTokens(ctx, address, opts)
	if err != nil {
		return nil, err
	}
	v := new(types.WalletTokens)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (e *Executor) RawWalletActivities(ctx context.Context, address string, opts *LimitOffsetOpts) ([]byte, error) {
	path := fmt.Sprintf("wallets/%s/activities", address)
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

func (e *Executor) WalletActivities(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletActivities, error) {
	b, err := e.RawWalletActivities(ctx, address, opts)
	if err != nil {
		return nil, err
	}
	v := new(types.WalletActivities)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (e *Executor) RawWalletOffersMade(ctx context.Context, address string, opts *LimitOffsetOpts) ([]byte, error) {
	path := fmt.Sprintf("wallets/%s/offers_made", address)
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

func (e *Executor) WalletOffersMade(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletOffersMade, error) {
	b, err := e.RawWalletOffersMade(ctx, address, opts)
	if err != nil {
		return nil, err
	}
	v := new(types.WalletOffersMade)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (e *Executor) RawWalletOffersReceived(ctx context.Context, address string, opts *LimitOffsetOpts) ([]byte, error) {
	path := fmt.Sprintf("wallets/%s/offers_received", address)
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

func (e *Executor) WalletOffersReceived(ctx context.Context, address string, opts *LimitOffsetOpts) (*types.WalletOffersReceived, error) {
	b, err := e.RawWalletOffersReceived(ctx, address, opts)
	if err != nil {
		return nil, err
	}
	v := new(types.WalletOffersReceived)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (e *Executor) RawWalletEscrowBalance(ctx context.Context, address string) ([]byte, error) {
	path := fmt.Sprintf("wallets/%s/escrow_balance", address)
	return e.callMethod(ctx, "GET", path, "")
}

func (e *Executor) WalletEscrowBalance(ctx context.Context, address string) (*types.WalletEscrowBalance, error) {
	b, err := e.RawWalletEscrowBalance(ctx, address)
	if err != nil {
		return nil, err
	}
	v := new(types.WalletEscrowBalance)
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
