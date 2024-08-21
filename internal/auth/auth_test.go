// internal/auth/auth_test.go

package auth

import (
	"fmt"
	"testing"
)

func TestGenerateAndValidateJwt(t *testing.T) {
	// Generate a JWT token
	tokenString, err := GenerateJwt("test_user")
	if err != nil {
		t.Fatalf("Error generating token: %v", err)
	}
	fmt.Println("Generated Token:", tokenString)

	// Validate the generated JWT token
	claims, err := ValidateJwt(tokenString)
	if err != nil {
		t.Fatalf("Error validating token: %v", err)
	}
	fmt.Println("Token Validated, Claims:", claims)

	// Check if the username is correct
	if claims.Username != "test_user" {
		t.Errorf("Expected username 'test_user', got '%s'", claims.Username)
	}
}
