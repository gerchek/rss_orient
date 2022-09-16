package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"rss/internal/services/fetchPosts/service"
	"rss/pkg/responses"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type FetchPostsController interface {
	// Fetch posts
	Fetch()
	// Get posts
	GetAllPosts(w http.ResponseWriter, r *http.Request)
}

type fetchPostsController struct {
	service service.FetchPostsService
	logger  *logrus.Logger
}

func NewFetchPostsController(service service.FetchPostsService, logger *logrus.Logger) FetchPostsController {
	return &fetchPostsController{
		service: service,
		logger:  logger,
	}
}

func (c *fetchPostsController) Fetch() {
	err := godotenv.Load("../../.env")
	if err != nil {
		c.logger.Warn(err)
	}

	time_sleep := os.Getenv("TIME_SLEEP")
	i, err := strconv.Atoi(time_sleep)
	if err != nil {
		c.logger.Warn(err)
	}
	// for {
	// 	go c.service.FetchTurkmenPortal()
	// 	go c.service.FetchOrient()
	// time.Sleep(time.Duration(int(i)) * time.Minute)
	// }
	for {
		links, err := c.service.LinkAll()
		if err != nil {
			c.logger.Warn(err)
		}
		for i := 0; i < len(links); i++ {
			c.service.Fetch(links[i])
		}
		time.Sleep(time.Duration(int(i)) * time.Minute)
	}
}

func (c *fetchPostsController) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sortBy")
	strLimit := r.URL.Query().Get("strLimit")
	strOffset := r.URL.Query().Get("strOffset")
	category := r.URL.Query().Get("category")
	fil_title := r.URL.Query().Get("fil_title")
	fil_link := r.URL.Query().Get("fil_link")
	fil_publish_date := r.URL.Query().Get("fil_publish_date")
	fil_summary := r.URL.Query().Get("fil_summary")
	fil_createdAt := r.URL.Query().Get("fil_createdAt")
	fil_updatedAt := r.URL.Query().Get("fil_updatedAt")

	parameters := map[string]interface{}{
		"sortBy":    sortBy,
		"strLimit":  strLimit,
		"strOffset": strOffset,
		"category":  category,
		// filter by field
		"fil_title":        fil_title,
		"fil_link":         fil_link,
		"fil_publish_date": fil_publish_date,
		"fil_summary":      fil_summary,
		"fil_createdAt":    fil_createdAt,
		"fil_updatedAt":    fil_updatedAt,
	}

	posts, err := c.service.GetAllPosts(parameters)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusInternalServerError)
		response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: posts}
		json.NewEncoder(w).Encode(response)
		return
	}
}
