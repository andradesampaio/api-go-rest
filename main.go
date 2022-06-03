package main

import (
	"fmt"

	"github.com/guilhermeonrails/go-rest-api/database"
	"github.com/guilhermeonrails/go-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
