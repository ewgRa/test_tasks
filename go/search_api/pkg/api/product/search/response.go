package search

import "github.com/ewgRa/test_tasks/go/search_api/pkg/api/product"

func newResponse() *response {
	return &response{}
}

type response struct {
	Products []*product.Product `json:"products"`
}
