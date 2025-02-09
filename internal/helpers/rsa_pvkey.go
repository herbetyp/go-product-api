package helpers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"log"
	"os"

	"github.com/herbetyp/go-product-api/configs/logger"
)

var jwtSignKey *rsa.PrivateKey

func init() {
	var err error
	jwtSignKey, err = readPrivateKeyFromFile(os.Getenv("JWT_SIGN_KEY_FILE"))
	if err != nil {
		log.Panic("Error reading private key file", err)
	}
}

func readPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	file, err := os.Open(filename)
	if err != nil {
		logger.Error("Error opening file", err)
		return nil, err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		logger.Error("Error reading file", err)
		return nil, err
	}

	data, _ := pem.Decode(buffer)
	privateKey, err := x509.ParsePKCS8PrivateKey(data.Bytes)
	if err != nil {
		logger.Error("Error parsing private key", err)
		return nil, err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		logger.Error("Error load private key", err)
		return nil, err
	}

	jwtSignKey = rsaPrivateKey
	return rsaPrivateKey, nil
}

func GetPrivateKey() *rsa.PrivateKey {
	return jwtSignKey
}
