package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"rss/internal/services/rss/service"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type RssController interface {
	Fetch()
	GetAll(w http.ResponseWriter, r *http.Request)
}

type rssController struct {
	service service.RssService
	logger  *logrus.Logger
}

func NewRssController(service service.RssService, logger *logrus.Logger) RssController {
	return &rssController{
		service: service,
		logger:  logger,
	}
}

func (c *rssController) Fetch() {
	err := godotenv.Load("../../.env")
	if err != nil {
		c.logger.Warn(err)
	}

	time_sleep := os.Getenv("TIME_SLEEP")
	i, err := strconv.Atoi(time_sleep)
	if err != nil {
		c.logger.Warn(err)
	}
	for {
		go c.service.FetchTurkmenPortal()
		go c.service.FetchOrient()
		time.Sleep(time.Duration(int(i)) * time.Minute)
	}
}

func (c *rssController) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := c.service.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(posts)
	}

}
