// crypto.go kee > 2019/11/11

package kiris

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

// AES 加密
func AESEncrypt(origin, key, mod string) string {
	originBytes, keyBytes := []byte(origin), []byte(key)
	// 分组密钥
	block, e := aes.NewCipher(keyBytes)
	if e != nil {
		log.Fatal(e)
	}
	// 获取密钥长度
	blockSize := block.BlockSize()
	// pkcs7
	originBytes = PKCS7Padding(originBytes, blockSize)
	// 模式
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])
	// encrypt
	cryted := make([]byte, len(originBytes))
	blockMode.CryptBlocks(cryted, originBytes)

	return Base64Encode(cryted)
}

// AES 解码
func AESDecrypt(cryted, key, mod string) string {
	crytedBytes, keyBytes := Base64Decode(cryted), []byte(key)
	// 分组秘钥
	block, e := aes.NewCipher(keyBytes)
	if e != nil {
		log.Fatal(e)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	origin := make([]byte, len(crytedBytes))
	// 解码
	blockMode.CryptBlocks(origin, crytedBytes)
	// pkcs7 去码
	origin = PKCS7UnPadding(origin)
	return string(origin)
}

// PKCS7 补码
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS7 去码
func PKCS7UnPadding(originBytes []byte) []byte {
	length := len(originBytes)
	unpadding := int(originBytes[length-1])
	return originBytes[:(length - unpadding)]
}

// Base64 转码
func Base64Encode(origin []byte) string {
	return base64.StdEncoding.EncodeToString(origin)
}

// Base64 解码
func Base64Decode(encode string) []byte {
	decode, _ := base64.StdEncoding.DecodeString(encode)
	return decode
}
