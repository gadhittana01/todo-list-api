package resthttp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gadhittana01/todolist/services"
	"github.com/go-chi/chi"
)

type activityHandler struct {
	activityService ActivityService
}

func newActivityHandler(activityService ActivityService) *activityHandler {
	return &activityHandler{
		activityService: activityService,
	}
}

func (p activityHandler) GetActivities(w http.ResponseWriter, r *http.Request) {
	resp := newResponse()

	res, err := p.activityService.GetActivities(context.Background())
	if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(res, w)
	return
}

func (p activityHandler) GetActivityByID(w http.ResponseWriter, r *http.Request) {
	resp := newResponse()

	id := chi.URLParam(r, "id")
	if id == "" {
		resp.setBadRequest("Invalid Request Parameter", w)
		return
	}

	aid, err := strconv.Atoi(id)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	res, err := p.activityService.GetActivityByID(context.Background(), services.GetActivityByID{
		ID: aid,
	})
	if err != nil && strings.Contains(err.Error(), "sql: no rows in result set") {
		resp.setNotFound(fmt.Sprintf("Activity with ID %s Not Found", id), w)
		return
	} else if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(res, w)
	return
}

func (p activityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	resp := newResponse()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	reqBody := services.CreateActivity{}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	if reqBody.Title == "" {
		resp.setBadRequest("title cannot be null", w)
		return
	}

	if reqBody.Email == "" {
		resp.setBadRequest("email cannot be null", w)
		return
	}

	res, err := p.activityService.CreateActivity(context.Background(), services.CreateActivity{
		Title: reqBody.Title,
		Email: reqBody.Email,
	})
	if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setCreated(res, w)
	return
}

func (p activityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	resp := newResponse()

	type UpdateActivitiyReq struct {
		Title string `json:"title"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	reqBody := UpdateActivitiyReq{}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		resp.setBadRequest("Invalid Request Parameter", w)
		return
	}

	aid, err := strconv.Atoi(id)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	if reqBody.Title == "" {
		resp.setBadRequest("title cannot be null", w)
		return
	}

	res, err := p.activityService.UpdateActivity(context.Background(), services.UpdateActivity{
		ID:    aid,
		Title: reqBody.Title,
	})
	if err != nil && strings.Contains(err.Error(), "sql: no rows in result set") {
		resp.setNotFound(fmt.Sprintf("Activity with ID %s Not Found", id), w)
		return
	} else if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setCreated(res, w)
	return
}

func (p activityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	resp := newResponse()

	id := chi.URLParam(r, "id")
	if id == "" {
		resp.setBadRequest("Invalid Request Parameter", w)
		return
	}

	aid, err := strconv.Atoi(id)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	_, err = p.activityService.DeleteActivity(context.Background(), services.DeleteActivity{
		ID: aid,
	})
	if err != nil && strings.Contains(err.Error(), "sql: no rows in result set") {
		resp.setNotFound(fmt.Sprintf("Activity with ID %s Not Found", id), w)
		return
	} else if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(map[string]interface{}{}, w)
	return
}
