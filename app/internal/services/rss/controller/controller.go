package controller

import (
	"encoding/json"
	"net/http"
	"rss/internal/services/rss/dto"
	"rss/internal/services/rss/service"
	"rss/internal/utils/customvalidator"
	"rss/pkg/responses"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type RssController interface {
	// Fetch posts
	Fetch()
	// Get posts
	GetAllPosts(w http.ResponseWriter, r *http.Request)
	// Links
	LinkAll(w http.ResponseWriter, r *http.Request)
	LinkCreate(w http.ResponseWriter, r *http.Request)
	LinkDelete(w http.ResponseWriter, r *http.Request)
	LinkUpdate(w http.ResponseWriter, r *http.Request)
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

	// time_sleep := os.Getenv("TIME_SLEEP")
	// i, err := strconv.Atoi(time_sleep)
	// if err != nil {
	// 	c.logger.Warn(err)
	// }
	// for {
	// 	go c.service.FetchTurkmenPortal()
	// 	go c.service.FetchOrient()
	// 	time.Sleep(time.Duration(int(i)) * time.Minute)
	// }
	for {
		links, err := c.service.LinkAll()
		if err != nil {
			c.logger.Warn(err)
		}
		for i := 0; i < len(links); i++ {
			c.service.Fetch(links[i].Link)
		}
		time.Sleep(10 * time.Second)
	}
}

func (c *rssController) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := c.service.GetAllPosts()
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

// LINKS

func (c *rssController) LinkAll(w http.ResponseWriter, r *http.Request) {
	links, err := c.service.LinkAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: links}
		json.NewEncoder(w).Encode(response)
		return
	}
}

func (c *rssController) LinkCreate(w http.ResponseWriter, r *http.Request) {
	var linkCreateDTO dto.LinkDto
	err := json.NewDecoder(r.Body).Decode(&linkCreateDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	err = customvalidator.ValidateStruct(&linkCreateDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	link, err := c.service.LinkCreate(linkCreateDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: link}
	json.NewEncoder(w).Encode(response)
}

func (c *rssController) LinkDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	// id = int(userId)
	intVar, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// _, err = c.service.LinkFindByID(intVar)
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	response := responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: "Link with specified ID not found!"}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	err = c.service.LinkDelete(intVar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: "Link successfully deleted!"}
	json.NewEncoder(w).Encode(response)
}

func (c *rssController) LinkUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	// id = int(userId)
	intVar, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	var linkCreateDTO dto.LinkDto
	err = json.NewDecoder(r.Body).Decode(&linkCreateDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	err = customvalidator.ValidateStruct(&linkCreateDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	link, err := c.service.LinkUpdate(linkCreateDTO, intVar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: link}
	json.NewEncoder(w).Encode(response)
}
