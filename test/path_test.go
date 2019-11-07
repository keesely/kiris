package kiris

import (
	"fmt"
	"testing"
)

func TestIsDir(t *testing.T) {
	dir1 := "~/"
	dir2 := "./data"
	dir3 := "./path.go"

	fmt.Println(StrPad("", "=", 100, 0))
	fmt.Println("Print Is Dir")
	fmt.Printf("> PATH: %s => %v\n", dir1, IsDir(dir1))
	fmt.Printf("> PATH: %s => %v\n", dir2, IsDir(dir2))
	fmt.Printf("> PATH: %s => %v\n", dir3, IsDir(dir3))
}

func TestIsFile(t *testing.T) {
	file1 := "~/.ssh/id_rsa"
	file2 := "./data"
	file3 := "./path.go"

	fmt.Println(StrPad("", "=", 100, 0))
	fmt.Println("Print Is File")
	fmt.Printf("> PATH: %s => %v\n", file1, IsFile(file1))
	fmt.Printf("> PATH: %s => %v\n", file2, IsFile(file2))
	fmt.Printf("> PATH: %s => %v\n", file3, IsFile(file3))

	fmt.Println(StrPad("", "-", 100, 0))
	fmt.Println("Print File Exits")
	fmt.Printf("> PATH: %s => %v\n", file1, FileExists(file1))
	fmt.Printf("> PATH: %s => %v\n", file2, FileExists(file2))
	fmt.Printf("> PATH: %s => %v\n", file3, FileExists(file3))
}
