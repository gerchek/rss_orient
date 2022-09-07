package service

import (
	"rss/internal/model"
	"rss/internal/services/rss/storage"

	"github.com/SlyMarbo/rss"
	"github.com/sirupsen/logrus"
)

type RssService interface {
	FetchTurkmenPortal()
	FetchOrient()
	GetAll() (data []*model.Post, err error)
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

func (s *rssService) FetchTurkmenPortal() {
	feed, err := rss.Fetch("https://turkmenportal.com/rss")
	if err != nil {
		s.logger.Warn(err)
	}
	s.storage.CreateTurkmenPortal(feed.Items)
}

func (s *rssService) FetchOrient() {
	feed, err := rss.Fetch("https://orient.tm/ru/rss")
	if err != nil {
		s.logger.Warn(err)
	}
	s.storage.CreateOrient(feed.Items)
}

func (s *rssService) GetAll() (data []*model.Post, err error) {
	posts, err := s.storage.GetAll()
	if err != nil {
		return nil, err
	} else {
		return posts, nil
	}
}
