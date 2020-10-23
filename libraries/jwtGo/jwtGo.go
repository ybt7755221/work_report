package jwtGo

import (
	"fmt"
	"time"
	"work_report/config"

	"github.com/dgrijalva/jwt-Go"
)

var jwtSecret = []byte(config.GetApolloString("JWTSECRET", "lh$6$@I#7b&gEk06M3"))

type Claims struct {
	Masterid   string `json:"masterid"`
	MasterName string `json:"master_name"`
	Mobile     string `json:"mobile"`
	FullNname  string `json:"full_name"`
	System     string `json:"system"`
	jwt.StandardClaims
}

func GenerateToken(params map[string]string) (string, error) {
	claims := Claims{
		params["masterid"],
		params["master_name"],
		params["mobile"],
		params["full_name"],
		params["system"],
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // 过期时间，必须设置
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	fmt.Println(jwtSecret)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
