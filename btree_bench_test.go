package btree

import (
	"math/rand"
	"testing"
)

func BenchmarkPut(t *testing.B) {
	pq(nfn())
	var bt *Btree[int]
	for i := 0; i < 10000000; i += 1 {
		v := rand.Int()
		if bt == nil {
			bt = &Btree[int]{v: v}
		} else {
			bt.Put(Icmp, v)
		}
	}
	pq(nfn())
	bi := bt.Newiter(Icmp, Imi)
	i := 0
	for bi.Next() {
		i += 1
	}
	pq("count", i, "time", nfn())
}
