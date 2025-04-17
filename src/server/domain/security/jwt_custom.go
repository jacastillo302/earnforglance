package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type UrlRecord struct {
	Name     string `json:"name"`
	IsRead   bool   `json:"is_read"`
	IsWrite  bool   `json:"is_write"`
	IsDelete bool   `json:"is_delete"`
	IsUpdate bool   `json:"is_update"`
}

type JwtCustomClaims struct {
	Name string      `json:"name"`
	ID   string      `json:"id"`
	Slug []UrlRecord `json:"slug"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
