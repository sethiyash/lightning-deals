package main

import (
	"allen/deals/database"
	"allen/deals/routes"
	"log"
	"net/http"
)

func main() {
	_, db := database.ConnectDB()
	//defer client.Disconnect(nil)

	router := routes.SetupRouts(db)

	log.Fatal(http.ListenAndServe(":8080", router))

}
