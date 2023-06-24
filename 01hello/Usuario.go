package main

import "fmt"

func main() {
	fmt.Println("structs in goland")
	usuario := Usuario{1, "santi", 51679947, Rol{1, "constructor"}}
	fmt.Println(usuario)
}

type Usuario struct {
	id     int
	nombre string
	cedula int
	rol    Rol
}
