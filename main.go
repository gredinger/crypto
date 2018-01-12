package main

import (
	"fmt"

	"github.com/gredinger/crypto/crypto"
)

func main() {
	fmt.Println(crypto.CeasarShiftRight("SAILTHESHIPS"))
	fmt.Println(crypto.CeasarShiftLeft("VDLOWKHVKLSV"))

	fmt.Println(crypto.CeasarShiftLeft("DWWDFNWKHJDWHVDWVXQGRZQ"))
}
