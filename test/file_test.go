/*************************************************************************
   > File Name: file_test.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2019.11.07
************************************************************************/
package kiris

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	err := FilePut("./test.txt", "这是一段测试文本\nThis is a test content")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("文本写入 / Write content")
	}

	err = FileAppend("./test.txt", "\n追加一段文本\nAppend content")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("文本追加 / Append content")
	}

	exists := FileExists("./test.txt")

	if true == exists {
		fmt.Println("文件已存在 / File exists")
		content, _ := FileGetContents("./test.txt")
		fmt.Println("文本内容(The content): \n```\n", content, "\n```")
	} else {
		fmt.Println("文件不存在 / File is not exists")
	}

	path := "~/.ssh/*.pub"
	list, _ := FileSearch(path)

	fmt.Println(StrPad("", "=", 100, 0))
	fmt.Println(StrPad(" Scanning the PATH ("+path+") ", "+", 100, KIRIS_STR_PAD_BOTH))
	fmt.Println(StrPad("", "=", 100, 0))
	fmt.Printf("Found Total: %v \n", len(list))
	fmt.Println(StrPad("", "-", 100, 0))
	for k, v := range list {
		fmt.Printf(" [%d] %-85s Found.\n", k, v)
	}
	fmt.Println(StrPad("", "=", 100, 0))

}
