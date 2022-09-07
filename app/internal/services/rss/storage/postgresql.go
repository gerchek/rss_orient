package storage

import (
	"rss/internal/model"

	"github.com/SlyMarbo/rss"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RssStorage interface {
	CreateTurkmenPortal(items []*rss.Item)
	CreateOrient(items []*rss.Item)
	GetAll() (data []*model.Post, err error)
}

type rssStorage struct {
	client *gorm.DB
	logger *logrus.Logger
}

func NewRssStorage(client *gorm.DB, logger *logrus.Logger) RssStorage {
	return &rssStorage{
		client: client,
		logger: logger,
	}
}

func (db *rssStorage) CreateTurkmenPortal(items []*rss.Item) {
	for _, item := range items {
		post := model.Post{
			Category: "TurkmenPortal",
			Title:    item.Title,
			Link:     item.Link,
			Date:     item.Date,
			Summary:  item.Summary,
		}
		db.client.Create(&post)
	}
}

func (db *rssStorage) CreateOrient(items []*rss.Item) {
	for _, item := range items {
		post := model.Post{
			Category: "Orient",
			Title:    item.Title,
			Link:     item.Link,
			Date:     item.Date,
			Summary:  item.Summary,
		}
		db.client.Create(&post)
	}
}

func (db *rssStorage) GetAll() (data []*model.Post, err error) {
	var posts []*model.Post
	if err := db.client.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
