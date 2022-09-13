package storage

import (
	"errors"
	"rss/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RssLinksStorage interface {
	// Links
	LinkAll() ([]*model.Link, error)
	LinkCreate(link *model.Link) (*model.Link, error)
	LinkDelete(link *model.Link) error
	LinkUpdate(link *model.Link) (*model.Link, error)
	LinkFindByID(link *model.Link, id int) (data *model.Link, err error)
}

type rssLinksStorage struct {
	client *gorm.DB
	logger *logrus.Logger
}

func NewRssLinksStorage(client *gorm.DB, logger *logrus.Logger) RssLinksStorage {
	return &rssLinksStorage{
		client: client,
		logger: logger,
	}
}

// LINKS

func (db *rssLinksStorage) LinkAll() ([]*model.Link, error) {
	var links []*model.Link
	if err := db.client.Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

func (db *rssLinksStorage) LinkCreate(link *model.Link) (*model.Link, error) {
	// if err := db.client.Create(link).Error; err != nil {
	// 	return nil, err
	// }
	// return link, nil
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	r := db.client.Where("source = ?", link.Source).Limit(1).Find(&link)
	if r.RowsAffected == 0 {
		if err := db.client.Create(link).Error; err != nil {
			return nil, err
		}
	} else {
		err1 := errors.New("there is already one")
		return nil, err1
	}
	return link, nil
}

func (db *rssLinksStorage) LinkFindByID(link *model.Link, id int) (data *model.Link, err error) {
	if err := db.client.First(link, id).Error; err != nil {
		return nil, err
	}
	return link, nil
}

func (db *rssLinksStorage) LinkDelete(link *model.Link) error {
	err := db.client.Delete(link).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *rssLinksStorage) LinkUpdate(role *model.Link) (data *model.Link, err error) {
	if err := db.client.Save(role).Error; err != nil {
		// fmt.Println(err.Error())
		return nil, err
	}
	return role, nil
}
