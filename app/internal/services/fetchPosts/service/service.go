package service

import (
	"context"
	"rss/internal/model"
	"rss/internal/services/fetchPosts/storage"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
)

type FetchPostsService interface {
	Fetch(link *model.Link)
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

func (s *fetchPostsService) Fetch(link *model.Link) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(link.Source, ctx)
	if err != nil {
		s.logger.Warn(err)
	} else {
		s.storage.CreatePosts(feed.Items, link.Name)
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
