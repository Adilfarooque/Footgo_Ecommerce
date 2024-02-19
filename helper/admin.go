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


func ValidateToken(tokenString string)(*authCustomClaimsAdmin,error){
	cfg,_ := config.LoadConfig()
	token,err := jwt.ParseWithClaims(tokenString,&authCustomClaimsAdmin{},func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("unexpected signing method :%v",token.Header["alg"])
		}
		return []byte(cfg.KEY_ADMIN),nil
	})
	if err != nil{
		return nil,err
	}
	if claims, ok := token.Claims.(*authCustomClaimsAdmin); ok && token.Valid{
		return claims,nil
	}
	return nil,fmt.Errorf("invalid token")
}