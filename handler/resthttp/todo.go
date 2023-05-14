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

type todoHandler struct {
	todoService TodoService
}

func newTodoHandler(todoService TodoService) *todoHandler {
	return &todoHandler{
		todoService: todoService,
	}
}

func (p todoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	resp := newResponse()

	id := r.URL.Query().Get("activity_group_id")
	if id == "" {
		id = "0"
	}

	pid, err := strconv.Atoi(id)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	res, err := p.todoService.GetTodos(context.Background(), pid)
	if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(res, w)
	return
}

func (p todoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
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

	res, err := p.todoService.GetTodoByID(context.Background(), services.GetTodoByID{
		ID: aid,
	})
	if err != nil && strings.Contains(err.Error(), "sql: no rows in result set") {
		resp.setNotFound(fmt.Sprintf("Todo with ID %s Not Found", id), w)
		return
	} else if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(res, w)
	return
}

func (p todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	resp := newResponse()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	reqBody := services.CreateTodo{}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	if reqBody.Title == "" {
		resp.setBadRequest("title cannot be null", w)
		return
	}

	if reqBody.ActivityGroupID == 0 {
		resp.setBadRequest("activity group id cannot be null", w)
		return
	}

	res, err := p.todoService.CreateTodo(context.Background(), services.CreateTodo{
		Title:           reqBody.Title,
		ActivityGroupID: reqBody.ActivityGroupID,
		IsActive:        reqBody.IsActive,
	})
	if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setCreated(res, w)
	return
}

func (p todoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	resp := newResponse()

	type UpdateTodoReq struct {
		Title    string `json:"title"`
		Priority string `json:"priority"`
		IsActive bool   `json:"is_active"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.setBadRequest(err.Error(), w)
		return
	}

	reqBody := UpdateTodoReq{}
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

	if reqBody.Priority == "" {
		resp.setBadRequest("priority cannot be null", w)
		return
	}

	res, err := p.todoService.UpdateTodo(context.Background(), services.UpdateTodo{
		ID:       aid,
		Title:    reqBody.Title,
		Priority: reqBody.Priority,
		IsActive: reqBody.IsActive,
	})
	if err != nil && strings.Contains(err.Error(), "sql: no rows in result set") {
		resp.setNotFound(fmt.Sprintf("Todo with ID %s Not Found", id), w)
		return
	} else if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setCreated(res, w)
	return
}

func (p todoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
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

	_, err = p.todoService.DeleteTodo(context.Background(), services.DeleteTodo{
		ID: aid,
	})
	if err != nil && strings.Contains(err.Error(), "sql: no rows in result set") {
		resp.setNotFound(fmt.Sprintf("Todo with ID %s Not Found", id), w)
		return
	} else if err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(map[string]interface{}{}, w)
	return
}
