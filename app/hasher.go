// IMPORTANT: THIS FILE CAN BE EDITED TO FIT YOUR NEEDS

package app

import "golang.org/x/crypto/bcrypt"

func HashString(toBeHashed string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(toBeHashed), 10)

	if err != nil {
		return "", err
	}

	hashedString := string(hashed)

	return hashedString, nil
}

func CompareHashedString(hashedString string, toBeCompared string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(toBeCompared))
}
