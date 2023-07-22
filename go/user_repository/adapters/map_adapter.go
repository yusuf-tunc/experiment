package adapters

import (
	"errors"
	"some_services/user_repository/structs"
	
)

type Database struct { Users map[string]*structs.User }

var Users = make(map[string]*structs.User)

var Persistence = &Database{Users}

func (adapter *Database) GetUsers() ([]*structs.User, error) {
	var result []*structs.User

	for _, user := range Users {
		result = append(result, user)
	}
	return result, nil
}

func (adapter *Database) CreateUser(ID, name, username, password, role string) (string, error) {
	user := &structs.User{ID: ID, Name: name, Username: username, Password: password, Role: role}
	Users[user.ID] = user
	return user.ID, nil
}

func (adapter *Database) GetUser(ID string) (*structs.User, error) {
	if !isUserExists(ID) { return nil, errors.New("USER_NOT_FOUND") }

	return Users[ID], nil
}

func (adapter *Database) DeleteUser(ID string) error {
	if !isUserExists(ID) { return errors.New("USER_NOT_FOUND") }

	delete(Users, ID)
	return nil
}

func (adapter *Database) UpdatePassword(ID, password string) error {
	user := Users[ID]
	user.Password = password
	return nil
}

func (adapter *Database) RemoveAllUsers() {
	Users = make(map[string]*structs.User)
}

func (adapter *Database) IsUsernameTaken(username string) bool {
	for _, user := range Users { 
		if user.Username == username { return true } 
	}
	return false
}

func isUserExists(ID string) bool {
	return Users[ID] != nil
}
