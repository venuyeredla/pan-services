package compress

import (
	"fmt"
)

const Code_valu_bits byte = 8
const MAX_FREQ uint16 = 64

type SymbolStats struct {
	symbol byte
	freq   uint16
	high   uint16
	low    uint16
}

var symArr []*SymbolStats

func ACompress(data []byte) {
	buildModel(data)
}

func buildModel(data []byte) {
	var symMap map[byte]uint16 = make(map[byte]uint16)
	for _, b := range data {
		val, present := symMap[b]
		if present == false {
			symMap[b] = 1
		} else {
			symMap[b] = val + 1
		}
	}
	keys := make([]byte, 0)
	for key := range symMap {
		keys = append(keys, key)
	}
	sortBytes(keys)
	for _, v := range keys {
		val := symMap[v]
		sym := &SymbolStats{symbol: v, freq: val}
		symArr = append(symArr, sym)
	}
	for _, s := range symArr {
		fmt.Print(string(s.symbol), "-", s.freq, ",")
	}
}
