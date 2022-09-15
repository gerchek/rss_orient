package storage

import (
	"rss/internal/model"
	"rss/pkg/gormquery"
	"strconv"

	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FetchPostsStorage interface {
	CreatePosts(items []*gofeed.Item, category string)
	GetAll(parameters map[string]interface{}) (data []*model.Post, err error)
	// Links
	LinkAll() ([]*model.Link, error)
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

func (db *fetchPostsStorage) CreatePosts(items []*gofeed.Item, category string) {
	for _, item := range items {
		post := model.Post{
			Category:     category,
			Title:        item.Title,
			Link:         item.Link,
			Publish_date: item.Published,
			Summary:      item.Description,
		}
		old_post := post

		r := db.client.Where("link = ?", &post.Link).Limit(1).Find(&old_post)
		if r.RowsAffected == 0 {
			db.client.Create(&post)
		} else {
			if old_post.Publish_date != post.Publish_date {
				//str := fmt.Sprintf("%s updated to %s", post.Date, new_post.Date)
				history := model.History{
					Old_published_at: old_post.Publish_date,
					New_published_at: post.Publish_date,
					PostID:           old_post.ID,
				}
				err := db.client.Model(&old_post).Association("HistoryList").Append(&history)
				if err != nil {
					db.logger.Warn(err)
				}
				if err := db.client.Model(model.Post{}).Where("id = ?", old_post.ID).Update("publish_date", post.Publish_date).Error; err != nil {
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
	if sortBy == "" {
		sortBy = "id.asc"
	}
	sortQuery, err := gormquery.ValidateAndReturnSortQuery(sortBy)
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
