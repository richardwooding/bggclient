package xml2

import (
	"github.com/richardwooding/bggclient/client"
)

type API struct {
	baseURL string
}

func NewXML2API(options client.Options) *API {
	return &API{
		baseURL: options.BaseURL,
	}
}
