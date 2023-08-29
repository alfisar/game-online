package databaseinternal

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Data       interface{} `json:"data"`
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
}

func Paginations(tblName string, page int, limit int, pagi *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	db.Debug().Table(tblName).Count(&totalRows)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	pagi.TotalPages = int(math.Ceil(float64(totalRows) / float64(limit)))
	pagi.TotalRows = totalRows
	pagi.Limit = limit
	pagi.Page = page

	return func(db *gorm.DB) *gorm.DB {
		return db.Debug().Table(tblName).Offset(offset).Limit(limit)
	}
}
