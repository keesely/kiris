// hash_test.go kee > 2019/11/11

package hash

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	str := "hello world"
	md5 := Md5(str)
	sha1 := Sha1(str)
	sha256 := Sha256(str)
	sha512 := Sha512(str)
	fmt.Printf("md5 => LEN(%d): %s\n", len(md5), md5)
	fmt.Printf("sha1 => LEN(%d): %s\n", len(sha1), sha1)
	fmt.Printf("sha256 => LEN(%d): %s\n", len(sha256), sha256)
	fmt.Printf("sha512 => LEN(%d): %s\n", len(sha512), sha512)
}
