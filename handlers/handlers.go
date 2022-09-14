package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-informations-system/dataloaders"
	"user-informations-system/models"
	modelsVm "user-informations-system/models/vm"
)

func Run() {
	http.HandleFunc("/", Home)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Page Object
	page := modelsVm.PageModelView{ID: 1, Name: "Users", Description: "Users List", URI: r.RequestURI}

	// Data Loader
	users := dataloaders.LoadUsers()
	interests := dataloaders.LoadInterests()
	interestMappings := dataloaders.LoadInterestMappingModel()

	// Response data preparation
	var newUsers []models.UserModel
	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, ininterest := range interests {
					if ininterest.ID == interestMapping.InterestID {
						user.Interests = append(user.Interests, ininterest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}
	viewModel := modelsVm.UserModelView{Page: page, Users: newUsers}
	bytes, _ := json.Marshal(viewModel)

	// Response data
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
