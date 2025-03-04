package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const (
	// Character sets for password generation
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// generateRandomInt generates a cryptographically secure random number between 0 and max-1
func generateRandomInt(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

// generatePassword creates a secure password with the specified length
func generatePassword(length int) (string, error) {
	// Combine all character sets
	allChars := lowercase + uppercase + numbers + symbols
	allCharsLen := len(allChars)

	// Ensure minimum length of 8 characters
	if length < 8 {
		length = 8
	}

	// Create a slice to store the password
	password := make([]byte, length)

	// Ensure at least one character from each set
	requiredChars := []string{lowercase, uppercase, numbers, symbols}
	for _, chars := range requiredChars {
		pos, err := generateRandomInt(length)
		if err != nil {
			return "", err
		}
		charPos, err := generateRandomInt(len(chars))
		if err != nil {
			return "", err
		}
		password[pos] = chars[charPos]
	}

	// Fill the rest of the password with random characters
	for i := 0; i < length; i++ {
		if password[i] == 0 { // If position is not filled
			pos, err := generateRandomInt(allCharsLen)
			if err != nil {
				return "", err
			}
			password[i] = allChars[pos]
		}
	}

	return string(password), nil
}

func main() {
	// Parse command line arguments
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Printf("Usage: %s <length>\n", os.Args[0])
		fmt.Println("Example: pwgen 16")
		os.Exit(1)
	}

	// Convert length argument to integer
	length, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error: Length must be a number\n")
		os.Exit(1)
	}

	// Generate password
	password, err := generatePassword(length)
	if err != nil {
		fmt.Printf("Error generating password: %v\n", err)
		os.Exit(1)
	}

	// Print the generated password
	fmt.Println(password)
} 