package handler

//
// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"log"
// 	"math/rand/v2"
// 	"net/http"
// 	"strconv"
//
// 	"github.com/devkaare/todo/model"
// 	"github.com/devkaare/todo/repository/todo"
// 	"github.com/devkaare/todo/views"
//
// 	"github.com/go-chi/chi/v5"
// )
//
// type Todo struct {
// 	Repo *todo.PostgresRepo
// }
//
// func (t *Todo) Health(w http.ResponseWriter, r *http.Request) {
// 	jsonResp, _ := json.Marshal(t.Repo.Health())
// 	_, _ = w.Write(jsonResp)
// }
//
// func (t *Todo) Create(w http.ResponseWriter, r *http.Request) {
// 	todo := &model.Todo{
// 		ID:          rand.Uint32N(2147483647),
// 		Title:       r.FormValue("title"),
// 		Description: r.FormValue("description"),
// 	}
// 	if _, err := t.Repo.GetTodoByID(todo.ID); err == errors.New("todo not found") && err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	if err := t.Repo.CreateTodo(todo); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	views.TodoPost(todo).Render(context.Background(), w)
// }
//
// func (t *Todo) List(w http.ResponseWriter, r *http.Request) {
// 	todos, err := t.Repo.GetTodoList()
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	views.TodoForm(todos).Render(context.Background(), w)
// }
//
// func (t *Todo) GetByID(w http.ResponseWriter, r *http.Request) {
// 	URLParam := chi.URLParam(r, "ID")
// 	id, err := strconv.Atoi(URLParam)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	todo, err := t.Repo.GetTodoByID(uint32(id))
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	views.TodoByIDForm(todo).Render(context.Background(), w)
// }
//
// func (t *Todo) UpdateByID(w http.ResponseWriter, r *http.Request) {
// 	URLParam := chi.URLParam(r, "ID")
// 	id, err := strconv.Atoi(URLParam)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	todo := &model.Todo{
// 		ID:          uint32(id),
// 		Title:       r.FormValue("title"),
// 		Description: r.FormValue("description"),
// 	}
//
// 	if err := t.Repo.UpdateTodoByID(todo); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	w.Write([]byte("<p>Successfully updated todo!</p>"))
// }
//
// func (t *Todo) DeleteByID(w http.ResponseWriter, r *http.Request) {
// 	URLParam := chi.URLParam(r, "ID")
// 	id, err := strconv.Atoi(URLParam)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	if err := t.Repo.DeleteTodoByID(uint32(id)); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	w.Write([]byte("<p>Successfully deleted todo!</p>"))
// }
//
// func (t *Todo) EditByID(w http.ResponseWriter, r *http.Request) {
// 	URLParam := chi.URLParam(r, "ID")
// 	id, err := strconv.Atoi(URLParam)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	todo, err := t.Repo.GetTodoByID(uint32(id))
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	views.TodoByIDPost(todo).Render(context.Background(), w)
// }
