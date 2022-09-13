package controller

import (
	"encoding/json"
	"net/http"
	"rss/internal/services/rssLinks/dto"
	"rss/internal/services/rssLinks/service"
	"rss/internal/utils/customvalidator"
	"rss/pkg/responses"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type RssLinksController interface {
	// Links
	LinkAll(w http.ResponseWriter, r *http.Request)
	LinkCreate(w http.ResponseWriter, r *http.Request)
	LinkDelete(w http.ResponseWriter, r *http.Request)
	LinkUpdate(w http.ResponseWriter, r *http.Request)
}

type rssLinksController struct {
	service service.RssLinksService
	logger  *logrus.Logger
}

func NewRssLinksController(service service.RssLinksService, logger *logrus.Logger) RssLinksController {
	return &rssLinksController{
		service: service,
		logger:  logger,
	}
}

// LINKS

func (c *rssLinksController) LinkAll(w http.ResponseWriter, r *http.Request) {
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

func (c *rssLinksController) LinkCreate(w http.ResponseWriter, r *http.Request) {
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

func (c *rssLinksController) LinkDelete(w http.ResponseWriter, r *http.Request) {
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

func (c *rssLinksController) LinkUpdate(w http.ResponseWriter, r *http.Request) {
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
