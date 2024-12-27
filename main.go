package main

import (
	"web/routes"
	"web/service"
)

func main() {
	service.InitDb()
	routes.InitRouter()
}
