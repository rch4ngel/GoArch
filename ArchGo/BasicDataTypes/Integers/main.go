package main

import "fmt"

// The Integer Type Has Three Different Variations (int, uint, uintptr). int and uint type each have Four Different Sizes
// Defined Below:
var a int8
var b int16
var c int32 // Synonym For rune 'Word Size'
var d int64

// Unsigned int's will be used for all non-negative values. If assigned a negative value it will result in an overflow error.
var ua uint8 // Synonym For byte
var ub uint16
var uc uint32
var ud uint64

var ip uintptr // Used Only For Low Level Programming. No Specified Width But Large Enough To Hold Bit Pattern Of A Pointer Value.

func main() {
	a = 4
	b = -5
	c = 6
	d = 7
	ua = 8
	ub = 9
	uc = 10
	ud = 11

	ip = 34
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d\n", a, b, c, d, ua, ub, uc, ud, ip)
}
