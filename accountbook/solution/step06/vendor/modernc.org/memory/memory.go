// Copyright 2017 The Memory Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package memory implements a memory allocator.
//
// Changelog
//
// 2017-10-03 Added alternative, unsafe.Pointer-based API.
//
// Benchmarks
//
// Intel® Core™ i5-4670 CPU @ 3.40GHz × 4
//
//  goos: linux
//  goarch: amd64
//  pkg: modernc.org/memory
//  BenchmarkFree16-4           	100000000	        15.3 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkFree32-4           	100000000	        21.3 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkFree64-4           	50000000	        35.9 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkCalloc16-4         	50000000	        26.6 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkCalloc32-4         	50000000	        30.1 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkCalloc64-4         	30000000	        38.1 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkGoCalloc16-4       	50000000	        29.3 ns/op	      16 B/op	       1 allocs/op
//  BenchmarkGoCalloc32-4       	50000000	        30.4 ns/op	      32 B/op	       1 allocs/op
//  BenchmarkGoCalloc64-4       	30000000	        37.9 ns/op	      64 B/op	       1 allocs/op
//  BenchmarkMalloc16-4         	100000000	        15.4 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkMalloc32-4         	100000000	        15.6 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkMalloc64-4         	100000000	        15.9 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkUnsafeFree16-4     	100000000	        14.4 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkUnsafeFree32-4     	100000000	        20.4 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkUnsafeFree64-4     	50000000	        34.1 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkUnsafeCalloc16-4   	50000000	        23.2 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkUnsafeCalloc32-4   	50000000	        28.0 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkUnsafeCalloc64-4   	50000000	        34.1 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkUnsafeMalloc16-4   	100000000	        13.8 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkUnsafeMalloc32-4   	100000000	        14.2 ns/op	       0 B/op	       0 allocs/op
//  BenchmarkUnsafeMalloc64-4   	100000000	        14.0 ns/op	       0 B/op	       0 allocs/op
//  PASS
//  ok  	modernc.org/memory	229.054s
package memory // import "modernc.org/memory"

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"

	"modernc.org/mathutil"
)

const mallocAllign = 16 // Must be >= 16

var (
	headerSize  = roundup(int(unsafe.Sizeof(page{})), mallocAllign)
	maxSlotSize = pageAvail >> 1
	osPageMask  = osPageSize - 1
	osPageSize  = os.Getpagesize()
	pageAvail   = pageSize - headerSize
	pageMask    = pageSize - 1
)

// if n%m != 0 { n += m-n%m }. m must be a power of 2.
func roundup(n, m int) int { return (n + m - 1) &^ (m - 1) }

type node struct {
	prev, next *node
}

type page struct {
	brk  int
	log  uint
	size int
	used int
}

// Allocator allocates and frees memory. Its zero value is ready for use.
type Allocator struct {
	allocs int // # of allocs.
	bytes  int // Asked from OS.
	cap    [64]int
	lists  [64]*node
	mmaps  int // Asked from OS.
	pages  [64]*page
	regs   map[*page]struct{}
}

func (a *Allocator) mmap(size int) (*page, error) {
	p, size, err := mmap(size)
	if err != nil {
		return nil, err
	}

	a.mmaps++
	a.bytes += size
	pg := (*page)(unsafe.Pointer(p))
	if a.regs == nil {
		a.regs = map[*page]struct{}{}
	}
	pg.size = size
	a.regs[pg] = struct{}{}
	return pg, nil
}

func (a *Allocator) newPage(size int) (*page, error) {
	size += headerSize
	p, err := a.mmap(size)
	if err != nil {
		return nil, err
	}

	p.log = 0
	return p, nil
}

func (a *Allocator) newSharedPage(log uint) (*page, error) {
	if a.cap[log] == 0 {
		a.cap[log] = pageAvail / (1 << log)
	}
	size := headerSize + a.cap[log]<<log
	p, err := a.mmap(size)
	if err != nil {
		return nil, err
	}

	a.pages[log] = p
	p.log = log
	return p, nil
}

func (a *Allocator) unmap(p *page) error {
	delete(a.regs, p)
	a.mmaps--
	return unmap(uintptr(unsafe.Pointer(p)), p.size)
}

// UintptrCalloc is like Calloc except it returns an uintptr.
func (a *Allocator) UintptrCalloc(size int) (r uintptr, err error) {
	if trace {
		defer func() {
			fmt.Fprintf(os.Stderr, "Calloc(%#x) %#x, %v\n", size, r, err)
		}()
	}
	if r, err = a.UintptrMalloc(size); r == 0 || err != nil {
		return 0, err
	}
	b := ((*rawmem)(unsafe.Pointer(r)))[:size]
	for i := range b {
		b[i] = 0
	}
	return r, nil
}

