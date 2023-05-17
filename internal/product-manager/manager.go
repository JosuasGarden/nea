package productmanager

import (
	"log"

	"github.com/knadh/koanf/v2"
	"github.com/knadh/listmonk/internal/product-manager/provider"
)

type ProductManager struct {
	provider []provider.ProductProvider
}

func InitProductManager(ko *koanf.Koanf, lo *log.Logger) *ProductManager {
	var (
		items              = ko.Slices("product_provider")
		configuredProvider = make([]provider.ProductProvider, 0, len(items))
	)

	// Load the config for the product provider.
	for _, item := range items {
		var (
			newProvider provider.ProductProvider
			err         error
		)

		if !item.Bool("enabled") {
			continue
		}

		switch providerType := item.String("type"); providerType {
		case "WooCommerce":
			newProvider, err = provider.NewWooCommerceProdiver(item)
		default:
			lo.Printf("unkown provider type: %s", providerType)
		}

		if err != nil {
			lo.Fatalf("error reading provider config: %v", err)
		} else {
			lo.Printf("setup %s", item.String("type"))
			configuredProvider = append(configuredProvider, newProvider)
		}
	}

	return &ProductManager{
		provider: configuredProvider,
	}
}

func GetAvalibleProductProvider() []interface{} {
	type p struct {
		Type   string      `json:"type"`
		Config interface{} `json:"config"`
	}

	return []interface{}{
		p{
			Type:   "WooCommerce",
			Config: provider.WooCommerceProviderConfig{},
		},
	}
}
