package types

type WalletTokens []struct {
	MintAddress          string `json:"mintAddress"`
	Owner                string `json:"owner"`
	Supply               int    `json:"supply"`
	Name                 string `json:"name"`
	UpdateAuthority      string `json:"updateAuthority"`
	PrimarySaleHappened  bool   `json:"primarySaleHappened"`
	SellerFeeBasisPoints int    `json:"sellerFeeBasisPoints"`
	Image                string `json:"image"`
	Attributes           []any  `json:"attributes"`
	Properties           any    `json:"properties"`
	ListStatus           string `json:"listStatus"`
}

type WalletActivities []struct {
	Signature      string  `json:"signature"`
	Type           string  `json:"type"`
	Source         string  `json:"source"`
	TokenMint      string  `json:"tokenMint"`
	Collection     string  `json:"collection"`
	Slot           int     `json:"slot"`
	BlockTime      int64   `json:"blockTime"`
	Buyer          string  `json:"buyer"`
	BuyerReferral  string  `json:"buyerReferral"`
	Seller         string  `json:"seller"`
	SellerReferral string  `json:"sellerReferral"`
	Price          float64 `json:"price"`
}

type WalletOffersMade []struct {
	PdaAddress   string  `json:"pdaAddress"`
	TokenMint    string  `json:"tokenMint"`
	AuctionHouse string  `json:"auctionHouse"`
	Buyer        string  `json:"buyer"`
	Price        float64 `json:"price"`
	TokenSize    int     `json:"tokenSize"`
	Expiry       int     `json:"expiry"`
}

type WalletOffersReceived []struct {
	PdaAddress   string  `json:"pdaAddress"`
	TokenMint    string  `json:"tokenMint"`
	AuctionHouse string  `json:"auctionHouse"`
	Buyer        string  `json:"buyer"`
	Price        float64 `json:"price"`
	TokenSize    int     `json:"tokenSize"`
	Expiry       int     `json:"expiry"`
}

type WalletEscrowBalance struct {
	BuyerEscrow string  `json:"buyerEscrow"`
	Balance     float64 `json:"balance"`
}

type Wallet struct {
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
}
