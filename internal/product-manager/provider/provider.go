package provider

import (
	"html/template"

	"github.com/knadh/listmonk/models"
)

type ProductProvider interface {
    ProvideTemplateFunctions(c *models.Campaign) template.FuncMap
}
