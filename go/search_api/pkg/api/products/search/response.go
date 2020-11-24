package search

import (
	"sync"
)

func newResponse() *response {
	productPool := &sync.Pool{
		New: func() interface{} { return &product{} },
	}

	return &response{productPool: productPool}
}

type response struct {
	Data []*product `json:"data"`

	productPool *sync.Pool
}

func (sr *response) product() *product {
	return sr.productPool.Get().(*product)
}

func (sr *response) reset() {
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
