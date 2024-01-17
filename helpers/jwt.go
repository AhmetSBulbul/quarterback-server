package helpers

import (
	"crypto/rsa"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/exp/maps"

	"github.com/golang-jwt/jwt"
)

var keys = make(map[string][]byte, 2)
var lock = sync.RWMutex{}

func readKeyFromFile(path string) []byte {
	if keys[path] != nil {
		return keys[path]
	} else {
		f, err := os.ReadFile(path)
		if err != nil {
			fmt.Errorf(err.Error())
			return []byte{}
		}
		lock.Lock()
		keys[path] = f
		lock.Unlock()
		return f
	}
}

func readPrivateKey(path string) (*rsa.PrivateKey, error) {
	pemFile := readKeyFromFile(path)
	return jwt.ParseRSAPrivateKeyFromPEM(pemFile)
}

func readPublicKey(path string) (*rsa.PublicKey, error) {
	pemFile := readKeyFromFile(path)
	return jwt.ParseRSAPublicKeyFromPEM(pemFile)
}

func CreateJWT(claims jwt.MapClaims) (string, error) {
	privateKeyPath, err := GetEnvOrFail("JWT_PRIVATE_KEY_PATH")
	if err != nil {
		return "", err
	}

	privateKey, err := readPrivateKey(privateKeyPath)
	if err != nil {
		return "", err
	}

	ttl := time.Minute * 120
	now := time.Now().UTC()
	allClaims := jwt.MapClaims{
		"iss": "rnd-auth",
		"exp": now.Add(ttl).Unix(),
	}
	maps.Copy(allClaims, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, allClaims)
	return token.SignedString(privateKey)
}

func ValidateJWT(token string) (*jwt.Token, error) {
	publicKeyPath, err := GetEnvOrFail("JWT_PUBLIC_KEY_PATH")
	if err != nil {
		return nil, err
	}

	publicKey, err := readPublicKey(publicKeyPath)
	if err != nil {
		return nil, err
	}

	keyFunc := jwt.Keyfunc(func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	return jwt.Parse(token, keyFunc)
}
