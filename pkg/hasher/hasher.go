package hasher

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/scrypt"
)

// Hasher allow generate a new password
type Hasher interface {
	Hash(pass string) ([]byte, []byte, error)
	Decrypt(pass string, salt []byte) ([]byte, error)
}

type hasher struct {
	scryptN   int
	scryptP   int
	scryptR   int
	scryptLen int
}

// New return an Hasher with the default configuration
func New() Hasher {
	c := DefaultConf()
	return NewWithConfig(c)
}

// NewWithConfig return an Hasher with the configuration parameters
func NewWithConfig(c Config) Hasher {
	return hasher{
		scryptN:   c.ScryptN,
		scryptP:   c.ScryptP,
		scryptR:   c.ScryptR,
		scryptLen: c.ScryptLen,
	}
}

// Hash the pass and return
// - password generated
// - salt used to generate the password
// - an error if occured
func (h hasher) Hash(pass string) ([]byte, []byte, error) {
	var err error
	var password []byte
	salt := make([]byte, 32)

	// generate random salt
	if _, err = io.ReadFull(rand.Reader, salt); err != nil {
		return nil, nil, err
	}

	password, err = scrypt.Key([]byte(pass), salt, h.scryptN, h.scryptR, h.scryptP, h.scryptLen)
	if err != nil {
		return nil, nil, err
	}
	return password, salt, nil
}

// Decrypt the pass with the salt parameter provided
func (h hasher) Decrypt(pass string, salt []byte) ([]byte, error) {
	password, err := scrypt.Key([]byte(pass), salt, h.scryptN, h.scryptR, h.scryptP, h.scryptLen)
	return password, err
}
