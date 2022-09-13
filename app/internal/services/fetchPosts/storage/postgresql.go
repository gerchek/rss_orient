package storage

import (
	"errors"
	"fmt"
	"reflect"
	"rss/internal/model"
	"strconv"
	"strings"

	"github.com/SlyMarbo/rss"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FetchPostsStorage interface {
	CreatePosts(items []*rss.Item, category string)
	GetAll(parameters map[string]interface{}) (data []*model.Post, err error)
	// Links
	LinkAll() ([]*model.Link, error)
	getPostFields() []string
	validateAndReturnSortQuery(sortBy string) (string, error)
	validateAndReturnFilterMap(filter string) (map[string]string, error)
	stringInSlice(strSlice []string, s string) bool
}

type fetchPostsStorage struct {
	client *gorm.DB
	logger *logrus.Logger
}

func NewFetchPostsStorage(client *gorm.DB, logger *logrus.Logger) FetchPostsStorage {
	return &fetchPostsStorage{
		client: client,
		logger: logger,
	}
}

func (db *fetchPostsStorage) CreatePosts(items []*rss.Item, category string) {
	for _, item := range items {
		post := model.Post{
			Category: category,
			Title:    item.Title,
			Link:     item.Link,
			Date:     item.Date,
			Summary:  item.Summary,
		}
		old_post := post

		r := db.client.Where("link = ?", &post.Link).Limit(1).Find(&old_post)
		if r.RowsAffected == 0 {
			db.client.Create(&post)
		} else {
			if old_post.Date != post.Date {
				//str := fmt.Sprintf("%s updated to %s", post.Date, new_post.Date)
				history := model.History{
					Old_published_at: old_post.Date,
					New_published_at: post.Date,
					PostID:           old_post.ID,
				}
				err := db.client.Model(&old_post).Association("HistoryList").Append(&history)
				if err != nil {
					db.logger.Warn(err)
				}
				if err := db.client.Model(model.Post{}).Where("id = ?", old_post.ID).Update("date", post.Date).Error; err != nil {
					db.logger.Warn(err)
				}
			}
		}
	}
}

func (db *fetchPostsStorage) GetAll(parameters map[string]interface{}) (data []*model.Post, err error) {
	if parameters["sortBy"] == nil {
		parameters["sortBy"] = "id.asc"
	}

	sortBy := parameters["sortBy"].(string)
	strLimit := parameters["strLimit"].(string)
	strOffset := parameters["strOffset"].(string)
	filter := parameters["filter"].(string)

	limit := -1
	if strLimit != "" {
		limit, err = strconv.Atoi(strLimit)
		if err != nil || limit < -1 {
			db.logger.Warn(err)
		}
	}
	offset := 1
	if strOffset != "" {
		offset, err = strconv.Atoi(strOffset)
		if err != nil || offset < 0 {
			db.logger.Warn(err)
		}
	}
	// filter start
	// filter end
	if sortBy == "" {
		sortBy = "id.asc"
	}
	sortQuery, err := db.validateAndReturnSortQuery(sortBy)
	if err != nil {
		db.logger.Warn(err)
	}
	new_offset := (offset - 1) * limit
	var posts []*model.Post
	if err := db.client.Where("title LIKE ? or link LIKE ? or summary LIKE ?", "%"+filter+"%", "%"+filter+"%", "%"+filter+"%").Limit(limit).Offset(new_offset).Order(sortQuery).Preload("HistoryList").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// LINKS

func (db *fetchPostsStorage) LinkAll() ([]*model.Link, error) {
	var links []*model.Link
	if err := db.client.Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

// ---------------------------------------------sortBy start------------------------------------------------------------

func (db *fetchPostsStorage) getPostFields() []string {
	var field []string
	var test model.Post
	v := reflect.ValueOf(test)
	for i := 0; i < v.Type().NumField(); i++ {
		field = append(field, v.Type().Field(i).Tag.Get("json"))
	}
	return field
}

func (db *fetchPostsStorage) validateAndReturnSortQuery(sortBy string) (string, error) {
	var userFields = db.getPostFields()
	splits := strings.Split(sortBy, ".")
	if len(splits) != 2 {
		return "", errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}
	field, order := splits[0], splits[1]
	if order != "desc" && order != "asc" {
		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	}
	if !db.stringInSlice(userFields, field) {
		return "", errors.New("unknown field in sortBy query parameter")
	}
	return fmt.Sprintf("%s %s", field, strings.ToUpper(order)), nil
}

func (db *fetchPostsStorage) stringInSlice(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if v == s {
			return true
		}
	}
	return false
}

// ---------------------------------------------sortBy end------------------------------------------------------------
// ---------------------------------------------filter start------------------------------------------------------------

func (db *fetchPostsStorage) validateAndReturnFilterMap(filter string) (map[string]string, error) {
	var userFields = db.getPostFields()
	splits := strings.Split(filter, ".")
	if len(splits) != 2 {
		return nil, errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}
	field, value := splits[0], splits[1]
	if !db.stringInSlice(userFields, field) {
		return nil, errors.New("unknown field in filter query parameter")
	}
	return map[string]string{field: value}, nil
}

// ---------------------------------------------filter end------------------------------------------------------------
