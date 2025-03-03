package dist

import (
	"bytes"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/lib"
)

func TestDecodeDistHeaderAtomCache(t *testing.T) {
	c := &distConnection{}
	c.cache = etf.NewAtomCache()
	a1 := etf.Atom("atom1")
	a2 := etf.Atom("atom2")
	c.cache.In.Atoms[1034] = &a1
	c.cache.In.Atoms[5] = &a2
	packet := []byte{
		131, 68, // start dist header
		5, 4, 137, 9, // 5 atoms and theirs flags
		10, 5, // already cached atom ids
		236, 3, 114, 101, 103, // atom 'reg'
		9, 4, 99, 97, 108, 108, //atom 'call'
		238, 13, 115, 101, 116, 95, 103, 101, 116, 95, 115, 116, 97, 116, 101, // atom 'set_get_state'
		104, 4, 97, 6, 103, 82, 0, 0, 0, 0, 85, 0, 0, 0, 0, 2, 82, 1, 82, 2, // message...
		104, 3, 82, 3, 103, 82, 0, 0, 0, 0, 245, 0, 0, 0, 2, 2,
		104, 2, 82, 4, 109, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	cacheExpected := []etf.Atom{"atom1", "atom2", "reg", "call", "set_get_state"}
	cacheInExpected := c.cache.In.Atoms
	a3 := etf.Atom("reg")
	a4 := etf.Atom("call")
	a5 := etf.Atom("set_get_state")
	cacheInExpected[492] = &a3
	cacheInExpected[9] = &a4
	cacheInExpected[494] = &a5

	packetExpected := packet[34:]
	cache, packet1, _ := c.decodeDistHeaderAtomCache(packet[2:], nil)

	if !bytes.Equal(packet1, packetExpected) {
		t.Fatal("incorrect packet")
	}

	if !reflect.DeepEqual(c.cache.In.Atoms, cacheInExpected) {
		t.Fatal("incorrect cacheIn")
	}

	if !reflect.DeepEqual(cache, cacheExpected) {
		t.Fatal("incorrect cache", cache)
	}

}

func TestEncodeDistHeaderAtomCache(t *testing.T) {

	b := lib.TakeBuffer()
	defer lib.ReleaseBuffer(b)

	senderAtomCache := make(map[etf.Atom]etf.CacheItem)
	encodingAtomCache := etf.TakeEncodingAtomCache()
	defer etf.ReleaseEncodingAtomCache(encodingAtomCache)

	senderAtomCache["reg"] = etf.CacheItem{ID: 1000, Encoded: false, Name: "reg"}
	senderAtomCache["call"] = etf.CacheItem{ID: 499, Encoded: false, Name: "call"}
	senderAtomCache["one_more_atom"] = etf.CacheItem{ID: 199, Encoded: true, Name: "one_more_atom"}
	senderAtomCache["yet_another_atom"] = etf.CacheItem{ID: 2, Encoded: false, Name: "yet_another_atom"}
	senderAtomCache["extra_atom"] = etf.CacheItem{ID: 10, Encoded: true, Name: "extra_atom"}
	senderAtomCache["potato"] = etf.CacheItem{ID: 2017, Encoded: true, Name: "potato"}

	// Encoded field is ignored here
	encodingAtomCache.Append(etf.CacheItem{ID: 499, Name: "call"})
	encodingAtomCache.Append(etf.CacheItem{ID: 1000, Name: "reg"})
	encodingAtomCache.Append(etf.CacheItem{ID: 199, Name: "one_more_atom"})
	encodingAtomCache.Append(etf.CacheItem{ID: 2017, Name: "potato"})

	expected := []byte{
		4, 185, 112, 0, // 4 atoms and theirs flags
		243, 4, 99, 97, 108, 108, // atom call
		232, 3, 114, 101, 103, // atom reg
		199, // atom one_more_atom, already encoded
		225, // atom potato, already encoded

	}

	l := &distConnection{}
	l.encodeDistHeaderAtomCache(b, senderAtomCache, encodingAtomCache)

	if !reflect.DeepEqual(b.B, expected) {
		t.Fatal("incorrect value")
	}

	b.Reset()
	encodingAtomCache.Append(etf.CacheItem{ID: 2, Name: "yet_another_atom"})

	expected = []byte{
		5, 49, 112, 8, // 5 atoms and theirs flags
		243,                      // atom call. already encoded
		232,                      // atom reg. already encoded
		199,                      // atom one_more_atom. already encoded
		225,                      // atom potato. already encoded
		2, 16, 121, 101, 116, 95, // atom yet_another_atom
		97, 110, 111, 116, 104, 101,
		114, 95, 97, 116, 111, 109,
	}
	l.encodeDistHeaderAtomCache(b, senderAtomCache, encodingAtomCache)

	if !reflect.DeepEqual(b.B, expected) {
		t.Fatal("incorrect value", b.B)
	}
}

func BenchmarkDecodeDistHeaderAtomCache(b *testing.B) {
	link := &distConnection{}
	packet := []byte{
		131, 68, // start dist header
		5, 4, 137, 9, // 5 atoms and theirs flags
		10, 5, // already cached atom ids
		236, 3, 114, 101, 103, // atom 'reg'
		9, 4, 99, 97, 108, 108, //atom 'call'
		238, 13, 115, 101, 116, 95, 103, 101, 116, 95, 115, 116, 97, 116, 101, // atom 'set_get_state'
		104, 4, 97, 6, 103, 82, 0, 0, 0, 0, 85, 0, 0, 0, 0, 2, 82, 1, 82, 2, // message...
		104, 3, 82, 3, 103, 82, 0, 0, 0, 0, 245, 0, 0, 0, 2, 2,
		104, 2, 82, 4, 109, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		link.decodeDistHeaderAtomCache(packet[2:], nil)
	}
}

func BenchmarkEncodeDistHeaderAtomCache(b *testing.B) {
	link := &distConnection{}
	buf := lib.TakeBuffer()
	defer lib.ReleaseBuffer(buf)

	senderAtomCache := make(map[etf.Atom]etf.CacheItem)
	encodingAtomCache := etf.TakeEncodingAtomCache()
	defer etf.ReleaseEncodingAtomCache(encodingAtomCache)

	senderAtomCache["reg"] = etf.CacheItem{ID: 1000, Encoded: false, Name: "reg"}
	senderAtomCache["call"] = etf.CacheItem{ID: 499, Encoded: false, Name: "call"}
	senderAtomCache["one_more_atom"] = etf.CacheItem{ID: 199, Encoded: true, Name: "one_more_atom"}
	senderAtomCache["yet_another_atom"] = etf.CacheItem{ID: 2, Encoded: false, Name: "yet_another_atom"}
	senderAtomCache["extra_atom"] = etf.CacheItem{ID: 10, Encoded: true, Name: "extra_atom"}
	senderAtomCache["potato"] = etf.CacheItem{ID: 2017, Encoded: true, Name: "potato"}

	// Encoded field is ignored here
	encodingAtomCache.Append(etf.CacheItem{ID: 499, Name: "call"})
	encodingAtomCache.Append(etf.CacheItem{ID: 1000, Name: "reg"})
	encodingAtomCache.Append(etf.CacheItem{ID: 199, Name: "one_more_atom"})
	encodingAtomCache.Append(etf.CacheItem{ID: 2017, Name: "potato"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		link.encodeDistHeaderAtomCache(buf, senderAtomCache, encodingAtomCache)
	}
}

func TestDecodeFragment(t *testing.T) {
	link := &distConnection{}

	link.checkCleanTimeout = 50 * time.Millisecond
	link.checkCleanDeadline = 150 * time.Millisecond
	link.fragments = make(map[uint64]*fragmentedPacket)

	// decode fragment with fragmentID=0 should return error
	fragment0 := []byte{protoDistFragment1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3}
	if _, e := link.decodeFragment(fragment0, nil); e == nil {
		t.Fatal("should be error here")
	}

	fragment1 := []byte{protoDistFragment1, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 3, 0, 1, 2, 3}
	fragment2 := []byte{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 2, 4, 5, 6}
	fragment3 := []byte{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 1, 7, 8, 9}

	expected := []byte{68, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// add first fragment
	if x, e := link.decodeFragment(fragment1, nil); x != nil || e != nil {
		t.Fatal("should be nil here", x, e)
	}
	// add second one
	if x, e := link.decodeFragment(fragment2, nil); x != nil || e != nil {
		t.Fatal("should be nil here", e)
	}
	// add the last one. should return *lib.Buffer with assembled packet
	if x, e := link.decodeFragment(fragment3, nil); x == nil || e != nil {
		t.Fatal("shouldn't be nil here", e)
	} else {
		// x should be *lib.Buffer
		if !reflect.DeepEqual(expected, x.B) {
			t.Fatal("exp:", expected, "got:", x.B)
		}
		lib.ReleaseBuffer(x)

		// map of the fragments should be empty here
		if len(link.fragments) > 0 {
			t.Fatal("fragments should be empty")
		}
	}

	link.checkCleanTimeout = 50 * time.Millisecond
	link.checkCleanDeadline = 150 * time.Millisecond
	// test lost fragment
	// add the first fragment and wait 160ms
	if x, e := link.decodeFragment(fragment1, nil); x != nil || e != nil {
		t.Fatal("should be nil here", e)
	}
	if len(link.fragments) == 0 {
		t.Fatal("fragments should have a record ")
	}
	// check and clean process should remove this record
	time.Sleep(360 * time.Millisecond)

	// map of the fragments should be empty here
	if len(link.fragments) > 0 {
		t.Fatal("fragments should be empty")
	}

	link.checkCleanTimeout = 0
	link.checkCleanDeadline = 0
	fragments := [][]byte{
		{protoDistFragment1, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 9, 0, 1, 2, 3},
		{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 8, 4, 5, 6},
		{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 7, 7, 8, 9},
		{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 6, 10, 11, 12},
		{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 5, 13, 14, 15},
		{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 4, 16, 17, 18},
		{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 3, 19, 20, 21},
		{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 2, 22, 23, 24},
		{protoDistFragmentN, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 1, 25, 26, 27},
	}
	expected = []byte{68, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}

	fragmentsReverse := make([][]byte, len(fragments))
	l := len(fragments)
	for i := 0; i < l; i++ {
		fragmentsReverse[l-i-1] = fragments[i]
	}

	var result *lib.Buffer
	var e error
	for i := range fragmentsReverse {
		if result, e = link.decodeFragment(fragmentsReverse[i], nil); e != nil {
			t.Fatal(e)
		}

	}
	if result == nil {
		t.Fatal("got nil result")
	}
	if !reflect.DeepEqual(expected, result.B) {
		t.Fatal("exp:", expected, "got:", result.B[:len(expected)])
	}
	// map of the fragments should be empty here
	if len(link.fragments) > 0 {
		t.Fatal("fragments should be empty")
	}

	// reshuffling 100 times
	for k := 0; k < 100; k++ {
		result = nil
		fragmentsShuffle := make([][]byte, l)
		rand.Seed(time.Now().UnixNano())
		for i, v := range rand.Perm(l) {
			fragmentsShuffle[v] = fragments[i]
		}

		for i := range fragmentsShuffle {
			if result, e = link.decodeFragment(fragmentsShuffle[i], nil); e != nil {
				t.Fatal(e)
			}

		}
		if result == nil {
			t.Fatal("got nil result")
		}
		if !reflect.DeepEqual(expected, result.B) {
			t.Fatal("exp:", expected, "got:", result.B[:len(expected)])
		}
	}
}
