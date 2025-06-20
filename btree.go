package btree

// author:wlb github/runsys email:iwlb@outlook.com create at:20250620T151011
// binary tree implement.
type Btree[T any] struct {
	v     T
	left  *Btree[T]
	right *Btree[T]
}

func (bt *Btree[T]) Put(cmp func(a, b T) int, v T) bool {
	cur := bt
	for {
		cmprl := cmp(cur.v, v)
		if cmprl < 0 {
			if cur.right != nil {
				cur = cur.right
			} else {
				cur.right = &Btree[T]{v: v}
				return true
			}
		} else if cmprl > 0 {
			if cur.left != nil {
				cur = cur.left
			} else {
				cur.left = &Btree[T]{v: v}
				return true
			}
		} else {
			cur.v = v
			return true
		}
	}
	return false
}

func (bt *Btree[T]) Get(cmp func(a, b T) int, fd T) (nd *Btree[T]) {
	if bt == nil {
		return nd
	}
	cur := bt
	for {
		cmprl := cmp(cur.v, fd)
		if cmprl < 0 {
			if cur.right != nil {
				cur = cur.right
			} else {
				return nd
			}
		} else if cmprl > 0 {
			if cur.left != nil {
				cur = cur.left
			} else {
				return nd
			}
		} else {
			return cur
		}
	}
	return nd
}

//return new root node;
func (bt *Btree[T]) Del(cmp func(a, b T) int, fd T) (rt *Btree[T]) {
	if bt == nil {
		return nil
	}
	var cpl, cpr *Btree[T]
	cur := bt
	for {
		cmprl := cmp(cur.v, fd)
		if cmprl < 0 {
			if cur.right != nil {
				cpl = nil
				cpr = cur
				cur = cur.right
			} else {
				return bt
			}
		} else if cmprl > 0 {
			if cur.left != nil {
				cpl = cur
				cpr = nil
				cur = cur.left
			} else {
				return bt
			}
		} else {
			if cur.left == nil && cur.right != nil {
				if cpl != nil {
					cpl.left = cur.right
				} else if cpr != nil {
					cpr.right = cur.right
				} else {
					return cur.right
				}
			} else if cur.left != nil && cur.right == nil {
				if cpl != nil {
					cpl.left = cur.left
				} else if cpr != nil {
					cpr.right = cur.left
				} else {
					return cur.left
				}
			} else if cur.left == nil && cur.right == nil {
				if cpl != nil {
					cpl.left = nil
				} else if cpr != nil {
					cpr.right = nil
				} else {
					return nil
				}
			} else {
				cu := cur.left
				for cu.right != nil {
					cu = cu.right
				}
				cu.right = cur.right.left
				cur.right.left = cur.left
				if cpl != nil {
					cpl.left = cur.right
				} else if cpr != nil {
					cpr.right = cur.right
				} else {
					return cur.right
				}
			}
			return bt
		}
	}
	return bt
}

type Btiter[T any] struct {
	bc  []*Btree[T] //Btree chain
	cur *Btree[T]
	fd  T
	cmp func(a, b T) int
}

func NewBtiter[T any](bc []*Btree[T], fd T, cmp func(a, b T) int) (bi *Btiter[T]) {
	bi = &Btiter[T]{bc: bc, fd: fd, cmp: cmp}
	return bi
}

func (bi *Btiter[T]) Next() bool {
	if len(bi.bc) == 0 {
		return false
	}
	if bi.cmp(bi.bc[len(bi.bc)-1].v, bi.fd) < 0 {
		return false
	}
	bi.cur = bi.bc[len(bi.bc)-1]
	if bi.bc[len(bi.bc)-1].right == nil {
		bi.bc = bi.bc[:len(bi.bc)-1]
		return true
	} else if bi.bc[len(bi.bc)-1].right != nil {
		n := len(bi.bc)
		cu := bi.bc[len(bi.bc)-1].right
		bi.bc = append(bi.bc, cu)
		for cu.left != nil {
			bi.bc = append(bi.bc, cu.left)
			cu = cu.left
		}
		bi.bc = append(bi.bc[:n-1], bi.bc[n:]...)
		return true
	}
	return false
}

func (bi *Btiter[T]) Value() (v *Btree[T]) {
	return bi.cur
}

// if fd==root Btiterator all;
func (bt *Btree[T]) Newiter(cmp func(a, b T) int, fd T) (bi *Btiter[T]) {
	cur := bt
	for cur != nil {
		cr := cmp(cur.v, fd)
		if cr < 0 {
			if cur.right != nil {
				cur = cur.right
			} else {
				break
			}
		} else if cr > 0 {
			if cur.left != nil {
				cur = cur.left
			} else {
				break
			}
		} else {
			break
		}
	}
	cu := cur
	bc := []*Btree[T]{}
	bc = append(bc, cu)
	for cu.left != nil {
		bc = append(bc, cu.left)
		cu = cu.left
	}
	return NewBtiter(bc, fd, cmp)
}

type Btbiter[T any] struct {
	bc  []*Btree[T] //Btree chain
	cur *Btree[T]
	fd  T
	cmp func(a, b T) int
}

func NewBtbiter[T any](bc []*Btree[T], fd T, cmp func(a, b T) int) (bi *Btbiter[T]) {
	bi = &Btbiter[T]{bc: bc, fd: fd, cmp: cmp}
	return bi
}

func (bi *Btbiter[T]) Next() bool {
	if len(bi.bc) == 0 {
		return false
	}
	if bi.cmp(bi.bc[len(bi.bc)-1].v, bi.fd) > 0 {
		return false
	}
	bi.cur = bi.bc[len(bi.bc)-1]
	if bi.bc[len(bi.bc)-1].left == nil {
		bi.bc = bi.bc[:len(bi.bc)-1]
		return true
	} else if bi.bc[len(bi.bc)-1].left != nil {
		n := len(bi.bc)
		cu := bi.bc[len(bi.bc)-1].left
		bi.bc = append(bi.bc, cu)
		for cu.right != nil {
			bi.bc = append(bi.bc, cu.right)
			cu = cu.right
		}
		bi.bc = append(bi.bc[:n-1], bi.bc[n:]...)
		return true
	}
	return false
}

func (bi *Btbiter[T]) Value() (v *Btree[T]) {
	return bi.cur
}

// back iter
func (bt *Btree[T]) Newbiter(cmp func(a, b T) int, fd T) (bi *Btbiter[T]) {
	cur := bt
	bc := []*Btree[T]{}
	bc = append(bc, cur)
	for cur != nil {
		cr := cmp(cur.v, fd)
		if cr < 0 {
			if cur.right != nil {
				cur = cur.right
				bc = append(bc, cur)
			} else {
				break
			}
		} else if cr > 0 {
			if cur.left != nil {
				cur = cur.left
				bc = append(bc, cur)
			} else {
				break
			}
		} else {
			break
		}
	}
	return NewBtbiter(bc, fd, cmp)
}
