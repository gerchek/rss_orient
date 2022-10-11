package storage

import (
	"rss/internal/model"
	"rss/pkg/gormquery"
	"strconv"
	"time"
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
		parsed_time, err := time.Parse(time.RFC1123Z, item.Published)
		if err != nil {
			db.logger.Warn(err)
		}
		post := model.Post{
			Category:     category,
			Title:        item.Title,
			Link:         item.Link,
			Publish_date: parsed_time,
			Str_pub_date: item.Published,
			Summary:      item.Description,
			//CreatedAt:    time.Now().Local().Add(time.Hour * time.Duration(5)),
			//UpdatedAt:    time.Now().Local().Add(time.Hour * time.Duration(5)),
		}
		old_post := post

		r := db.client.Where("link = ?", &post.Link).Limit(1).Find(&old_post)
		if r.RowsAffected == 0 {
			db.client.Create(&post)
		} else {
			if old_post.Str_pub_date != post.Str_pub_date {
				//str := fmt.Sprintf("%s updated to %s", post.Date, new_post.Date)
				history := model.History{
					Old_published_at: old_post.Publish_date,
					New_published_at: post.Publish_date,
					PostID:           old_post.ID,
					//CreatedAt:    time.Now().Local().Add(time.Hour * time.Duration(5)),
					//UpdatedAt:    time.Now().Local().Add(time.Hour * time.Duration(5)),
				}
				err := db.client.Model(&old_post).Association("HistoryList").Append(&history)
				if err != nil {
					db.logger.Warn(err)
				}
				if err := db.client.Model(model.Post{}).Where("id = ?", old_post.ID).Updates(map[string]interface{}{"publish_date": post.Publish_date, "str_pub_date": post.Str_pub_date}).Error; err != nil {
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
	category := parameters["category"].(string)
	// filter by field
	fil_title := parameters["fil_title"].(string)
	fil_link := parameters["fil_link"].(string)
	fil_publish_date := parameters["fil_publish_date"].(string)
	fil_summary := parameters["fil_summary"].(string)
	fil_createdAt := parameters["fil_createdAt"].(string)
	fil_updatedAt := parameters["fil_updatedAt"].(string)

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
		sortBy = "id.desc"
	}
	sortQuery, err := gormquery.ValidateAndReturnSortQuery(sortBy)
	if err != nil {
		db.logger.Warn(err)
	}
	new_offset := (offset - 1) * limit
	var posts []*model.Post
	if err := db.client.Where("category LIKE ? and title LIKE ? and link LIKE ? and publish_date LIKE ? and summary LIKE ? and created_at LIKE ? and updated_at LIKE ? ", "%"+category+"%", "%"+fil_title+"%", "%"+fil_link+"%", "%"+fil_publish_date+"%", "%"+fil_summary+"%", "%"+fil_createdAt+"%", "%"+fil_updatedAt+"%").Limit(limit).Offset(new_offset).Order(sortQuery).Preload("HistoryList").Find(&posts).Error; err != nil {
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
