package main

import (
	"Daveslist/routes"
	"Daveslist/services/databases"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	databases.InitDatabase()
	// databases.MockRole()
	// databases.MockUser()
	r := routes.SetupRouter()
	// running
	r.Run(":8080")
}
