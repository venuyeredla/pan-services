package vrdbms

import (
	"fmt"
	"os"
	"testing"
)

func TestRdbms(t *testing.T) {
	fmt.Println("Testing RDBMS")
}

func TestWrite(pageNum int) {

	page := new(Page)
	page.Num = pageNum
	pageSize := os.Getpagesize()
	bytes := make([]byte, 0, pageSize)
	for i := 0; i < pageSize; i++ {
		b := byte(i / 255)
		bytes = append(bytes, b)
	}
	fmt.Println("Slice Size is : ", len(bytes))
	page.Bytes = bytes
	Write(page)
}

func testRead(pageNum int) {
	page := Read(pageNum)
	fmt.Printf("Page number is : %v", page.Num)
}
