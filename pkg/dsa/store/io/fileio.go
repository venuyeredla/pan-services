package io

import (
	"log"
	"os"
)

/*
Reading/writing whole file is ReadFile() or WriteFile().
*/
const file_name string = "/Users/venugopal/Documents/work/Gapp/tdata/test.txt"

func ReadFile() {
	data, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)

}

func WriteFile() {

	text := "Learing go language is good for futuure."

	err := os.WriteFile(file_name, []byte(text), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func RandomRead() {
	f, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 15)
	_, er := f.ReadAt(data, 2)
	if er != nil {
		log.Fatal(er)
	}

	os.Stdout.Write(data)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func RandomWrite() {
	f, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	appendText := "Written by random write."
	wcount, werr := f.WriteAt([]byte(appendText), 40)
	if werr != nil {
		log.Fatal(werr)
	}
	log.Printf("Written : %v", wcount)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
