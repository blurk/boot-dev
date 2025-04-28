package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPasswordHash(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}

	return nil
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	if userID == uuid.Nil {
		return "", errors.New("userID is empty")
	}

	if len(tokenSecret) == 0 {
		return "", errors.New("secret is empty")
	}

	issuedAt := time.Now().UTC()
	expiresAt := issuedAt.Add(expiresIn)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		Subject:   userID.String(),
	})

	signedString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse JWT: %w", err)
	}

	if !token.Valid {
		return uuid.Nil, errors.New("Invalid JWT")
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get subject from claims: %w", err)
	}

	userId, err := uuid.Parse(subject)
	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	authorization := headers.Get("Authorization")

	if len(authorization) == 0 {
		return "", fmt.Errorf("authorization headers is missing")
	}

	parts := strings.SplitN(authorization, " ", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("malformed authorization header")
	}

	scheme := strings.ToLower(parts[0])
	if scheme != "bearer" {
		return "", fmt.Errorf("incorrect authorization scheme: expected Bearer, got %s", parts[0])
	}

	return parts[1], nil
}

func MakeRefreshToken() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)

	if err != nil {
		return "", nil
	}

	token := hex.EncodeToString(key)

	return token, nil
}

func GetAPIKey(headers http.Header) (string, error) {
	authorization := headers.Get("Authorization")

	if len(authorization) == 0 {
		return "", fmt.Errorf("authorization headers is missing")
	}

	parts := strings.SplitN(authorization, " ", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("malformed authorization header")
	}

	scheme := strings.ToLower(parts[0])
	if scheme != "apikey" {
		return "", fmt.Errorf("incorrect authorization scheme: expected ApiKey, got %s", parts[0])
	}

	return parts[1], nil
}
