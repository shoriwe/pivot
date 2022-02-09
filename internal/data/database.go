package data

import (
	"encoding/hex"
	"github.com/shoriwe/pivot/internal/data/objects"
	"golang.org/x/crypto/sha3"
)

type Database interface {
	Login(email, password string) (user *objects.User, succeed bool, err error)
	Register(user *objects.User) (succeed bool, err error)
}

func CalcHash(s string) []byte {
	hash := sha3.New512()
	hash.Write([]byte(s))
	return hash.Sum(nil)
}

func CheckPasswords(password, hashedPassword string) bool {
	return hex.EncodeToString(CalcHash(password)) == hashedPassword
}

type Connection struct {
	DB    Database
	Cache *Cache
}

func NewConnection(database Database) *Connection {
	return &Connection{
		DB:    database,
		Cache: NewCache(),
	}
}
