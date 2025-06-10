package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venuyeredla/pan-services/internal/models"
)

type StatiContent struct {
	TemplateMap map[string]*template.Template
}

func (sc *StatiContent) Preprocess() {
	sc.TemplateMap = make(map[string]*template.Template)
	//sc.TemplateMap["todo"] = template.Must(template.ParseFiles("webapp/html/layout.html"))
	sc.TemplateMap["formtmpl"] = template.Must(template.ParseFiles("wstatic/form.html"))
}

func (sc *StatiContent) GenrateForm(c *gin.Context) {
	details := models.ContactInfo{
		Email: c.Request.FormValue("email"),
	}
	/*
		// Native way of loading template and generating htmls
		template.Must(template.ParseFiles("wstatic/form.html"))
		formtmpl := sc.TemplateMap["formtmpl"]
			 do something with details
		 formtmpl.Execute(w, d)
	*/
	_ = details

	data := models.TodoPageData{
		PageTitle: "My TODO list",
		Todos: []models.Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	d := struct {
		Success bool
		Todo    models.TodoPageData
	}{Success: false, Todo: data}

	c.HTML(http.StatusOK, "form.html", d)

}
