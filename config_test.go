package tlsdk

import (
	"fmt"
	"os"
	"testing"

	"github.com/codingeasygo/util/converter"
)

func TestWriteConfig(t *testing.T) {
	pub, pri, _ := CreateRSAKeyPair()

	fmt.Printf("pub %v\n", pub)
	fmt.Printf("pri %v\n", pri)

	config := &Config{
		APIURL: "https://api.example.com",
		MchID:  "M170123456789",
		APPID:  "6629fecd9b8e41000700000d",
		RSATLPublicKey: `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv6JJl9tqyKN5ldc7YE0A
ztSR8U1J4jcymK0Um6E3MxiXI3iIc9smQh8wsMMyymBE41cQkZUxbsoLhm30MABE
cmKCkHoF+jfUib0qJU422Led3ymkaBQzfo9BrBd4D8Aq72LPrw2IaMqZBwF/BRT/
1XAjGeauhKCEK4koXWvG7aRgVHEVmMKVKJIZKBCr7+9Fl3JTwy5jn5k15rSOrFeS
5F6MeFg2xIqeKDBzMXtho4+89kGyiGi0MIA7mO7eCrUCt/7P0WIHKmvdWh6foQhP
lzQNdpZQuEN5j8+tnPorm4FqT5UdMOLRjGj0cIozsjUn9xZhFtMyRhqTSzeP76uE
pwIDAQAB
-----END RSA PUBLIC KEY-----
`,
		RSAMCHPrivateKey: pri,
		RSAMCHPublicKey:  pub,
	}

	file, err := os.Open("config.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	err = os.WriteFile("config.json", []byte(converter.JSON(config)), 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetRSAMchPub(t *testing.T) {
	config, err := NewConfigWithFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.RSAMCHPublicKey)
}

func TestConfig(t *testing.T) {
	config, err := NewConfigWithFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("config: %+v\n", config)
}
