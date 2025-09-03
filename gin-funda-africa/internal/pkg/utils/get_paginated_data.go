package utils

import (
	"gin-funda-africa/internal/modules/users"

	"github.com/jinzhu/copier"
)

type PaginateData struct {
	Model      interface{}
	Filter     string
	Dto        interface{}
	Page       int
	PageSize   int
	FilterVariable interface{}
}

type QueryListData struct {
	Model      interface{}
	Filter     string
	Dto        interface{}
	FilterVariable interface{}
}

type Results struct {
	Data       []interface{}
	TotalCount int64
	Page       int
	PageSize   int
}

func GetPaginatedData(paginateData *PaginateData) (*Results, error) {
	var data []interface{}
	var totalCount int64

	countQuery := users.Database.Model(paginateData.Model).Where(paginateData.Filter, paginateData.FilterVariable)

	if err := countQuery.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	page := paginateData.Page

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * paginateData.PageSize

	query := users.Database.Model(paginateData.Model).Where(paginateData.Filter, paginateData.FilterVariable).Limit(paginateData.PageSize).Offset(offset)

	if err := query.Find(paginateData.Model).Error; err != nil {
		return nil, err
	}

	var responses []interface{}
	if err := copier.Copy(&responses, &data); err != nil {
		return nil, err
	}

	return &Results{
		Data: responses,
		TotalCount: totalCount,
		Page: page,
		PageSize: paginateData.PageSize,
	}, nil
}

func QueryList(queryListData *QueryListData) ([]interface{}, error) {
	var data []interface{}

	query := users.Database.Model(queryListData.Model).Where(queryListData.Filter, queryListData.FilterVariable)

	if err := query.Find(queryListData.Model).Error; err != nil {
		return nil, err
	}

	var responses []interface{}
	if err := copier.Copy(&responses, &data); err != nil {
		return nil, err
	}

	return responses, nil
}
