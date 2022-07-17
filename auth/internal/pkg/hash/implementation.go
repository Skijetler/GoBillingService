package hash

import (
	"encoding/hex"

	"github.com/Skijetler/GoBillingService/auth/internal/conf"
	"golang.org/x/crypto/argon2"
)

type PasswordHasher interface {
	GenerateHash(password string) string
	CompareHash(hash, password string) bool
}

type Hasher struct {
	salt       []byte
	iterations uint32
	memory     uint32
	threads    uint8
	keyLen     uint32
}

func NewPasswordHasher(c *conf.Hasher) PasswordHasher {
	if c.Iterations == 0 {
		c.Iterations = 1
	}
	if c.Memory == 0 {
		c.Memory = 64 * 1024
	}
	if c.Threads == 0 {
		c.Threads = 1
	}
	if c.KeyLen == 0 {
		c.KeyLen = 32
	}

	return &Hasher{
		salt:       []byte(c.Salt),
		iterations: c.Iterations,
		memory:     c.Memory,
		threads:    uint8(c.Threads),
		keyLen:     c.KeyLen,
	}
}

func (h *Hasher) GenerateHash(password string) string {
	hash := argon2.IDKey([]byte(password), h.salt, h.iterations, h.memory, h.threads, h.keyLen)
	return hex.EncodeToString(hash)
}

func (h *Hasher) CompareHash(hash, password string) bool {
	passHash := h.GenerateHash(password)
	return hash == passHash
}
