package v2

type product struct {
	ProductID  int            `json:"-"`
	UUID       string         `json:"uuid"`
	Name       string         `json:"name"`
	Brand      string         `json:"brand"`
	Stock      int            `json:"stock"`
	SellerUUID string         `json:"-"`
	Seller     *productSeller `json:"seller"`
}

type productSeller struct {
	UUID string                 `json:"uuid"`
	Link map[string]interface{} `json:"_links"`
}