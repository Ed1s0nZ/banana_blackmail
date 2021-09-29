package RSA_ENC_KEY

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var publickey = FileLoad("mypublic.pem")

func FileLoad(filepath string) []byte {
	privatefile, err := os.Open(filepath)
	defer privatefile.Close()
	if err != nil {
		return nil
	}
	privateKey := make([]byte, 2048)
	num, err := privatefile.Read(privateKey)
	return privateKey[:num]
}

func RSA_ENC_KEY(orgidata []byte) ([]byte, error) { //加密函数
	block, _ := pem.Decode(publickey)
	if block == nil {
		return nil, errors.New("public key is bad")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, orgidata) //加密
}

func RSAEncrypt(key string) string { // 主RSAEncrypt
	var data []byte
	var err error
	data, err = RSA_ENC_KEY([]byte(key))
	if err != nil {
		fmt.Println("错误", err)
	}
	// fmt.Println("加密：", base64.StdEncoding.EncodeToString(data))
	return base64.StdEncoding.EncodeToString(data)

}
