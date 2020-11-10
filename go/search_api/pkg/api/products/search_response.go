package products

import (
	"sync"
)

func newSearchResponse() *searchResponse {
	productPool := &sync.Pool{
		New: func() interface{} { return &product{} },
	}

	return &searchResponse{productPool: productPool}
}

type searchResponse struct {
	Data []*product `json:"data"`

	productPool *sync.Pool
}

func (sr *searchResponse) product() *product {
	return sr.productPool.Get().(*product)
}

func (sr *searchResponse) reset() {
	for _, product := range sr.Data {
		sr.productPool.Put(product)
	}

	sr.Data = sr.Data[:0]
}

type product struct {
	Title string  `json:"title"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
