package models

type Pagination struct {
	PageNumber uint64 `form:"pageNumber,default=1" binding:"number,gte=1"`
	PageSize   uint32 `form:"pageSize,default=10" binding:"number,gte=10,lte=100"`
}

type PaginationMetadataAPI struct {
	CurrentPageNumber uint64 `json:"currentPageNumber"`
	PageCount         uint64 `json:"pageCount"`
	PageSize          uint32 `json:"pageSize"`
	RecordCount       uint64 `json:"recordCount"`
}
