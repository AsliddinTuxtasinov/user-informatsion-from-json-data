package models

type UserModel struct {
	ID        int
	FirstName string
	LastName  string
	Interests []InterestsModel
	Profile   string
}
