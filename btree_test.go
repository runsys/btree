package btree

import (
	"math/rand"
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	is1 := []int{}
	is2 := []int{}
	var bt *Btree[int]
	m := make(map[int]bool)
	for i := 0; i < 100000; i += 1 {
		v := rand.Int()
		if m[v] == false {
			m[v] = true
			is1 = append(is1, v)
			if bt == nil {
				bt = &Btree[int]{v: v}
			} else {
				bt.Put(Icmp, v)
			}
			gv := bt.Get(Icmp, v)
			if gv == nil {
				E("error", v, i)
			} else if gv.v != v {
				E("error", v, i)
			}
		}
	}
	sort.Ints(is1)
	bi := bt.Newiter(Icmp, Imi)
	i := 0
	for bi.Next() {
		v := bi.cur.v
		if is1[i] != v {
			E("error", i, is1[i], v)
		}
		i += 1
	}
	bbi := bt.Newbiter(Icmp, Ima)
	i = len(is1) - 1
	for bbi.Next() {
		v := bbi.cur.v
		if is1[i] != v {
			E("error", i, is1[i], v, len(is1))
		}
		i -= 1
	}

	dec := 0
	for k, _ := range m {
		delete(m, k)
		bt = bt.Del(Icmp, k)
		dec += 1
		if dec >= 10000 {
			break
		}
	}
	for k, _ := range m {
		is2 = append(is2, k)
	}
	sort.Ints(is2)
	bi = bt.Newiter(Icmp, Imi)
	i = 0
	for bi.Next() {
		v := bi.cur.v
		if is2[i] != v {
			E("error", i, is2[i], v)
		}
		i += 1
	}
	bbi = bt.Newbiter(Icmp, Ima)
	i = len(is2) - 1
	for bbi.Next() {
		v := bbi.cur.v
		if is2[i] != v {
			E("error", i, is2[i], v, len(is2))
		}
		i -= 1
	}

}
