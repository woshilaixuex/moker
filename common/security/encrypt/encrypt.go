package encrypt

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/scrypt"
)

// 盐加密: salt + hashCode
type (
	CodeData struct {
		Salt     string
		HashCode string
	}
	EncryptPass interface {
		Encrypt(password string) error
		DeCode(password string) (bool, error)
	}
)

func NewCodeData() *CodeData {
	return &CodeData{}
}

func (data *CodeData) Encrypt(password string) (*CodeData, error) {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		logx.Errorf(err.Error())
	}

	//  使用Scrypt算法生成哈希值
	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
	if err != nil {
		logx.Errorf(err.Error())
	}

	data.HashCode = base64.StdEncoding.EncodeToString(hash)
	data.Salt = base64.StdEncoding.EncodeToString(salt)
	return data, nil
}
func (data *CodeData) DeCode(password, hashcode, salt string) (bool, error) {

	retrievedHash, err := base64.StdEncoding.DecodeString(hashcode)
	if err != nil {

	}
	retrievedSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {

	}
	newHash, err := scrypt.Key([]byte(password), retrievedSalt, 16384, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}

	return subtle.ConstantTimeCompare(newHash, retrievedHash) == 1, nil
}
