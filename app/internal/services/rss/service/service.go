package service

import (
	"rss/internal/model"
	"rss/internal/services/rss/dto"
	"rss/internal/services/rss/storage"
	"strings"

	"github.com/SlyMarbo/rss"
	"github.com/sirupsen/logrus"
)

type RssService interface {
	Fetch(link string)
	// FetchOrient()
	GetAllPosts() (data []*model.Post, err error)
	// Links
	LinkAll() ([]*model.Link, error)
	LinkCreate(linkDto dto.LinkDto) (*model.Link, error)
	LinkDelete(id int) error
	LinkUpdate(linkDTO dto.LinkDto, id int) (data *model.Link, err error)
	LinkFindByID(id int) (*model.Link, error)
}

type rssService struct {
	storage storage.RssStorage
	logger  *logrus.Logger
}

func NewRssService(storage storage.RssStorage, logger *logrus.Logger) RssService {
	return &rssService{
		storage: storage,
		logger:  logger,
	}
}

func (s *rssService) Fetch(link string) {
	feed, err := rss.Fetch(link)
	if err != nil {
		s.logger.Warn(err)
	} else {
		category := strings.Split(link, "/")
		s.storage.CreatePosts(feed.Items, category[2])
	}

}

func (s *rssService) GetAllPosts() (data []*model.Post, err error) {
	posts, err := s.storage.GetAll()
	if err != nil {
		return nil, err
	} else {
		return posts, nil
	}
}

// GET ALL LINKS
func (s *rssService) LinkAll() ([]*model.Link, error) {
	links, err := s.storage.LinkAll()
	if err != nil {
		return nil, err
	} else {
		return links, nil
	}
}

// CREATE LINK
func (s *rssService) LinkCreate(linkDto dto.LinkDto) (*model.Link, error) {
	link := &model.Link{
		Link: linkDto.Link,
	}
	link, err := s.storage.LinkCreate(link)
	if err != nil {
		return nil, err
	}
	return link, nil
}

// LINK FIND BY ID
func (s *rssService) LinkFindByID(id int) (data *model.Link, err error) {
	var link model.Link
	data, err = s.storage.LinkFindByID(&link, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// LINK DELETE
func (s *rssService) LinkDelete(id int) error {
	var link model.Link
	_, err := s.storage.LinkFindByID(&link, id)
	if err != nil {
		return err
	}
	err = s.storage.LinkDelete(&link)
	if err != nil {
		return err
	}
	return nil
}

// LINK UPDATE
func (s *rssService) LinkUpdate(linkDTO dto.LinkDto, id int) (data *model.Link, err error) {
	var oldLink model.Link
	_, err = s.storage.LinkFindByID(&oldLink, id)
	if err != nil {
		return nil, err
	}
	oldLink.Link = linkDTO.Link
	data, err = s.storage.LinkUpdate(&oldLink)
	if err != nil {
		return nil, err
	}
	return data, nil
}
