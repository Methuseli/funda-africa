package utils

import (
	"fmt"
	"gin-funda-africa/internal/modules/users"
	"reflect"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type PaginateData struct {
	Model    interface{}
	Dto      interface{}
	Page     int
	PageSize int
	Params   interface{}
}

type QueryListData struct {
	Model          interface{}
	Filter         string
	Dto            interface{}
	FilterVariable interface{}
}

type Results struct {
	Data       []interface{}
	TotalCount int64
	Page       int
	PageSize   int
}

type Params struct {
	Version      *string `form:"version"`
	Country      *string `form:"country"`
	Board        *string `form:"board"`
	TopicID      *string `form:"topic_id"`
	CurriculumID *string `form:"curriculum_id"`
	Code         *string `form:"code"`
	Level        *string `form:"level"`
	SubjectID    *string `form:"subject_id"`
	Name         *string `form:"name"`
}

func GetPaginatedData(paginateData *PaginateData) (*Results, error) {
	var data []interface{}
	var totalCount int64

	query := users.Database.Model(paginateData.Model)

	rv := reflect.ValueOf(paginateData.Params)
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if !field.IsNil() && field.Elem().String() != "" {
			fieldName := rv.Type().Field(i).Tag.Get("form")
			switch fieldName {
			case "topic_id", "curriculum_id", "subject_id":
				uuid, err := uuid.Parse(*field.Interface().(*string))
				if err != nil {
					return nil, err
				}
				query = query.Where(fmt.Sprintf("%s = ?", fieldName), uuid)
			default:
				query = query.Where(fmt.Sprintf("%s = ?", fieldName), field.Elem().Interface())
			}

		}
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	page := paginateData.Page

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * paginateData.PageSize

	query = query.Limit(paginateData.PageSize).Offset(offset)

	if err := query.Find(paginateData.Model).Error; err != nil {
		return nil, err
	}

	var responses []interface{}
	if err := copier.Copy(&responses, &data); err != nil {
		return nil, err
	}

	return &Results{
		Data:       responses,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   paginateData.PageSize,
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
