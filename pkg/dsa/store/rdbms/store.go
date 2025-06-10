// vdb
package vrdbms

import (
	"fmt"
	"log"
	"os"
)

type Page struct {
	Num   int
	Bytes []byte
}

func NewPage(pageNum int) *Page {
	return &Page{
		Num:   pageNum,
		Bytes: make([]byte, 0, os.Getpagesize()),
	}
}

const fileName = "/home/venugopal/Documents/Code/GoData/go.db"

func dbExists() bool {
	_, err := os.Open(fileName)
	if err == nil {
		return true
	} else {
		return false
	}

}

func Write(page *Page) bool {
	file := getFile(fileName)
	size, err := file.WriteAt(page.Bytes, page.offSet())
	file.Close()
	if err == nil {
		log.Println("No of bytes written", size, "in page", page.Num)
		return true
	} else {
		return false
	}
}

func Read(num int) *Page {
	file := getFile(fileName)
	page := new(Page)
	page.Num = num
	size, err := file.ReadAt(page.Bytes, page.offSet())
	if err == nil {
		log.Println("No of bytes read: ", size, "from page :", page.Num)
	} else {
		log.Println("Error in reading data from file : page num ", page.Num, err)
	}
	fmt.Println("Hello from VIO package!")
	return page
}

func getFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Properties", file.Name())
	return file
}

func (page *Page) offSet() int64 {
	offset := os.Getpagesize() * page.Num
	return int64(offset)
}

func (page *Page) writeByte(b byte) {

}
