package common

type PageRequest struct {
	PageIndex int64 `form:"pageIndex"`
	PageSize  int64 `form:"pageSize"`
}
