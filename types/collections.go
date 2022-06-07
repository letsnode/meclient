package types

type Collections []struct {
	Symbol      string   `json:"symbol"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Twitter     string   `json:"twitter"`
	Discord     string   `json:"discord"`
	Website     string   `json:"website"`
	Categories  []string `json:"categories"`
}

type CollectionsListings []struct {
	PdaAddress   string  `json:"pdaAddress"`
	AuctionHouse string  `json:"auctionHouse"`
	TokenAddress string  `json:"tokenAddress"`
	TokenMint    string  `json:"tokenMint"`
	Seller       string  `json:"seller"`
	TokenSize    int     `json:"tokenSize"`
	Price        float64 `json:"price"`
}

type CollectionActivities []struct {
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

type CollectionStats struct {
	Symbol      string  `json:"symbol"`
	FloorPrice  float64 `json:"floorPrice"`
	ListedCount int     `json:"listedCount"`
	VolumeAll   int64   `json:"volumeAll"`
}
