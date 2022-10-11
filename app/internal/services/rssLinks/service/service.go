package service

import (
	"rss/internal/model"
	"rss/internal/services/rssLinks/dto"
	"rss/internal/services/rssLinks/storage"
	"github.com/sirupsen/logrus"
)

type RssLinksService interface {
	// Links
	LinkAll() ([]*model.Link, error)
	LinkCreate(linkDto dto.LinkDto) (*model.Link, error)
	LinkDelete(id int) error
	LinkUpdate(linkDTO dto.LinkDto, id int) (data *model.Link, err error)
	LinkFindByID(id int) (*model.Link, error)
}

type rssLinksService struct {
	storage storage.RssLinksStorage
	logger  *logrus.Logger
}

func NewRssLinksService(storage storage.RssLinksStorage, logger *logrus.Logger) RssLinksService {
	return &rssLinksService{
		storage: storage,
		logger:  logger,
	}
}

// GET ALL LINKS
func (s *rssLinksService) LinkAll() ([]*model.Link, error) {
	links, err := s.storage.LinkAll()
	if err != nil {
		return nil, err
	} else {
		return links, nil
	}
}

// CREATE LINK
func (s *rssLinksService) LinkCreate(linkDto dto.LinkDto) (*model.Link, error) {
	link := &model.Link{
		Name:   linkDto.Name,
		Source: linkDto.Source,
		//CreatedAt: time.Now().Local().Add(time.Hour * time.Duration(5)),
		//UpdatedAt: time.Now().Local().Add(time.Hour * time.Duration(5)),
	}
	link, err := s.storage.LinkCreate(link)
	if err != nil {
		return nil, err
	}
	return link, nil
}

// LINK FIND BY ID
func (s *rssLinksService) LinkFindByID(id int) (data *model.Link, err error) {
	var link model.Link
	data, err = s.storage.LinkFindByID(&link, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// LINK DELETE
func (s *rssLinksService) LinkDelete(id int) error {
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
func (s *rssLinksService) LinkUpdate(linkDTO dto.LinkDto, id int) (data *model.Link, err error) {
	var oldLink model.Link
	_, err = s.storage.LinkFindByID(&oldLink, id)
	if err != nil {
		return nil, err
	}

	oldLink.Source = linkDTO.Source
	oldLink.Name = linkDTO.Name
	//oldLink.UpdatedAt = time.Now().Local().Add(time.Hour * time.Duration(5))
	data, err = s.storage.LinkUpdate(&oldLink)
	if err != nil {
		return nil, err
	}
	return data, nil
}
