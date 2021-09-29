package Directory_traversal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	AES_Decryption "github.com/ed11s00n/blackmail_dec/AES_DEcryption"
)

var wg sync.WaitGroup
var k int

// 遍历pathName下所有文件
func GetAllFile(pathName string) {
	conn := pathName
	local_dir := conn
	err := filepath.Walk(local_dir, file_path)
	fmt.Println("Temp总共文件数量:", k)
	if err != nil {
		// fmt.Println("路径获取错误")
		return
	}
}

func file_path(FileName string, fi os.FileInfo, err error) error {
	if fi.IsDir() {
		return nil
	}
	k++
	fmt.Println("filename:", FileName) // 输出文件名字 ->  byte读取文件 aes加密文件内容 将加密后的文件新建到 fileName.banana
	// 读取文件AES 加密
	if strings.HasSuffix(FileName, ".banana") { // 判断后缀是否为.banana 文件 只解密.banana后缀的文件
		wg.Add(1)
		go func() {
			defer wg.Add(-1)
			AES_Decryption.File_Processing(FileName) // 这里实现解密

		}()
		wg.Wait()
		return nil
	}
	return nil
}
func Directory_traversal(pathName string) { //主
	GetAllFile(pathName)
}
