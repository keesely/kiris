// crypto_test.go kee > 2019/11/11

package kiris

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	orig := "这是一段密码明文"
	key := "123456781234567812345678"
	fmt.Println("Key: ", len(key), " => ", key)
	fmt.Println("原文：", orig)

	encryptCode := AESEncrypt(orig, key, "cbc")
	fmt.Println("密文：", Base64Encode(encryptCode))

	//key = Md5Hash("hello world")
	//key = Md5Hash("Keesely.net")
	decryptCode := AESDecrypt(encryptCode, key, "cbc")
	fmt.Println("解密结果：", decryptCode)
}
