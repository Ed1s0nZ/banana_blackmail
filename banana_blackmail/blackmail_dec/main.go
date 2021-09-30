package main

import (
	"github.com/ed11s00n/blackmail_dec/Directory_traversal"
)

func main() {
	var pathName string
	flag.StringVar(&pathName, "pathName", "C:/Users/13107/Desktop/识别沙箱检测", "输入要加密的路径")
	flag.Parse()
// 	pathName := "C:/Users/13107/Desktop/识别沙箱检测"
	Directory_traversal.Directory_traversal(pathName) // 解密文件
}
