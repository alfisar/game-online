package hashgenerator

import "golang.org/x/crypto/bcrypt"

// Generate generates hash of data
func Generate(data string) (result string, err error) {
	pwd := []byte(data)
	pwd, err = bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}
	result = string(pwd[:])
	return
}

// Verify verifies password with hashedPassword
func Verify(hashedPassword string, password string) (result bool, err error) {
	hshPwd := []byte(hashedPassword)
	pwd := []byte(password)
	err = bcrypt.CompareHashAndPassword(hshPwd, pwd)
	if err != nil {
		return
	}
	result = true
	return
}
