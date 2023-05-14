package provider

import (
	"html/template"
	"net/http"

	"github.com/knadh/koanf/v2"
)

type WooCommerceProvider struct {
	endpoint       string
	consumerKey    string
	consumerSecret string
	httpClient     *http.Client
}

type WooCommerceProviderConfig struct {
	Endpoint       string `json:"endpoint"`
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

func NewWooCommerceProdiver(ko *koanf.Koanf) (*WooCommerceProvider, error) {
	var conifg WooCommerceProviderConfig

	if err := ko.UnmarshalWithConf("config", &conifg, koanf.UnmarshalConf{Tag: "json"}); err != nil {
		return nil, err
	}

	return &WooCommerceProvider{
		endpoint:       conifg.Endpoint,
		consumerKey:    conifg.ConsumerKey,
		consumerSecret: conifg.ConsumerSecret,
		httpClient:     &http.Client{},
	}, nil
}

func (w *WooCommerceProvider) RegisterFunctions(func_map *template.FuncMap) {

}

// func (c *Client) GetProductByID(productID int) (*Product, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wp-json/wc/v3/products/%d", c.endpoint, productID), nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req.SetBasicAuth(c.consumerKey, c.consumerSecret)
//
// 	res, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()
//
// 	if res.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
// 	}
//
// 	var product Product
// 	if err := json.NewDecoder(res.Body).Decode(&product); err != nil {
// 		return nil, err
// 	}
//
// 	return &product, nil
// }
//
// func NewClient(endpoint string, consumerKey string, consumerSecret string) *Client {
// 	client := Client{
// 		endpoint:        endpoint,
// 		consumerKey:    consumerKey,
// 		consumerSecret: consumerSecret,
// 		httpClient:     &http.Client{},
// 	}
//
// 	return &client
// }
