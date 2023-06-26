package main

import (
	"fmt"
	"net/http"
	"proyectoort/controller"
)

func main() {
	controller := controller.NewController()
	controller.SetupRoutes()

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}