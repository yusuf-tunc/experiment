package main

import (
	"errors"
	"some_services/user_repository/interfaces"
	"some_services/user_repository/structs"
)

type UserRepository struct {
	persistence interfaces.IPersistence
	util interfaces.IUtil
}

func (UR *UserRepository) GetUsers() ([]*structs.User, error) {
	return UR.persistence.GetUsers()
}

func (UR *UserRepository) CreateUser(name, username, password, role string) (string, error) {
	usernameTaken := UR.persistence.IsUsernameTaken(username)
	if usernameTaken { return "", errors.New("USERNAME_TAKEN") }

	encrypted, err := UR.util.EncryptPassword(password)
	if err != nil { return "", err }
	
	ID, err := UR.util.GenerateUUID()
	if err != nil { return "", err }

	return UR.persistence.CreateUser(ID, name, username, encrypted, role)
}

func (UR *UserRepository) GetUser(ID string) (*structs.User, error) {
	return UR.persistence.GetUser(ID)
}

func (UR *UserRepository) DeleteUser(ID string) error {
	return UR.persistence.DeleteUser(ID)
}

func (UR *UserRepository) UpdatePassword(ID, password string) error {
	encrypted, err := UR.util.EncryptPassword(password)
	if err != nil { return err }

	return UR.persistence.UpdatePassword(ID, encrypted)
}

func (UR *UserRepository) VerifyPassword(password, encryptedPassword string) (bool, error) {
	return UR.util.VerifyPassword(password, encryptedPassword)
}

func (UR *UserRepository) RemoveAllUsers() {
	UR.persistence.RemoveAllUsers()
}
