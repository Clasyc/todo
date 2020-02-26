package main

import (
	"fmt"
	"github.com/Clasyc/todo/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// homeHandler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "template/index.html")
}

// createHandler
func createHandler(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("title")

	if v == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	t := model.NewTask(v, false)
	err := t.Create()
	if err != nil {
		Error(w, err.Error(), 500)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// deleteHandler
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, _ := strconv.Atoi(id)

	t, err := model.GetTask(i)
	if err != nil {
		Error(w, err.Error(), 404)
	}

	err = t.Delete()
	if err != nil {
		Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// toggleHandler
func toggleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, _ := strconv.Atoi(id)

	t, err := model.GetTask(i)
	if err != nil {
		Error(w, err.Error(), 404)
	}

	err = t.Toggle(!t.Done)
	if err != nil {
		Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// headerHandler
func headerHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("ID").(string)

	fmt.Println("id inside handler: ", id)
	return
}