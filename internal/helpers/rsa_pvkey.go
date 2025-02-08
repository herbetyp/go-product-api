package helpers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	// "log"
	"os"

	"github.com/herbetyp/go-product-api/configs/logger"
)

var jwtSignKey *rsa.PrivateKey

func init() {
	var err error
	jwtSignKey, err = readPrivateKeyFromFile(os.Getenv("SIGN_KEY_FILENAME"))
	if err != nil {
		logger.Error("Error reading private key file", err)
		panic(err)
	}
}

func readPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	data, _ := pem.Decode(buffer)
	privateKey, err := x509.ParsePKCS8PrivateKey(data.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("expected RSA private key")
	}

	jwtSignKey = rsaPrivateKey
	return rsaPrivateKey, nil
}

func GetPrivateKey() *rsa.PrivateKey {
	return jwtSignKey
}
