package main

import (
	"fmt"

	"github.com/rafaellimasnp/api_go_rest/models"
	"github.com/rafaellimasnp/api_go_rest/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
