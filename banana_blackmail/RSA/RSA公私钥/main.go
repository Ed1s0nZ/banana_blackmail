// RSA公钥和私钥生成：
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func RSAKeyGen(bits int) error {
	privatekey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("私钥文件生成失败")
	}
	fmt.Println("私钥为：", privatekey)
	derStream := x509.MarshalPKCS1PrivateKey(privatekey)
	block := &pem.Block{
		Type:  "RSA Private key",
		Bytes: derStream,
	}
	privatefile, err := os.Create("myprivatekey.pem")
	defer privatefile.Close()
	err = pem.Encode(privatefile, block)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	publickey := &privatekey.PublicKey
	fmt.Println("公钥为：", publickey)
	derpkix, err := x509.MarshalPKIXPublicKey(publickey)
	block = &pem.Block{
		Type:  "RSA Public key",
		Bytes: derpkix,
	}
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	publickfile, err := os.Create("mypublic.pem")
	defer publickfile.Close()
	err = pem.Encode(publickfile, block)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func main() {
	var bits int
	flag.IntVar(&bits, "b", 1024, "密码默认长度1024")
	flag.Parse()
	err := RSAKeyGen(bits)
	if err != nil {
		fmt.Println("RSA密码文件生成失败")
	}
	fmt.Println("RSA密码生成成功")
}
