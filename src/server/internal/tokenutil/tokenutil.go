package tokenutil

import (
	"fmt"
	"time"

	customers "earnforglance/server/domain/customers"
	domain "earnforglance/server/domain/security"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *customers.Customer, slugs []domain.UrlRecord, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &domain.JwtCustomClaims{
		Name: user.FirstName,
		ID:   user.ID.Hex(),
		Slug: slugs,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *customers.Customer, slugs []domain.UrlRecord, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID: user.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["id"].(string), nil
}

func ExtractSlugsFromToken(requestToken string, secret string) ([]domain.UrlRecord, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	fmt.Println("claims", claims)
	slugsRaw, ok := claims["slug"]
	if !ok {
		return nil, fmt.Errorf("slugs not found in token")
	}

	slugsIface, ok := slugsRaw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("slugs claim is not an array")
	}

	slugs := make([]domain.UrlRecord, len(slugsIface))
	for i, v := range slugsIface {
		fmt.Println("v", v)
		slugs[i], ok = v.(domain.UrlRecord)
		if !ok {
			return nil, fmt.Errorf("slugs claim contains non-string value")
		}
	}

	fmt.Println("claims", claims)
	return slugs, nil
}
