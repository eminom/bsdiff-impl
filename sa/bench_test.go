package sa

import (
	"crypto/rand"
	mrand "math/rand"
	"sync"
	"testing"
	"time"
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

var doOnce sync.Once

var inputStream [][]byte

func initInput() {
	doOnce.Do(func() {
		for i := 0; i < 10; i++ {
			buffer := make([]byte, 16*1024+mrand.Intn(256*1024))
			rand.Read(buffer)
			inputStream = append(inputStream, buffer)
		}
	})
}

func BenchmarkBsdiffUno(b *testing.B) {
	a := new(StdDoubleAlgo)
	for i := 0; i < b.N; i++ {
		for _, v := range inputStream {
			a.Sort(v)
		}
	}

}

func BenchmarkBsdiffCero(b *testing.B) {
	a := new(ModifiedSort)
	for i := 0; i < b.N; i++ {
		for _, v := range inputStream {
			a.Sort(v)
		}
	}
}
