package main

import (
	"log"

	"github.com/nestorraulcgomez/Twittor/bd"
	"github.com/nestorraulcgomez/Twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
