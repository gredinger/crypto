package crypto

import (
	"strings"
)

//Alphabet is the upper case 26 letters
var Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Ceasar does a 3 character shift with a rot 25 (capitals)
// Deprecated: Original credit to some Roman dude a couple millennium ago.
//
// ABCDEFGHIJKLMNOPQRSTUVWXYZ
// vvvvvvvvvvvvvvvvvvvvvvvvvv
// XYZABCDEFGHIJKLMNOPQRSTUVW
// ---
// Success! You've found buried treasure. Always keep an eye out for stuff like this.
// It could lead to interesting stories ;)
// And who knows what it might bring
// DWWDFNWKHJDWHVDWVXQGRZQ
// You'll make a great person of yourself .
func CeasarShiftLeft(input string) string {
	sr := ""
	for _, r := range strings.ToUpper(input) {
		ix := strings.IndexRune(Alphabet, r) // Index of char
		if ix < 3 {
			ix = ix + 26
		}
		tv := ix - 3
		sr += string(Alphabet[tv])
	}
	return sr
}

//CeasarShiftRight does a 3 character shift to the right with a rot 25 (captials)
// Deprecated: Original credit to some Roman dude a couple millennium ago.
//
// XYZABCDEFGHIJKLMNOPQRSTUVW
// VVVVVVVVVVVVVVVVVVVVVVVVVV
// ABCDEFGHIJKLMNOPQRSTUVWXYZ
//-----
// How'd you get here? Hit me up.
// Literally the author's name @ gmail.com. ;)
// <3
//
func CeasarShiftRight(input string) string {
	sr := ""
	for _, r := range strings.ToUpper(input) {
		sr += string(Alphabet[(strings.IndexRune(Alphabet, r)+3)%26])
	}
	return sr
}
