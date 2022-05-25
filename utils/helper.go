package utils

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"golang/config"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v4"
)

type LoginUser struct {
	Userid       string `json:"userId"`
	TokenLogin   string `json:"tokenLogin"`
}

func ParseToJson() (string, bool) {
	loginUser := LoginUser {
		Userid: "",
		TokenLogin: "",
	}

	jsonParse, jsonParseError := json.Marshal(loginUser)
	if jsonParseError != nil {
		return "", true
	}

	return string(jsonParse), false
}

func EncryptToken(expired int) (string, bool) {
	secretJWT := config.SecretJwt
	client := jwt.MapClaims {}
	client["exp"] = time.Now().Add(time.Second * time.Duration(expired)).Unix()
	encrypt := jwt.NewWithClaims(jwt.SigningMethodHS256, client)
	token, err := encrypt.SignedString([]byte(secretJWT))
	if err != nil {
		return "", true
	}
	return token, false
}

func DecryptToken(tokenStr string) (jwt.MapClaims, bool) {
	secretJWT := config.SecretJwt
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretJWT, nil
	})

	if err != nil {
		return nil, true
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, false
	} else {
		return nil, true
	}
}

func SHA256(text string) string {
	algorithm := crypto.SHA256.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
			if !unicode.IsSpace(ch) {
					b.WriteRune(ch)
			}
	}
	return b.String()
}

func IsValidEmail (email string) bool {
	var regex = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	
	return regex.MatchString(email)
}