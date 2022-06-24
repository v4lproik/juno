package trie

var EmptyPath = &Path{0, []byte{}}

type Path struct {
	length int
	bytes  []byte
}

func NewPath(length int, b []byte) *Path {
	bs := &Path{length, make([]byte, (length+7)/8)}
	offset := len(b) - len(bs.bytes)
	if offset < 0 {
		copy(bs.bytes[-offset:], b)
	} else {
		copy(bs.bytes, b[offset:])
	}
	for i := 1; i <= len(bs.bytes)*8-bs.length; i++ {
		bs.Clear(-i)
	}
	return bs
}

func (bs *Path) Set(i int) {
	i += len(bs.bytes)*8 - bs.length
	bs.bytes[i/8] |= 1 << (7 - i%8)
}

func (bs *Path) Clear(i int) {
	i += len(bs.bytes)*8 - bs.length
	bs.bytes[i/8] &^= (1 << (7 - i%8))
}

func (bs *Path) Get(i int) bool {
	i += len(bs.bytes)*8 - bs.length
	return bs.bytes[i/8]&(1<<(7-i%8)) != 0
}

func (bs *Path) Len() int {
	return bs.length
}

func (bs *Path) Bytes() []byte {
	return NewPath(bs.length, bs.bytes).bytes
}

func (path *Path) Walked(walked int) *Path {
	return NewPath(path.Len()-walked, path.bytes)
}

func (path *Path) Prefix(length int) *Path {
	result := NewPath(length, []byte{})
	for i := 0; i < length; i++ {
		if path.Get(i) {
			result.Set(i)
		}
	}
	return result
}

func (path *Path) String() string {
	res := ""
	for i := 0; i < path.Len(); i++ {
		if path.Get(i) {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}

func (path *Path) LongestCommonPrefix(other *Path) int {
	n := 0
	for ; n < path.Len() && n < other.Len(); n++ {
		if path.Get(n) != other.Get(n) {
			break
		}
	}
	return n
}