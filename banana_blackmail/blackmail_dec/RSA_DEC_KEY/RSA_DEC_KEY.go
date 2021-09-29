package RSA_DEC_KEY

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

var privatekey = FileLoad("myprivatekey.pem")

func FileLoad(filepath string) []byte {
	privatefile, err := os.Open(filepath)
	defer privatefile.Close()
	if err != nil {
		return nil
	}
	privateKey := make([]byte, 2048)
	num, err := privatefile.Read(privateKey)
	if err != nil {
		return nil
	}
	return privateKey[:num]
}
func RSADecrypt(cipertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privatekey)
	if block == nil {
		return nil, errors.New("private key is bad")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipertext)
}

// key := "cOMzS+CqcN7+8jboX2YlblEWX8c0AR21v2WwCg8xadKyGKZEzd8Imzt1gEzQRwsrYXi+Zf71ok+XBvqaoYhbOBlI17tsBBElJjW7hHat65RH08T3W0v3Qp0xsPyUnVZhjFSBoSYaNFAV8BelR5AkszfaGCjId9GmZeCHBuXamL8="
//主 传入RSA加密的AES key
func RSA_DEC_KEY(RSA_key string) string { //主 传入RSA加密的AES key
	// var data []byte
	var err error
	// fmt.Println(RSA_key)

	data, err := base64.StdEncoding.DecodeString(RSA_key)
	if err != nil {
		fmt.Println("错误", err)
	}
	origData, err := RSADecrypt(data) //解密
	if err != nil {
		fmt.Println("错误", err)
	}
	// fmt.Println("解密:", string(origData))
	return string(origData)
	//pk := FileLoad("myprivatekey.pem")
	//fmt.Println(string(pk))
}
