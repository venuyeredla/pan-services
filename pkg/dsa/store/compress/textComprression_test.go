package compress

import (
	"fmt"
	"os"
	"testing"
)

func TestHCompression(t *testing.T) {
	t.Skip()
	//fileName := "C:\\Work\\opensource\\data\\sample.txt"
	fileName := "./wtest.txt"

	file, err := os.Open(fileName)
	if err == nil {
		var b []byte = make([]byte, 310, os.Getpagesize())
		fmt.Println("Size:", len(b), " Capacity:", cap(b))
		size, err := file.Read(b)
		fmt.Println("Opened file & no of bytes read :", size)
		if err == nil {
			compressed := Hcompress(b)
			Hdecode(compressed)
		} else {
			fmt.Println("Error in opening file", fileName)
		}

	} else {
		fmt.Println("Error in opening file", fileName)
	}
}

//vcomp.AbsCompress("00101000101110101001010")
//symDict := vcomp.Count(bytes)
//testCompression()
//vcomp.BitReadWriteTest()
