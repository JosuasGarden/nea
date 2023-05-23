package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/knadh/koanf/v2"
	"github.com/knadh/listmonk/models"
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

func (w *WooCommerceProvider) ProvideTemplateFunctions(c *models.Campaign) template.FuncMap {
	return template.FuncMap{"WooProduct": func(productID int) (template.HTML, error) {
		wooProduct, err := w.GetProductByID(int(productID))
		if err != nil {
			return "", fmt.Errorf("Error getting WooCommerce product: %v", err)
		}

		product := models.Product{
			Name:        wooProduct.Name,
			Price:       wooProduct.Price,
			URL:         wooProduct.Permalink,
			TitleImage:  wooProduct.Images[len(wooProduct.Images)-1].Src,
			Description: wooProduct.Description,
		}

		out := bytes.Buffer{}
		err = c.ProductTpl.Execute(&out, product)
		if err != nil {
			return "", err
		}

		return template.HTML(out.Bytes()), nil
	}}
}

func (w *WooCommerceProvider) GetProductByID(productID int) (*WooCommerceProduct, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wp-json/wc/v3/products/%d", w.endpoint, productID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(w.consumerKey, w.consumerSecret)

	res, err := w.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	var product WooCommerceProduct
	if err := json.NewDecoder(res.Body).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

type WooCommerceProduct struct {
	ID                int64         `json:"id"`
	Name              string        `json:"name"`
	Slug              string        `json:"slug"`
	Permalink         string        `json:"permalink"`
	DateCreated       string        `json:"date_created"`
	DateCreatedGmt    string        `json:"date_created_gmt"`
	DateModified      string        `json:"date_modified"`
	DateModifiedGmt   string        `json:"date_modified_gmt"`
	Type              string        `json:"type"`
	Status            string        `json:"status"`
	Featured          bool          `json:"featured"`
	CatalogVisibility string        `json:"catalog_visibility"`
	Description       string        `json:"description"`
	ShortDescription  string        `json:"short_description"`
	Sku               string        `json:"sku"`
	Price             string        `json:"price"`
	RegularPrice      string        `json:"regular_price"`
	SalePrice         string        `json:"sale_price"`
	DateOnSaleFrom    interface{}   `json:"date_on_sale_from"`
	DateOnSaleFromGmt interface{}   `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      interface{}   `json:"date_on_sale_to"`
	DateOnSaleToGmt   interface{}   `json:"date_on_sale_to_gmt"`
	PriceHTML         string        `json:"price_html"`
	OnSale            bool          `json:"on_sale"`
	Purchasable       bool          `json:"purchasable"`
	TotalSales        int64         `json:"total_sales"`
	Virtual           bool          `json:"virtual"`
	Downloadable      bool          `json:"downloadable"`
	Downloads         []interface{} `json:"downloads"`
	DownloadLimit     int64         `json:"download_limit"`
	DownloadExpiry    int64         `json:"download_expiry"`
	ExternalURL       string        `json:"external_url"`
	ButtonText        string        `json:"button_text"`
	TaxStatus         string        `json:"tax_status"`
	TaxClass          string        `json:"tax_class"`
	ManageStock       bool          `json:"manage_stock"`
	StockQuantity     interface{}   `json:"stock_quantity"`
	StockStatus       string        `json:"stock_status"`
	Backorders        string        `json:"backorders"`
	BackordersAllowed bool          `json:"backorders_allowed"`
	Backordered       bool          `json:"backordered"`
	SoldIndividually  bool          `json:"sold_individually"`
	Weight            string        `json:"weight"`
	Dimensions        struct {
		Length string `json:"length"`
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"dimensions"`
	ShippingRequired bool          `json:"shipping_required"`
	ShippingTaxable  bool          `json:"shipping_taxable"`
	ShippingClass    string        `json:"shipping_class"`
	ShippingClassID  int64         `json:"shipping_class_id"`
	ReviewsAllowed   bool          `json:"reviews_allowed"`
	AverageRating    string        `json:"average_rating"`
	RatingCount      int64         `json:"rating_count"`
	RelatedIDS       []int64       `json:"related_ids"`
	UpsellIDS        []interface{} `json:"upsell_ids"`
	CrossSellIDS     []interface{} `json:"cross_sell_ids"`
	ParentID         int64         `json:"parent_id"`
	PurchaseNote     string        `json:"purchase_note"`
	Categories       []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"categories"`
	Tags []struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
		Count       int64  `json:"count"`
		Links       Links  `json:"_links"`
	} `json:"tags"`
	Images []struct {
		ID              int64  `json:"id"`
		DateCreated     string `json:"date_created"`
		DateCreatedGmt  string `json:"date_created_gmt"`
		DateModified    string `json:"date_modified"`
		DateModifiedGmt string `json:"date_modified_gmt"`
		Src             string `json:"src"`
		Name            string `json:"name"`
		Alt             string `json:"alt"`
	} `json:"images"`
	Attributes        []interface{} `json:"attributes"`
	DefaultAttributes []interface{} `json:"default_attributes"`
	Variations        []interface{} `json:"variations"`
	GroupedProducts   []interface{} `json:"grouped_products"`
	MenuOrder         int64         `json:"menu_order"`
	MetaData          []interface{} `json:"meta_data"`
	Links             Links         `json:"_links"`
}

type Links struct {
	Self       []Collection `json:"self"`
	Collection []Collection `json:"collection"`
}

type Collection struct {
	Href string `json:"href"`
}
