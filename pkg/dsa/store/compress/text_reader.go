package compress

import (
	"fmt"
	"os"
)

func getText() []byte {
	fileName := "/home/venugopal/Documents/Code/GoData/sample.txt"
	file, err := os.Open(fileName)
	if err == nil {
		var b []byte = make([]byte, 310, os.Getpagesize())
		fmt.Println("Size:", len(b), " Capacity:", cap(b))
		size, _ := file.Read(b)
		fmt.Println("Number of bytes read", size)
		return b
	} else {
		fmt.Println("Error in opening file", fileName)
		return nil
	}
}
