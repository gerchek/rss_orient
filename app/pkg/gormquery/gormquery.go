package gormquery

import (
	"errors"
	"fmt"
	"reflect"
	"rss/internal/model"
	"strings"
)

// type Query struct{}

// ---------------------------------------------sortBy start------------------------------------------------------------

func GetPostFields() []string {
	var field []string
	var test model.Post
	v := reflect.ValueOf(test)
	for i := 0; i < v.Type().NumField(); i++ {
		field = append(field, v.Type().Field(i).Tag.Get("json"))
	}
	return field
}

func ValidateAndReturnSortQuery(sortBy string) (string, error) {
	var userFields = GetPostFields()
	splits := strings.Split(sortBy, ".")
	if len(splits) != 2 {
		return "", errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}
	field, order := splits[0], splits[1]
	if order != "desc" && order != "asc" {
		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	}
	if !StringInSlice(userFields, field) {
		return "", errors.New("unknown field in sortBy query parameter")
	}
	return fmt.Sprintf("%s %s", field, strings.ToUpper(order)), nil
}

func StringInSlice(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if v == s {
			return true
		}
	}
	return false
}

// ---------------------------------------------sortBy end------------------------------------------------------------
// ---------------------------------------------filter start------------------------------------------------------------

func ValidateAndReturnFilterMap(filter string) (map[string]string, error) {
	var userFields = GetPostFields()
	splits := strings.Split(filter, ".")
	if len(splits) != 2 {
		return nil, errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}
	field, value := splits[0], splits[1]
	if !StringInSlice(userFields, field) {
		return nil, errors.New("unknown field in filter query parameter")
	}
	return map[string]string{field: value}, nil
}

// ---------------------------------------------filter end------------------------------------------------------------
