package service

import (
	"rss/internal/model"
	"rss/internal/services/fetchPosts/storage"
	"strings"

	"github.com/SlyMarbo/rss"
	"github.com/sirupsen/logrus"
)

type FetchPostsService interface {
	Fetch(link string)
	// FetchOrient()
	GetAllPosts(parameters map[string]interface{}) (data []*model.Post, err error)
	// ALL LINKS
	LinkAll() ([]*model.Link, error)
}

type fetchPostsService struct {
	storage storage.FetchPostsStorage
	logger  *logrus.Logger
}

func NewFetchPostsService(storage storage.FetchPostsStorage, logger *logrus.Logger) FetchPostsService {
	return &fetchPostsService{
		storage: storage,
		logger:  logger,
	}
}

func (s *fetchPostsService) Fetch(link string) {
	feed, err := rss.Fetch(link)
	if err != nil {
		s.logger.Warn(err)
	} else {
		category := strings.Split(link, "/")
		s.storage.CreatePosts(feed.Items, category[2])
	}

}

func (s *fetchPostsService) GetAllPosts(parameters map[string]interface{}) (data []*model.Post, err error) {
	posts, err := s.storage.GetAll(parameters)
	if err != nil {
		return nil, err
	} else {
		return posts, nil
	}
}

// GET ALL LINKS
func (s *fetchPostsService) LinkAll() ([]*model.Link, error) {
	links, err := s.storage.LinkAll()
	if err != nil {
		return nil, err
	} else {
		return links, nil
	}
}
