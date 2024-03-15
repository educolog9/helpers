/*
 * File: generate_token.go
 * Author: David Bujosa
 * Date: 3/15/2024
 * Description: This file contains the implementation of the GenerateRandomSalt, EncodeStructToBase64,
 * and DecodeBase64ToStruct functions.
 */
package functions

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"time"
)

type EncodeToken struct {
	Email string    `json:"email"`
	Salt  string    `json:"salt"`
	Exp   time.Time `json:"exp"`
}

// GenerateRandomSalt generates a random salt string.
// It uses the crypto/rand package to generate a random byte slice and encodes it using base64.URLEncoding.
// The resulting string is suitable for use as a salt in password hashing.
func GenerateRandomSalt() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// EncodeStructToBase64 encodes email to a base64 string.
func EncodeStructToBase64(email string, duration time.Duration) (string, error) {
	s := EncodeToken{
		Email: email,
		Salt:  GenerateRandomSalt(),
		Exp:   time.Now().Add(duration),
	}

	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	base64Data := base64.URLEncoding.EncodeToString(jsonData)
	return base64Data, nil
}

// DecodeBase64ToStruct decodes a base64 string to a struct.
func DecodeBase64ToStruct(base64Data string) (EncodeToken, error) {
	var s EncodeToken

	jsonData, err := base64.URLEncoding.DecodeString(base64Data)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(jsonData, &s)
	if err != nil {
		return s, err
	}

	if time.Now().After(s.Exp) {
		return s, errors.New("token has expired")
	}

	return s, nil
}
