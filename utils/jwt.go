package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("SECRET_KOSKUY_TOKEN") // ubah jadi rahasia kamu sendiri nanti

// Struct untuk klaim token
type JWTClaim struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Fungsi untuk generate token
func GenerateJWT(userID uint, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JWTClaim{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Fungsi untuk validasi token
func ValidateToken(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
