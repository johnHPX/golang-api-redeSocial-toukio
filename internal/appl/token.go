package appl

import (
	"API-RS-TOUKIO/configs"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken retorna um token assinado com as permissions de usuario
func CreateToken(userID int64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(configs.SecretKey))
}

// validateToken verifica se o token passado na requisição é valido
func validateToken(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, returnCheckKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalido")
}

// ExtractUsuarioID retorna o usuarioId que está salvo no token
func ExtractUserID(r *http.Request) (int64, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, returnCheckKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("token inválido")
}

// pega o valor do token
func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// Bearer asdlkdjsakl -> asdlkdjsakl

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// retorna a chave de verificação
func returnCheckKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return configs.SecretKey, nil
}
