package utils

import (
	crand "crypto/rand"
	"encoding/hex"
	"math/rand"
	"sync"
	"testing"
)

var testrandom *rand.Rand

// ThreadSafeRand provides a thread safe version of math/rand.Rand using
// the same technique used in the math/rand package to make the top level
// functions thread safe.
func ThreadSafeRand(seed int64) *rand.Rand {
	temp, _ := rand.NewSource(seed).(rand.Source64)
	return rand.New(&lockedSource{src: &temp})
}

type T testing.T

type lockedSource struct {
	lk  sync.Mutex
	src *rand.Source64
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.Int63()
	//n=r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Uint64() (n uint64) {
	r.lk.Lock()
	//n = r.src.Uint64()
	n = r.Uint64()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	// r.src.Seed(seed)
	r.Seed(seed)
	r.lk.Unlock()
}

// seedPos implements Seed for a lockedSource without a race condiiton.
func (r *lockedSource) seedPos(seed int64, readPos *int8) {
	r.lk.Lock()
	// r.src.Seed(seed)
	r.Seed(seed)
	*readPos = 0
	r.lk.Unlock()
}

// read implements Read for a lockedSource without a race condition.
func (r *lockedSource) read(p []byte, readVal *int64, readPos *int8) (n int, err error) {
	r.lk.Lock()
	// n, err = read(p, r.src.Int63, readVal, readPos)
	n, err = read(p, r.Int63, readVal, readPos)
	r.lk.Unlock()
	return
}

func read(p []byte, int63 func() int64, readVal *int64, readPos *int8) (n int, err error) {
	pos := *readPos
	val := *readVal
	for n = 0; n < len(p); n++ {
		if pos == 0 {
			val = int63()
			pos = 7
		}
		p[n] = byte(val)
		val >>= 8
		pos--
	}
	*readPos = pos
	*readVal = val
	return
}

func RandSlice(length int) []byte {
	slice := make([]byte, length)
	if _, err := crand.Read(slice); err != nil {
		panic(err)
	}
	return slice
}

func RandHex(length int) string {
	return hex.EncodeToString(RandSlice(length / 2))
}

func RandStr(length int) string {
	return string(RandSlice(length))
}
