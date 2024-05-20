package controllers

import (
	"allen/deals/models"
	"allen/deals/utils"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DealController struct {
	DB *mongo.Database
}

func (dc *DealController) CreateDeal(w http.ResponseWriter, r *http.Request) {
	var deal models.Deal
	if err := json.NewDecoder(r.Body).Decode(&deal); err != nil {
		log.Fatal(err)
		return
	}

	deal.Id = primitive.NewObjectID().Hex()

	if err := deal.CreateDeal(dc.DB); err != nil {
		log.Fatal(err)
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, deal)

}

func (dc *DealController) UpdateDeal(w http.ResponseWriter, r *http.Request) {
	var deal models.Deal
	if err := json.NewDecoder(r.Body).Decode(&deal); err != nil {
		log.Fatal(err)
		return
	}

	deal.Id = primitive.NewObjectID().Hex()

	if err := deal.UpdateDeal(dc.DB); err != nil {
		log.Fatal(err)
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, deal)
}

func (dc *DealController) EndDeal(w http.ResponseWriter, r *http.Request) {
	var deal models.Deal
	if err := json.NewDecoder(r.Body).Decode(&deal); err != nil {
		log.Fatal(err)
		return
	}
	
	if err := deal.EndDeal(dc.DB); err != nil {
		log.Fatal(err)
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, deal)

}

func (dc *DealController) ClaimDeal(w http.ResponseWriter, r *http.Request) {
	var deal models.Deal
	if err := json.NewDecoder(r.Body).Decode(&deal); err != nil {
		log.Fatal(err)
		return
	}
	
	if err := deal.ClaimDeal(dc.DB, "userID"); err != nil {
		log.Fatal(err)
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, deal)

}


