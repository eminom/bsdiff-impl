package sa

import (
	"bytes"
	"crypto/rand"
	mrand "math/rand"
	"testing"
	"time"
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

func randString(n int) string {
	rv := ""
	for i := 0; i < n; i++ {
		rv += string(mrand.Intn(2) + 'a')
	}
	return rv
}

func TestBsdiffBasic(t *testing.T) {
	testBsdiffX(new(RawSort), 1000, t)
	testBinaries(new(RawSort), 100, t)
}

func TestBsdiffCero(t *testing.T) {
	testBsdiffX(new(ModifiedSort), 1000, t)
	testBinaries(new(ModifiedSort), 100, t)
}

func TestBsdiffUno(t *testing.T) {
	testBsdiffX(new(StdDoubleAlgo), 1000, t)
	testBinaries(new(StdDoubleAlgo), 100, t)
}

func testBsdiffX(sa SASort, kRun int, t *testing.T) {
	for i := 0; i < 16*1024; i = i*2 + 1 {
		for k := 0; k < kRun; k++ {
			buff := []byte(randString((i + k)))
			v := sa.Sort(buff)
			for j := 1; j < len(v); j++ {
				if -1 != bytes.Compare(buff[v[j-1]:], buff[v[j]:]) {
					t.Logf("Test on <%v> failed", string(buff))
					t.Logf("result: %v", v)
					t.Logf("[%v]: %v", v[j-1], string(buff[v[j-1]:]))
					t.Logf("[%v]: %v", v[j], string(buff[v[j]:]))
					t.Fatalf("No no no: <%v>", string(buff))
				}
			}
		}
	}
}

func testBinaries(sa SASort, kRun int, t *testing.T) {
	var buff [1024 * 16]byte
	for i := 0; i < 1000; i++ {
		l := mrand.Intn(1024 * 16)
		m, _ := rand.Read(buff[:l])
		v := sa.Sort(buff[:m])
		for j := 1; j < len(v); j++ {
			if -1 != bytes.Compare(buff[v[j-1]:m], buff[v[j]:m]) {
				t.Fatalf("failed for %v", sa)
			}
		}
	}
}
