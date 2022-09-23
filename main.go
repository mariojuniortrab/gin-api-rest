package main

import (
	"github.com/mariojuniortrab/gin-api-rest/database"
	"github.com/mariojuniortrab/gin-api-rest/routes"
)

func main() {
	database.DatabaseConnect()
	routes.HandleRequests()
}
