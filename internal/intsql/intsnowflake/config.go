package intsnowflake

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/snowflakedb/gosnowflake"
)

type Config struct {
	Account        string `json:"account"`
	User           string `json:"user"`
	PrivateKeyFile string `json:"private_key_file"`
}

func ToSnowflakeConfig(fileName string) (*gosnowflake.Config, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("reading snowflake config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("parsing snowflake config file: %w", err)
	}

	rsaBytes, err := os.ReadFile(config.PrivateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("reading private key file: %w", err)
	}

	privPem, _ := pem.Decode(rsaBytes)
	if privPem.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("expected private key pem type PRIVATE KEY.  Found: %v", privPem.Type)
	}
	privPemBytes := privPem.Bytes

	var parsedKey interface{}
	parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes) // returns interface{}
	if err != nil {
		return nil, fmt.Errorf("parsing private key: %v", err)
	}

	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("expected private key to be of type *rsa.PrivateKey.  Found: %T", privateKey)
	}

	sfConfig := &gosnowflake.Config{
		Account:       config.Account,
		User:          config.User,
		Authenticator: gosnowflake.AuthTypeJwt,
		PrivateKey:    privateKey,
	}

	return sfConfig, nil

}
