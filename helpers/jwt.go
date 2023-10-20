package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/Ma1y0/gist-clone/model"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user model.UserModel) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// claims["authorized"] = true
	claims["user_id"] = user.ID

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT2(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}

		return os.Getenv("JWT_SECRET"), nil
	})
	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	} else {
		return false, nil
	}
}

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error parsing token: %v", err)
	}

	// Validate the token
	if !token.Valid {
		return nil, fmt.Errorf("Token is not valid")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Error extracting claims")
	}

	return claims, nil
}

func ExtractIdFromJWT(jwtToken any) (string, error) {
	jwtMap, ok := jwtToken.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("JWT claims aren't a map")
	}

	userIDVal, ok := jwtMap["user_id"]
	if !ok {
		return "", fmt.Errorf("user_id not found in the map")
	}

	userID, ok := userIDVal.(string)
	if !ok {
		return "", fmt.Errorf("user_id value is not of type string")
	}

	return userID, nil
}
