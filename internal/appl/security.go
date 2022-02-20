package appl

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPassword virificar se a senha passada é igual a da salva no banco de dados
func CheckPassword(passwordWithHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(senhaString))
}
