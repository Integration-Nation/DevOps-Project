package security

import (
    "github.com/golang-jwt/jwt/v4"
    "os"
    "time"
)

// GenerateJWT generates a new JWT token for a given user ID and username.
func GenerateJWT(userID int, username string) (string, error) {
    // Get the JWT secret from environment variables
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        return "", errors.New("JWT secret not set in environment variables")
    }

    claims := jwt.MapClaims{
        "user_id":   userID,           // Add user ID as a claim
        "username":  username,         // Add username as a claim
        "exp":       time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
    }

    // Create a new token with the claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with the secret key from the environment
    return token.SignedString([]byte(jwtSecret))
}
