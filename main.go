package main

import (
	"fmt"

	"github.com/eduzera0/api-rest-go.git/database"
	"github.com/eduzera0/api-rest-go.git/models"
	"github.com/eduzera0/api-rest-go.git/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
