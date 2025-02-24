package handler

import (
	// "context"
	// "encoding/json"
	"fmt"
	"log"
	// "errors"
	// "log"
	// "math/rand/v2"
	"net/http"
	"strconv"

	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository/query"
	// "github.com/devkaare/todo/views"
	"github.com/go-chi/chi/v5"
)

type User struct {
	Repo *query.PostgresRepo
}

func (u *User) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.Repo.GetUsers()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// jsonResp, _ := json.Marshal(users)
	// _, _ = w.Write(jsonResp)

	fmt.Fprintln(w, users)
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := &model.User{
		Email:    email,
		Password: password, // TODO: Hash pass
	}

	_, err := u.Repo.CreateUser(user) // TODO: Use returned userID for auth
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	URLParam := chi.URLParam(r, "ID")
	userID, err := strconv.Atoi(URLParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, _, err := u.Repo.GetUserByUserID(uint32(userID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, user)
}

// func (u *User) Create(w http.ResponseWriter, r *http.Request) {
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
// func (u *User) List(w http.ResponseWriter, r *http.Request) {
// 	todos, err := t.Repo.GetTodoList()
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	views.TodoForm(todos).Render(context.Background(), w)
// }
//
// func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
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
// func (u *User) UpdateByID(w http.ResponseWriter, r *http.Request) {
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
// func (u *User) DeleteByID(w http.ResponseWriter, r *http.Request) {
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
// func (u *User) EditByID(w http.ResponseWriter, r *http.Request) {
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
