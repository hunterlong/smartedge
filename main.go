package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

type Key struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
}

type Message struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	Pubkey    string `json:"pubkey"`
}

const (
	keyFile = "private.pem"
	keySize = 2056
)

func main() {
	args := os.Args

	if len(args) != 2 {
		handleError(fmt.Errorf("there was no message inputted"))
	}

	input := args[1]

	key, err := OpenKey(keyFile, nil)
	if err != nil {
		if key, err = CreatePrivateKey(keyFile); err != nil {
			handleError(err)
		}
	}
	message, err := key.Encrypt(input)
	if err != nil {
		handleError(err)
	}

	fmt.Println(message.JSON())
}

// handleError will print the error and exit the application
func handleError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

// CreatePrivateKey will create a private.pem file that includes the private key
func CreatePrivateKey(filename string) (*Key, error) {
	priv, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, err
	}
	err = priv.Validate()
	if err != nil {
		return nil, err
	}
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)
	priv, pub, err := bytesToPrivateKey(string(privBytes), nil)
	if err != nil {
		return nil, err
	}
	if err := ioutil.WriteFile(filename, privBytes, 0655); err != nil {
		return nil, err
	}
	key := &Key{
		Private: priv,
		Public:  pub,
	}
	return key, nil
}

// JSON will return the JSON string of a message
func (m Message) JSON() string {
	out, _ := json.Marshal(m)
	return string(out)
}

// ToBase64 converts []byte to a string
func ToBase64(sig []byte) string {
	s1 := base64.StdEncoding.EncodeToString(sig)
	s2 := ""
	var LEN int = 76
	for len(s1) > 76 {
		s2 = s2 + s1[:LEN] + "\n"
		s1 = s1[LEN:]
	}
	s2 = s2 + s1
	return s2
}

// PublicKey returns the public key as a string
func (k *Key) PublicKey() string {
	pubASN1, err := x509.MarshalPKIXPublicKey(k.Public)
	if err != nil {
		handleError(err)
	}
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})
	return string(pubBytes)
}

// PrivateKey returns the private key as a string
func (k *Key) PrivateKey() string {
	priv := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(k.Private),
	})
	return string(priv)
}

// Encrypt will encrypt a string and output a Message
func (k *Key) Encrypt(msg string) (*Message, error) {
	if len(msg) > 250 {
		return nil, fmt.Errorf("message length must be under 250 charcters")
	}
	shaHash := sha256.New()
	shaHash.Write([]byte(msg))
	ciphertext, err := rsa.EncryptOAEP(shaHash, rand.Reader, k.Public, []byte(msg), nil)
	if err != nil {
		return nil, err
	}
	return &Message{msg, ToBase64(ciphertext), k.PublicKey()}, nil
}

// OpenKey opens a private key file and returns a Key object
func OpenKey(filename string, password []byte) (*Key, error) {
	keyData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	priv, pub, err := bytesToPrivateKey(string(keyData), password)
	if err != nil {
		return nil, err
	}
	key := &Key{
		Private: priv,
		Public:  pub,
	}
	return key, nil
}

func KeyFromData(data string, pass []byte) (*Key, error) {
	priv, pub, err := bytesToPrivateKey(data, pass)
	if err != nil {
		return nil, err
	}
	key := &Key{
		Private: priv,
		Public:  pub,
	}
	return key, err
}

// bytesToPrivateKey converts []byte from the private key file and returns
// a rsa.PrivateKey and rsa.PublicKey for the Key object
func bytesToPrivateKey(priv string, pass []byte) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(priv))
	if block == nil {
		return nil, nil, fmt.Errorf("incorrect RSA private key format")
	}
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, pass)
		if err != nil {
			return nil, nil, err
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, nil, err
	}
	pub := key.PublicKey
	return key, &pub, err
}
