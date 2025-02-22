package models

type Pagination struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Search    string `json:"search"`
	OrderBy   string `json:"orderBy"`
	OrderDir  string `json:"orderDir"`
	TotalRows int64  `json:"totalRows"`
}

type PaginationResult struct {
	Items      interface{} `json:"items"`
	Page       int         `json:"page"`
	Limit      int        `json:"limit"`
	TotalRows  int64      `json:"totalRows"`
	TotalPages int        `json:"totalPages"`
}