package provider

import (
    "html/template"
)

type ProductProvider interface {
    RegisterFunctions(*template.FuncMap)
}
