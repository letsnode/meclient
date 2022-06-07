package types

type TokenMetadata struct {
	MintAddress          string        `json:"mintAddress"`
	Owner                string        `json:"owner"`
	Supply               int           `json:"supply"`
	Collection           string        `json:"collection"`
	Name                 string        `json:"name"`
	UpdateAuthority      string        `json:"updateAuthority"`
	PrimarySaleHappened  string        `json:"primarySaleHappened"`
	SellerFeeBasisPoints string        `json:"sellerFeeBasisPoints"`
	Image                string        `json:"image"`
	AnimationUrl         string        `json:"animationUrl"`
	ExternalUrl          string        `json:"externalUrl"`
	Attributes           []interface{} `json:"attributes"`
	Properties           []interface{} `json:"properties"`
	Category             string        `json:"category"`
	Creators             []struct {
		Address string `json:"address"`
		Share   int    `json:"share"`
	} `json:"creators"`
}

type TokenListings []struct {
	PdaAddress   string  `json:"pdaAddress"`
	AuctionHouse string  `json:"auctionHouse"`
	TokenAddress string  `json:"tokenAddress"`
	TokenMint    string  `json:"tokenMint"`
	Seller       string  `json:"seller"`
	TokenSize    int     `json:"tokenSize"`
	Price        float64 `json:"price"`
	Expiry       int     `json:"expiry"`
}

type TokenOfferReceived []struct {
	PdaAddress    string  `json:"pdaAddress"`
	TokenMint     string  `json:"tokenMint"`
	AuctionHouse  string  `json:"auctionHouse"`
	Buyer         string  `json:"buyer"`
	BuyerReferral string  `json:"buyerReferral"`
	TokenSize     int     `json:"tokenSize"`
	Price         float64 `json:"price"`
	Expiry        int     `json:"expiry"`
}

type TokenActivities []struct {
	Signature        string  `json:"signature"`
	Type             string  `json:"type"`
	Source           string  `json:"source"`
	TokenMint        string  `json:"tokenMint"`
	CollectionSymbol string  `json:"collectionSymbol"`
	Slot             int     `json:"slot"`
	BlockTime        int64   `json:"blockTime"`
	BuyerReferral    string  `json:"buyerReferral"`
	Seller           string  `json:"seller"`
	SellerReferral   string  `json:"sellerReferral"`
	Price            float64 `json:"price"`
}
