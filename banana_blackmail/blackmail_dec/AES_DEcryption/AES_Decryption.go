package AES_Decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ed11s00n/blackmail_dec/RSA_DEC_KEY"
)

var iv = "0000000000000000"

func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// 删除原加密文件
func Delete_File(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		return
	} else {
		return
	}
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//  解密文件
func AES_Decryption(key string, data string) []byte { // key为AES加密文件时随机生成的key , data为加密以后的数据，返回byte类型原数据
	var info_list = [...]string{data}
	original_data, _ := AesDecrypt(info_list[0], []byte(key))
	// fmt.Println(string(original_data))
	return original_data
}

// 主 对传入文件进行处理
func File_Processing(fileName string) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	RSA_key := lineSlice[0]
	key := RSA_DEC_KEY.RSA_DEC_KEY(RSA_key)
	data := lineSlice[4]
	Delete_File(fileName) // 删除加密文件
	original_data := AES_Decryption(key, data)
	err2 := ioutil.WriteFile(fileName[:len(fileName)-7], original_data, 0666)
	if err2 != nil {
		return
	}
}
