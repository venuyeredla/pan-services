// vcomp project vcomp.go
package compress

import (
	"fmt"
)

type HuffmanNode struct {
	key         byte
	freq        uint16
	left, right *HuffmanNode
}

var hNodes []*HuffmanNode
var heapSize int

/*
*
Huffman coding compression
*/
func Hcompress(bytes []byte) []byte {
	fmt.Println("Actual data size:", len(bytes))
	bitStream := &BitStream{}
	symMap := buildSymMap(bytes, bitStream)
	codeMap := buildHCodes(symMap)
	bitStream.WriteUint16(uint16(len(bytes)))
	for _, b := range bytes {
		code := codeMap[b]
		for _, c := range code {
			bitStream.WriteBit(byte(c))
		}
	}
	bitStream.close()
	fmt.Println("\nCompressed size :", bitStream.size())
	fmt.Println("")
	return bitStream.getBytes()
}

func Hdecode(compressed []byte) {
	bitStream := &BitStream{}
	bitStream.bytes = compressed
	keySize := int(bitStream.ReadUint16())
	var symMap map[byte]uint16 = make(map[byte]uint16)
	for i := 0; i < keySize; i++ {
		symMap[bitStream.readByte()] = bitStream.ReadUint16()
	}
	codeMap := buildHCodes(symMap)

	fmt.Println("CodeMap size:", len(codeMap))
	dataSize := int(bitStream.ReadUint16())
	var decodeMap map[string]byte = make(map[string]byte)
	for k, v := range codeMap {
		decodeMap[v] = k
		fmt.Print(v, "-", k, " , ")
	}
	code := ""
	for decoded := 0; decoded <= dataSize; decoded++ {
		_, bit := bitStream.ReadBit()
		code = code + string(bit)
		_, exist := decodeMap[code]
		fmt.Println("code", code)
		for exist == false {
			_, bit = bitStream.ReadBit()
			code = code + string(bit)
			fmt.Println("code", code)
			_, exist = decodeMap[code]
		}
		//break
		rcode, _ := decodeMap[code]
		fmt.Printf("%c", rcode)
		code = ""

	}

}

func (hnode *HuffmanNode) toString() {
	fmt.Println(hnode.key, "(", string(hnode.key), ")-", hnode.freq, hnode.left, " , ", hnode.right)
}

func buildSymMap(data []byte, bitStream *BitStream) map[byte]uint16 {
	var symMap map[byte]uint16 = make(map[byte]uint16)
	for _, b := range data {
		val, present := symMap[b]
		if present == false {
			symMap[b] = 1
		} else {
			symMap[b] = val + 1
		}
	}
	bitStream.WriteUint16(uint16(len(symMap)))
	for k, v := range symMap {
		bitStream.writeByte(k)
		bitStream.WriteUint16(v)
	}
	return symMap
}

func buildHCodes(symMap map[byte]uint16) map[byte]string {
	buildMinHeap(symMap)
	buildHuffManTree()
	codeMap := make(map[byte]string, len(symMap))
	createCodeMap(codeMap, hNodes[0], "")
	for k, v := range codeMap {
		fmt.Print(k, "-", v, ",")
	}
	return codeMap
}

func buildMinHeap(symMap map[byte]uint16) {
	size := len(symMap)
	hNodes = make([]*HuffmanNode, size)
	keys := make([]byte, 0)
	for key, _ := range symMap {
		keys = append(keys, key)
	}
	sortBytes(keys)
	heapSize = -1
	for _, v := range keys {
		val := symMap[v]
		add(&HuffmanNode{key: v, freq: val})
	}
	printArr()
	fmt.Println("Codes")
}

func add(hNode *HuffmanNode) {
	heapSize = heapSize + 1
	hNodes[heapSize] = hNode
	var i int = heapSize
	for i != 0 && validateHeapfy((i-1)/2, i) { // (i-1)/2 gives the parent of node
		swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func buildHuffManTree() {
	for heapSize > 0 {
		//fmt.Println("current heapSize & indexex :", len(hNodes), " - ", heapSize)
		first := extractMin()
		second := extractMin()
		cumFreq := first.freq + second.freq
		node := &HuffmanNode{freq: cumFreq}
		node.left = first
		node.right = second
		//fmt.Print("First(", string(first.key), ",", first.freq, "),  Second(", string(second.key), ",", second.freq, "),  ")
		//node.toString()
		//fmt.Println("Inserting node at :", heapSize)
		add(node)
	}
}

func extractMin() *HuffmanNode {
	temp := hNodes[0]
	hNodes[0] = hNodes[heapSize]
	hNodes[heapSize] = nil
	heapSize = heapSize - 1
	minHeapfy(0)
	//printArr()
	return temp
}

func printArr() {
	fmt.Print("Heap size: ", len(hNodes), "=")
	for i := range hNodes {
		n := hNodes[i]
		fmt.Print(n.key, "-", n.freq, " ,")
	}
}

// Heafies from top to bottom.
func minHeapfy(i int) {
	l := 2*i + 1
	r := 2*i + 2
	smallest := i
	if l < heapSize && validateHeapfy(smallest, l) {
		smallest = l
	}
	if r < heapSize && validateHeapfy(smallest, r) {
		smallest = r
	}
	if smallest != i {
		swap(i, smallest)
		minHeapfy(smallest)
	}
}

func createCodeMap(codeMap map[byte]string, node *HuffmanNode, str string) {
	if node != nil {
		if node.left != nil {
			createCodeMap(codeMap, node.left, str+"0")
		}
		if node.right != nil {
			createCodeMap(codeMap, node.right, str+"1")
		}
		if node.left == nil && node.right == nil {
			codeMap[node.key] = str
		}
	}
}

func validateHeapfy(j int, i int) bool {
	parent := hNodes[j]
	child := hNodes[i]
	return parent.freq > child.freq
}

func swap(first int, second int) {
	temp := hNodes[first]
	hNodes[first] = hNodes[second]
	hNodes[second] = temp
}

func sortBytes(b []byte) {
	var i, j, min_idx int
	n := len(b)
	for i = 0; i < n-1; i = i + 1 {
		min_idx = i
		for j = i + 1; j < n; j = j + 1 {
			if b[j] < b[min_idx] {
				min_idx = j
			}
		}
		temp := b[min_idx]
		b[min_idx] = b[i]
		b[i] = temp
	}
	/*
		for _, v := range b {
			fmt.Print(v, ",")
		}
	*/
}
