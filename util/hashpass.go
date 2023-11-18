package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(pwd string) string {
	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice

	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Clm(m map[string]interface{}) (*ClaimBody, bool) {
	var c ClaimBody
	c.Phone = m["phone"].(string)
	for _, v := range m["roles"].([]interface{}) {
		c.Roles = append(c.Roles, v.(string))
	}
	c.UserId = m["user_id"].(string)
	return &c, true
}

type ClaimBody struct {
	UserId string `json:"user_id"`
	Roles  []string `json:"roles"`
	Phone  string `json:"phone"`
}
 