package interfaces

type IUtil interface {
	EncryptPassword(password string) (string, error)
	VerifyPassword(password, encryptedPassword string) (bool, error)
	GenerateUUID() (string, error)
}
