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
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["usuariosId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(configs.SecretKey))
}

// validateToken verifica se o token passado na requisição é valido
func validateToken(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, erro := jwt.Parse(tokenString, returnCheckKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalido")
}

// ExtractUsuarioID retorna o usuarioId que está salvo no token
func ExtractUsuarioID(r *http.Request) (uint64, error) {
	tokenString := ExtractToken(r)
	fmt.Println(tokenString)
	token, erro := jwt.Parse(tokenString, returnCheckKey)
	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["usuariosId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return userID, nil
	}

	return 0, errors.New("Token inválido")
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// Bearer asdlkdjsakl

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnCheckKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return configs.SecretKey, nil
}