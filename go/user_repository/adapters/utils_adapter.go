package adapters

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"

)

type Utilities struct{}

var Util *Utilities = &Utilities{}

func (e *Utilities) GenerateUUID() (string, error) {
	ID, err := uuid.NewUUID()
	return ID.String(), err
}

func (e *Utilities) EncryptPassword(s string) (string, error) {
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if error != nil { return "", error }
	return string(hashedPassword), nil
}

func (e *Utilities) VerifyPassword(password, encryptedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))

	if err != nil {
		return false, errors.New("PASSWORD_NOT_MATCH")
	}
	return true, nil
}
