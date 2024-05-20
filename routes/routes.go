package routes

import (
	"allen/deals/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouts(db *mongo.Database) *mux.Router {
	router := mux.NewRouter()

	dealController := &controllers.DealController{DB: db}

	router.HandleFunc("/createDeal", dealController.CreateDeal).Methods("POST")
	router.HandleFunc("/updateDeal", dealController.UpdateDeal).Methods("PUT")
	router.HandleFunc("/endDeal", dealController.UpdateDeal).Methods("")
	router.HandleFunc("/claimDeal", dealController.ClaimDeal).Methods("GET")
	return router
}
