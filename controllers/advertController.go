package controllers

import (
	"fmt"
	"net/http"
	"rest/di"
	"rest/repository"
	"rest/utils"
)

var rep *repository.AdvertRepository

func findRepository() {
	di.Container.Invoke(func(r *repository.AdvertRepository) {
		rep = r
	})
}

func ListAdverts(w http.ResponseWriter, r *http.Request) {
	findRepository()
	adverts, err := rep.FindAll()
	if err != nil {
		utils.Respond(w, utils.Message(false, fmt.Sprintf("error: %s", err)))
		return
	}
	resp := utils.Message(true, "success")
	resp["data"] = adverts
	utils.Respond(w, resp)
}

func ViewAdvert(w http.ResponseWriter, r *http.Request) {
	findRepository()
}

func CreateAdvert(w http.ResponseWriter, r *http.Request) {
	findRepository()
}
