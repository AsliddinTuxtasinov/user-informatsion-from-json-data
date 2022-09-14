package dataloaders

import (
	"encoding/json"
	"user-informations-system/models"
	"user-informations-system/utils"
)

// Load json data to struct model
func LoadUsers() []models.UserModel {
	bytes, _ := utils.ReadFile("./json/users.json")
	var data []models.UserModel
	json.Unmarshal(bytes, &data)
	return data
}

// Load json data to struct model
func LoadInterests() []models.InterestsModel {
	bytes, _ := utils.ReadFile("./json/interests.json")
	var data []models.InterestsModel
	json.Unmarshal(bytes, &data)
	return data
}

// Load json data to struct model
func LoadInterestMappingModel() []models.InterestMappingModel {
	bytes, _ := utils.ReadFile("./json/userInterestMapping.json")
	var data []models.InterestMappingModel
	json.Unmarshal(bytes, &data)
	return data
}
