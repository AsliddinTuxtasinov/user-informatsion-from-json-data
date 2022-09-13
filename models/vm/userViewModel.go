package models

import "user-informations-system/models"

type UserModelView struct {
	Page  PageModelView
	Users []models.UserModel
}
