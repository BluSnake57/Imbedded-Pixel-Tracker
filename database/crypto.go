package Database

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// GenerateRandomBytes returns securely generated random bytes
func generaterandombytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomStringURLSafe returns a URL-safe, base64 encoded
func generaterandomstringURLsafe(n int) (string, error) {
	b, err := generaterandombytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

// GenerateRandomStringURLSafe panics if the called functions error
func GenerateRandomStringURLSafe(n int) string {
	token, err := generaterandomstringURLsafe(n)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		panic(err)
	}
	return token
}

// Hashes the users clientID for storage
func HashClientID(clientID string) string {
	hashedID, err := bcrypt.GenerateFromPassword([]byte(clientID), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing ID:", err)
		return string(hashedID)
	}

	log.Printf("Hashed ID: %s\n", string(hashedID))

	return string(hashedID)
}

// Verifies the clientID against the stored hash
func VerifyClientID(hashedID, clientID string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedID), []byte(clientID))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Println("Invalid ID")
		} else {
			log.Println("Error verifying ID:", err)
		}
		return false
	}
	return true
}
