package views

import (
	"fmt"
	"net/http"
)

type TodoView struct {
}

func (v *TodoView) MostrarError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Fprintf(w, "Error: %v", err)
}

func (v *TodoView) MostrarLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mostrar formulario de login")
}

func (v *TodoView) MostrarDashboard(w http.ResponseWriter, r *http.Request, username string) {
	fmt.Fprintf(w, "Dashboard de %s", username)
}