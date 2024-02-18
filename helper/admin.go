package helper

import (
	"fmt"
	"time"

	"github.com/Adilfarooque/Footgo_Ecommerce/config"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/golang-jwt/jwt"
)

type authCustomClaimsAdmin struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	cfg, _ := config.LoadConfig()
	claims := &authCustomClaimsAdmin{
		Firstname: admin.Fristname,
		Lastname:  admin.Lastname,
		Email:     admin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString([]byte(cfg.KEY_ADMIN))
	if err != nil {
		fmt.Println("Error is", err)
		return "", err
	}
	return tokenString, nil
}
