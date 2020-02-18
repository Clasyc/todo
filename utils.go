package main

import (
	"github.com/Clasyc/todo/model"
	"html/template"
	"net/http"
)

type Response struct {
	Tasks []*model.Task
}

func render(w http.ResponseWriter, name string) {
	tmpl, _ := template.ParseFiles(name)

	tasks, err := model.GetTasks()
	if err != nil {
		Error(w, err.Error(), 400)
		return
	}

	err = tmpl.Execute(w, Response{tasks})
	if err != nil {
		Error(w, err.Error(), 500)
		return
	}
}