package main

import (
	"api-rest/internal/database"
	"api-rest/internal/routers"
	"fmt"
)

func main() {
	fmt.Println("Initializing Application")
	database.Connect()
	routers.HandleRequest()
}