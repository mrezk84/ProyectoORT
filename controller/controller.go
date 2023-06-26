package controller

import (
	"fmt"
	"net/http"
	"proyectoort/model"
	"proyectoort/view"
)

type TodoController struct {
	model *model.UserModel
	view  *view.TodoView
}

func NewController() *TodoController {
	return &TodoController{
		model: &model.UserModel{},
		view:  &view.TodoView{},
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
