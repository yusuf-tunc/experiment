package interfaces

import ("some_services/user_repository/structs")

type IPersistence interface {
	GetUsers() ([]*structs.User, error)
	CreateUser(ID, name, username, password, role string) (string, error)
	GetUser(ID string) (*structs.User, error)
	DeleteUser(ID string) error
	UpdatePassword(ID, password string) error
	IsUsernameTaken(username string) bool
	RemoveAllUsers()
}