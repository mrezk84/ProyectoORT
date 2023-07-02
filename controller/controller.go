package controller

import (
	"fmt"
	"net/http"
	"proyectoort/models"
	"proyectoort/views"
)

type TodoController struct {
	model *models.UserModel
	view  *views.TodoView
}



func NewController() *TodoController {
	return &TodoController{
		model: &models.UserModel{},
		view:  &views.TodoView{},
	}
}

func (c *TodoController) SetupRoutes() {
	http.HandleFunc("/", c.ListarTodosHandler)
	http.HandleFunc("/login", c.LoginHandler)
	http.HandleFunc("/dashboard", c.DashboardHandler)
}

func (c *TodoController) ListarTodosHandler(w http.ResponseWriter, r *http.Request) {
	c.view.MostrarError(w, r, fmt.Errorf("Acceso no autorizado"))
}

func (c *TodoController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	c.view.MostrarLogin(w, r)
}

func (c *TodoController) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	username := "Usuario"
	c.view.MostrarDashboard(w, r, username)
}
