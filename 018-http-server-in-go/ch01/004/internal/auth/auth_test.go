package auth

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestHashPassword(t *testing.T) {
	testCases := []struct {
		password      string
		expectedError bool
	}{
		{password: "validPassword", expectedError: false},
		{password: "", expectedError: false},
		{password: "verylongpassword123456789012345678901234567890", expectedError: false},
	}

	for _, tc := range testCases {
		hashed, err := HashPassword(tc.password)
		if tc.expectedError {
			if err == nil {
				t.Errorf("HashPassword(%q) expected error, got nil", tc.password)
			}
		} else {
			if err != nil {
				t.Errorf("HashPassword(%q) returned unexpected error: %v", tc.password, err)
			}
			if hashed == tc.password {
				t.Errorf("HashPassword(%q) returned the original password", tc.password)
			}
		}
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "test-password"
	hashed, _ := HashPassword(password) //error ignored for test

	testCases := []struct {
		hash          string
		password      string
		expectedError bool
	}{
		{hash: hashed, password: password, expectedError: false},
		{hash: hashed, password: "wrong-password", expectedError: true},
		{hash: "invalid-hash", password: password, expectedError: true},
	}

	for _, tc := range testCases {
		err := CheckPasswordHash(tc.hash, tc.password)
		if tc.expectedError {
			if err == nil {
				t.Errorf("CheckPasswordHash(%q, %q) expected error, got nil", tc.hash, tc.password)
			}
		} else {
			if err != nil {
				t.Errorf("CheckPasswordHash(%q, %q) returned unexpected error: %v", tc.hash, tc.password, err)
			}
		}
	}
}

func TestMakeJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "test-secret"
	expiresIn := time.Hour

	t.Run("Successful JWT Creation", func(t *testing.T) {
		tokenString, err := MakeJWT(userID, tokenSecret, expiresIn)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if tokenString == "" {
			t.Fatal("Expected non-empty token string, got empty")
		}

		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		})
		if err != nil {
			t.Fatalf("Expected valid token, got error: %v", err)
		}
		if !token.Valid {
			t.Fatal("Expected token to be valid")
		}

		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			t.Fatal("Expected RegisteredClaims type")
		}
		if claims.Subject != userID.String() {
			t.Errorf("Expected subject %s, got %s", userID.String(), claims.Subject)
		}
		if claims.Issuer != "chirpy" {
			t.Errorf("Expected issuer 'chirpy', got %s", claims.Issuer)
		}
	})

	t.Run("Invalid Token Secret", func(t *testing.T) {
		invalidSecret := "" // Empty secret should cause signing error
		tokenString, err := MakeJWT(userID, invalidSecret, expiresIn)
		if err == nil {
			t.Fatal("Expected error with invalid secret, got none")
		}
		if tokenString != "" {
			t.Errorf("Expected empty token string, got %s", tokenString)
		}
	})

	t.Run("Token Expiration Validation", func(t *testing.T) {
		shortExpiresIn := time.Second
		tokenString, err := MakeJWT(userID, tokenSecret, shortExpiresIn)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Wait for token to expire
		time.Sleep(2 * time.Second)

		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		})
		if err == nil || token.Valid {
			t.Fatal("Expected expired token to be invalid")
		}
	})
}

func TestValidateJWT(t *testing.T) {
	// Test Case 1: Valid Token
	t.Run("Valid Token", func(t *testing.T) {
		userID := uuid.New()
		secret := "test-secret"
		expiresIn := 2 * time.Hour
		tokenString, err := MakeJWT(userID, secret, expiresIn) // Reuse MakeJWT for convenience
		if err != nil {
			t.Fatalf("Failed to create test token: %v", err)
		}

		parsedUserID, err := ValidateJWT(tokenString, secret)
		if err != nil {
			t.Errorf("ValidateJWT failed: %v", err)
		}
		if parsedUserID != userID {
			t.Errorf("ValidateJWT returned incorrect UserID: got %v, want %v", parsedUserID, userID)
		}
	})

	// Test Case 2: Invalid Signature
	t.Run("Invalid Signature", func(t *testing.T) {
		userID := uuid.New()
		secret := "test-secret"
		wrongSecret := "wrong-secret"
		expiresIn := 2 * time.Hour
		tokenString, err := MakeJWT(userID, secret, expiresIn) // Use correct secret here
		if err != nil {
			t.Fatalf("Failed to create test token: %v", err)
		}

		_, err = ValidateJWT(tokenString, wrongSecret) // Use incorrect secret here
		if err == nil {
			t.Errorf("ValidateJWT should have failed with invalid signature, got nil error")
		}
	})

	// Test Case 3: Invalid Token Format
	t.Run("Invalid Token Format", func(t *testing.T) {
		invalidTokenString := "invalid-token-string"
		secret := "test-secret" // Secret is needed to construct the token.

		_, err := ValidateJWT(invalidTokenString, secret)
		if err == nil {
			t.Errorf("ValidateJWT should have failed with invalid token format, got nil error")
		}
	})
}

func TestGetBearerToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/protected", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	testCases := []struct {
		jwtToken      string
		expected      string
		expectedError bool
	}{{
		jwtToken: "Bearer aaaa.bbb.ccc", expected: "aaaa.bbb.ccc", expectedError: false},
		{jwtToken: "A aaaa.bbb.ccc", expected: "aaaa.bbb.ccc", expectedError: true},
		{jwtToken: "aaa-aaaa-aaa", expected: "aaaa.bbb.ccc", expectedError: true},
	}

	for _, tc := range testCases {
		req.Header.Set("Authorization", tc.jwtToken)
		token, err := GetBearerToken(req.Header)

		if tc.expectedError {
			if err == nil {
				t.Errorf("GetBearerToken(%v) expected error, got nil", req.Header)
			}
		} else {
			if err != nil {
				t.Errorf("GetBearerToken(%v) returned unexpected error: %v", req.Header, err)
			}

			if token != tc.expected {
				t.Errorf("GetBearerToken(%q) doesn't extract the right token", req.Header)
			}
		}
	}
}
