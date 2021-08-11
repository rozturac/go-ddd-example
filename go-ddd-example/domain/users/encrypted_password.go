package users

import (
	"crypto/sha256"
	"go-ddd-example/domain/common"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type EncryptedPassword struct {
	PasswordHash [32]byte `json:"password_hash"`
	PasswordSalt [32]byte `json:"password_salt"`
}

func NewEncryptedPassword(password string) *EncryptedPassword {

	if len(password) == 0 {
		panic(common.IsNullOrEmptyError("password"))
	}

	var (
		passwordAsByte []byte
		passwordSalt   [32]byte
		passwordHash   [32]byte
	)

	passwordAsByte = []byte(password)
	passwordSalt = getPasswordSalt()

	for index, item := range passwordAsByte {
		passwordAsByte[index] = item ^ passwordSalt[index%32]
	}

	passwordHash = sha256.Sum256(passwordAsByte)

	return &EncryptedPassword{PasswordSalt: passwordSalt, PasswordHash: passwordHash}
}

func (ep *EncryptedPassword) VerifyPassword(password string) bool {
	if len(password) == 0 {
		panic(common.IsNullOrEmptyError("password"))
	}

	var (
		passwordAsByte []byte
		passwordHash   [32]byte
	)

	passwordAsByte = []byte(password)

	for index, item := range passwordAsByte {
		passwordAsByte[index] = item ^ ep.PasswordSalt[index%32]
	}

	passwordHash = sha256.Sum256(passwordAsByte)

	for i := 0; i < len(ep.PasswordHash); i++ {
		if ep.PasswordHash[i] != passwordHash[i] {
			return false
		}
	}

	return true
}

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func getPasswordSalt() [32]byte {
	var result [32]byte
	copy(result[:], []byte(stringWithCharset(32, charset)))
	return result
}