// UintptrFree is like Free except its argument is an uintptr, which must have
// been acquired from UintptrCalloc or UintptrMalloc or UintptrRealloc.
func (a *Allocator) UintptrFree(p uintptr) (err error) {
	if trace {
		defer func() {
			fmt.Fprintf(os.Stderr, "Free(%#x) %v\n", p, err)
		}()
	}
	if p == 0 {
		return nil
	}

	a.allocs--
	pg := (*page)(unsafe.Pointer(p &^ uintptr(pageMask)))
	log := pg.log
	if log == 0 {
		a.bytes -= pg.size
		return a.unmap(pg)
	}

	n := (*node)(unsafe.Pointer(p))
	n.prev = nil
	n.next = a.lists[log]
	if n.next != nil {
		n.next.prev = n
	}
	a.lists[log] = n
	pg.used--
	if pg.used != 0 {
		return nil
	}

	for i := 0; i < pg.brk; i++ {
		n := (*node)(unsafe.Pointer(uintptr(unsafe.Pointer(pg)) + uintptr(headerSize+i<<log)))
		switch {
		case n.prev == nil:
			a.lists[log] = n.next
			if n.next != nil {
				n.next.prev = nil
			}
		case n.next == nil:
			n.prev.next = nil
		default:
			n.prev.next = n.next
			n.next.prev = n.prev
		}
	}

	if a.pages[log] == pg {
		a.pages[log] = nil
	}
	a.bytes -= pg.size
	return a.unmap(pg)
}

// UintptrMalloc is like Malloc except it returns an uinptr.
func (a *Allocator) UintptrMalloc(size int) (r uintptr, err error) {
	if trace {
		defer func() {
			fmt.Fprintf(os.Stderr, "Malloc(%#x) %#x, %v\n", size, r, err)
		}()
	}
	if size < 0 {
		panic("invalid malloc size")
	}

	if size == 0 {
		return 0, nil
	}

	a.allocs++
	log := uint(mathutil.BitLen(roundup(size, mallocAllign) - 1))
	if 1<<log > maxSlotSize {
		p, err := a.newPage(size)
		if err != nil {
			return 0, err
		}

		return uintptr(unsafe.Pointer(p)) + uintptr(headerSize), nil
	}

	if a.lists[log] == nil && a.pages[log] == nil {
		if _, err := a.newSharedPage(log); err != nil {
			return 0, err
		}
	}

	if p := a.pages[log]; p != nil {
		p.used++
		p.brk++
		if p.brk == a.cap[log] {
			a.pages[log] = nil
		}
		return uintptr(unsafe.Pointer(p)) + uintptr(headerSize+(p.brk-1)<<log), nil
	}

	n := a.lists[log]
	p := (*page)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) &^ uintptr(pageMask)))
	a.lists[log] = n.next
	if n.next != nil {
		n.next.prev = nil
	}
	p.used++
	return uintptr(unsafe.Pointer(n)), nil
}

// UintptrRealloc is like Realloc except its first argument is an uintptr,
// which must have been returned from UintptrCalloc, UintptrMalloc or
// UintptrRealloc.
func (a *Allocator) UintptrRealloc(p uintptr, size int) (r uintptr, err error) {
	if trace {
		defer func() {
			fmt.Fprintf(os.Stderr, "UnsafeRealloc(%#x, %#x) %#x, %v\n", p, size, r, err)
		}()
	}
	switch {
	case p == 0:
		return a.UintptrMalloc(size)
	case size == 0 && p != 0:
		return 0, a.UintptrFree(p)
	}

	us := UintptrUsableSize(p)
	if us > size {
		return p, nil
	}

	if r, err = a.UintptrMalloc(size); err != nil {
		return 0, err
	}

	if us < size {
		size = us
	}
	copy((*rawmem)(unsafe.Pointer(r))[:size], (*rawmem)(unsafe.Pointer(p))[:size])
	return r, a.UintptrFree(p)
}

// UintptrUsableSize is like UsableSize except its argument is an uintptr,
// which must have been returned from UintptrCalloc, UintptrMalloc or
// UintptrRealloc.
func UintptrUsableSize(p uintptr) (r int) {
	if trace {
		defer func() {
			fmt.Fprintf(os.Stderr, "UsableSize(%#x) %#x\n", p, r)
		}()
	}
	if p == 0 {
		return 0
	}

	return usableSize(p)
}

func usableSize(p uintptr) (r int) {
	pg := (*page)(unsafe.Pointer(p &^ uintptr(pageMask)))
	if pg.log != 0 {
		return 1 << pg.log
	}

	return pg.size - headerSize
}

