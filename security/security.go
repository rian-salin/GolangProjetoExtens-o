package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash gera um hash seguro para a senha informada.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Verify compara uma senha em texto puro com um hash armazenado.
func Verify(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
