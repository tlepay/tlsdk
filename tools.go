package tlsdk

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"sync"
	"time"
)

var MarchineID = 1
var seqOrderID uint16
var lckOrderID = sync.RWMutex{}

func NewOrderID() (orderID string) {
	lckOrderID.Lock()
	defer lckOrderID.Unlock()
	seqOrderID++
	timeStr := time.Now().Format("20060102150405")
	return fmt.Sprintf("%v%02d%05d", timeStr, MarchineID, seqOrderID)
}

func RSAVerify(pubKey, data []byte, base64Sign string) (err error) {
	hash := crypto.SHA256
	h := hash.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(pubKey)
	if block == nil {
		fmt.Printf("public key %v\n", string(pubKey))
		return fmt.Errorf("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub := pubInterface.(*rsa.PublicKey)

	sign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return
	}
	return rsa.VerifyPKCS1v15(pub, hash, hashed, sign)
}

func RSASign(privKey []byte, data []byte) (sign string, err error) {
	hash := crypto.SHA256
	h := hash.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	block, _ := pem.Decode(privKey)
	if block == nil {
		err = fmt.Errorf("private key error")
		return
	}
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), hash, hashed)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(signature)
	return
}

var (
	defaultBits = 2048
)

func CreateRSAKeyPair() (pub, privPKCS8 string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, defaultBits)
	if err != nil {
		return
	}
	publicKey := &privateKey.PublicKey
	// // 序列化私钥为PKCS1格式字符串
	// privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	// privateKeyPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})
	// privPKCS1 = string(privateKeyPem)
	// 序列化公钥为字符串
	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyBytes})
	pub = string(publicKeyPem)

	// 序列化私钥为PKCS8格式字符串
	privateKeyPKCS8, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	privateKeyPemPKCS8 := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privateKeyPKCS8})
	privPKCS8 = string(privateKeyPemPKCS8)
	return
}