// Calloc is like Malloc except the allocated memory is zeroed.
func (a *Allocator) Calloc(size int) (r []byte, err error) {
	p, err := a.UintptrCalloc(size)
	if err != nil {
		return nil, err
	}

	var b []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh.Cap = usableSize(p)
	sh.Data = p
	sh.Len = size
	return b, nil
}

// Close releases all OS resources used by a and sets it to its zero value.
//
// It's not necessary to Close the Allocator when exiting a process.
func (a *Allocator) Close() (err error) {
	for p := range a.regs {
		if e := a.unmap(p); e != nil && err == nil {
			err = e
		}
	}
	*a = Allocator{}
	return err
}

// Free deallocates memory (as in C.free). The argument of Free must have been
// acquired from Calloc or Malloc or Realloc.
func (a *Allocator) Free(b []byte) (err error) {
	if b = b[:cap(b)]; len(b) == 0 {
		return nil
	}

	return a.UintptrFree(uintptr(unsafe.Pointer(&b[0])))
}

// Malloc allocates size bytes and returns a byte slice of the allocated
// memory. The memory is not initialized. Malloc panics for size < 0 and
// returns (nil, nil) for zero size.
//
// It's ok to reslice the returned slice but the result of appending to it
// cannot be passed to Free or Realloc as it may refer to a different backing
// array afterwards.
func (a *Allocator) Malloc(size int) (r []byte, err error) {
	p, err := a.UintptrMalloc(size)
	if p == 0 || err != nil {
		return nil, err
	}

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&r))
	sh.Cap = usableSize(p)
	sh.Data = p
	sh.Len = size
	return r, nil
}

// Realloc changes the size of the backing array of b to size bytes or returns
// an error, if any.  The contents will be unchanged in the range from the
// start of the region up to the minimum of the old and new  sizes.   If the
// new size is larger than the old size, the added memory will not be
// initialized.  If b's backing array is of zero size, then the call is
// equivalent to Malloc(size), for all values of size; if size is equal to
// zero, and b's backing array is not of zero size, then the call is equivalent
// to Free(b).  Unless b's backing array is of zero size, it must have been
// returned by an earlier call to Malloc, Calloc or Realloc.  If the area
// pointed to was moved, a Free(b) is done.
func (a *Allocator) Realloc(b []byte, size int) (r []byte, err error) {
	var p uintptr
	if b = b[:cap(b)]; len(b) != 0 {
		p = uintptr(unsafe.Pointer(&b[0]))
	}
	if p, err = a.UintptrRealloc(p, size); p == 0 || err != nil {
		return nil, err
	}

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&r))
	sh.Cap = usableSize(p)
	sh.Data = p
	sh.Len = size
	return r, nil
}

// UsableSize reports the size of the memory block allocated at p, which must
// point to the first byte of a slice returned from Calloc, Malloc or Realloc.
// The allocated memory block size can be larger than the size originally
// requested from Calloc, Malloc or Realloc.
func UsableSize(p *byte) (r int) { return UintptrUsableSize(uintptr(unsafe.Pointer(p))) }

// UnsafeCalloc is like Calloc except it returns an unsafe.Pointer.
func (a *Allocator) UnsafeCalloc(size int) (r unsafe.Pointer, err error) {
	p, err := a.UintptrCalloc(size)
	if err != nil {
		return nil, err
	}

	return unsafe.Pointer(p), nil
}

// UnsafeFree is like Free except its argument is an unsafe.Pointer, which must
// have been acquired from UnsafeCalloc or UnsafeMalloc or UnsafeRealloc.
func (a *Allocator) UnsafeFree(p unsafe.Pointer) (err error) { return a.UintptrFree(uintptr(p)) }

// UnsafeMalloc is like Malloc except it returns an unsafe.Pointer.
func (a *Allocator) UnsafeMalloc(size int) (r unsafe.Pointer, err error) {
	p, err := a.UintptrMalloc(size)
	if err != nil {
		return nil, err
	}

	return unsafe.Pointer(p), nil
}

// UnsafeRealloc is like Realloc except its first argument is an
// unsafe.Pointer, which must have been returned from UnsafeCalloc,
// UnsafeMalloc or UnsafeRealloc.
func (a *Allocator) UnsafeRealloc(p unsafe.Pointer, size int) (r unsafe.Pointer, err error) {
	q, err := a.UintptrRealloc(uintptr(p), size)
	if err != nil {
		return nil, err
	}

	return unsafe.Pointer(q), nil
}

// UnsafeUsableSize is like UsableSize except its argument is an
// unsafe.Pointer, which must have been returned from UnsafeCalloc,
// UnsafeMalloc or UnsafeRealloc.
func UnsafeUsableSize(p unsafe.Pointer) (r int) { return UintptrUsableSize(uintptr(p)) }
