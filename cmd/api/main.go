package main

import (
	"github.com/leMoreira/api_simples/cmd/api/bd"
	"github.com/leMoreira/api_simples/cmd/api/routes"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	bd.ConectaBaseDeDados()

	routes.Routes()

}
