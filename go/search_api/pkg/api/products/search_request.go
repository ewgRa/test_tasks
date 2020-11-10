package products

func newSearchRequest() *searchRequest {
	return &searchRequest{}
}

type searchRequest struct {
	Query   string `form:"q,default="`
	Brand   string `form:"brand,default="`
	OrderBy string `form:"order_by,default=price" binding:"eq=title|eq=brand|eq=price|eq=stock"`
	Sort    string `form:"sort,default=asc" binding:"eq=asc|eq=desc"`
	Offset  int    `form:"offset,default=0" binding:"gte=0"`
	Limit   int    `form:"limit,default=10" binding:"gte=1"`
}
