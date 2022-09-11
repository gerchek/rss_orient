package storage

import (
	"rss/internal/model"

	"github.com/SlyMarbo/rss"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RssStorage interface {
	CreatePosts(items []*rss.Item, category string)
	GetAll() (data []*model.Post, err error)
	// Links
	LinkAll() ([]*model.Link, error)
	LinkCreate(link *model.Link) (*model.Link, error)
	LinkDelete(link *model.Link) error
	LinkUpdate(link *model.Link) (*model.Link, error)
	LinkFindByID(link *model.Link, id int) (data *model.Link, err error)
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

func (db *rssStorage) CreatePosts(items []*rss.Item, category string) {
	for _, item := range items {
		post := model.Post{
			Category: category,
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

// LINKS

func (db *rssStorage) LinkAll() ([]*model.Link, error) {
	var links []*model.Link
	if err := db.client.Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

func (db *rssStorage) LinkCreate(link *model.Link) (*model.Link, error) {
	if err := db.client.Create(link).Error; err != nil {
		return nil, err
	}
	return link, nil
}

func (db *rssStorage) LinkFindByID(link *model.Link, id int) (data *model.Link, err error) {
	if err := db.client.First(link, id).Error; err != nil {
		return nil, err
	}
	return link, nil
}

func (db *rssStorage) LinkDelete(link *model.Link) error {
	err := db.client.Delete(link).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *rssStorage) LinkUpdate(role *model.Link) (data *model.Link, err error) {
	if err := db.client.Save(role).Error; err != nil {
		// fmt.Println(err.Error())
		return nil, err
	}
	return role, nil
}
