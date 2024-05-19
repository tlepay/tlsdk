package tlsdk

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	APIURL           string `json:"api_url" valid:"api_url,r|s,l:0;"`                         // API地址
	MchID            string `json:"mch_id" valid:"mch_id,r|s,l:0;"`                           // 商户号
	APPID            string `json:"app_id" valid:"app_id,r|s,l:0;"`                           // APPID
	RSATLPublicKey   string `json:"rsa_tl_public_key" valid:"rsa_tl_public_key,r|s,l:0;"`     // RSA通联公钥
	RSAMCHPrivateKey string `json:"rsa_mch_private_key" valid:"rsa_mch_private_key,r|s,l:0;"` // RSA商户私钥
	RSAMCHPublicKey  string `json:"rsa_mch_public_key" valid:"rsa_mch_public_key,r|s,l:0;"`   // RSA商户公钥
}

func NewConfigWithFile(fileDir string) (config *Config, err error) {
	file, err := os.Open(fileDir)
	if err != nil {
		return
	}
	defer file.Close()
	config = &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return
	}
	if len(config.APIURL) < 1 {
		err = fmt.Errorf("api_url is required")
		return
	}
	if len(config.MchID) < 1 {
		err = fmt.Errorf("mch_id is required")
		return
	}
	if len(config.APPID) < 1 {
		err = fmt.Errorf("app_id is required")
		return
	}
	if len(config.RSATLPublicKey) < 1 {
		err = fmt.Errorf("rsa_tl_public_key is required")
		return
	}
	if len(config.RSAMCHPrivateKey) < 1 {
		err = fmt.Errorf("rsa_mch_private_key is required")
		return
	}
	config.RSAMCHPrivateKey = formatRSAPrivKey(config.RSAMCHPrivateKey)
	config.RSATLPublicKey = formatRSAPublicKey(config.RSATLPublicKey)
	return
}

func formatRSAPublicKey(public string) string {
	const pubKeyHead = "-----BEGIN RSA PUBLIC KEY-----\n"
	const pubKeyTail = "\n-----END RSA PUBLIC KEY-----"

	if !strings.Contains(public, pubKeyHead) {
		public = pubKeyHead + public
	}
	if !strings.Contains(public, pubKeyTail) {
		public = public + pubKeyTail
	}
	return public
}

func formatRSAPrivKey(priv string) string {
	// 使用常量来定义私钥的头部和尾部，方便维护
	const privKeyHead = "-----BEGIN RSA PRIVATE KEY-----\n"
	const privKeyTail = "\n-----END RSA PRIVATE KEY-----"

	// 检查是否已包含头部和尾部
	if !strings.Contains(priv, privKeyHead) {
		priv = privKeyHead + priv
	}
	if !strings.Contains(priv, privKeyTail) {
		priv = priv + privKeyTail
	}
	return priv
}
