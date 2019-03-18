package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/imulab/soteria/pkg/oauth"
	"github.com/imulab/soteria/pkg/utility"
	"hash"
)

type HmacShaStrategy interface {
	// Generate a new key with signature.
	Generate(entropy uint) (key string, sig string, err error)
	// Verify a key against its signature.
	Verify(key, sig string) error
}

// Create a new HMAC-SHA strategy using SHA-256 algorithm.
func NewHmacSha256Strategy(signingKey []byte) (HmacShaStrategy, error) {
	return newHmacSha(signingKey, 32, sha256.New)
}

// Create a new HMAC-SHA strategy using SHA-384 algorithm.
func NewHmacSha384Strategy(signingKey []byte) (HmacShaStrategy, error) {
	return newHmacSha(signingKey, 48, sha512.New384)
}

// Create a new HMAC-SHA strategy using SHA-512 algorithm.
func NewHmacSha512Strategy(signingKey []byte) (HmacShaStrategy, error) {
	return newHmacSha(signingKey, 64, sha512.New)
}

var (
	b64 = base64.URLEncoding.WithPadding(base64.NoPadding)
)

func newHmacSha(signingKey []byte, sigKeyLen int, hashFunc	func() hash.Hash) (*hmacSha, error) {
	if len(signingKey) != sigKeyLen {
		return nil, fmt.Errorf("signing key length must be exactly %d bits", sigKeyLen * 8)
	}

	// copy for safety
	copiedSigningKey := make([]byte, len(signingKey))
	copy(copiedSigningKey, signingKey)

	return &hmacSha{
		signingKey:copiedSigningKey,
		hashFunc:hashFunc,
	}, nil
}

type hmacSha struct {
	// Secret key for signing the generated key
	signingKey 		[]byte
	// Function to generate a new hmac-sha hash
	hashFunc		func() hash.Hash
}

func (h *hmacSha) Generate(entropy uint) (string, string, error) {
	rawKey, err := utility.RandomBytes(entropy)
	if err != nil {
		return "", "", err
	}

	rawSig := h.sign(rawKey, h.signingKey)

	return b64.EncodeToString(rawKey), b64.EncodeToString(rawSig), nil
}

func (h *hmacSha) Verify(key, sig string) error {
	var (
		err	error
	)

	rawKey, err := b64.DecodeString(key)
	if err != nil {
		return err
	}

	rawSig, err := b64.DecodeString(sig)
	if err != nil {
		return err
	}

	expectSig := h.sign(rawKey, h.signingKey)
	if !hmac.Equal(expectSig, rawSig) {
		return oauth.ErrSignatureMismatch
	}

	return nil
}

func (h *hmacSha) sign(data []byte, key []byte) []byte {
	hs := hmac.New(h.hashFunc, key[:])
	hs.Write(data)
	return hs.Sum(nil)
}