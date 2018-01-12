package main

import (
	"fmt"

	"github.com/gredinger/crypto/crypto"
)

// Just making sure we can encode things correctly. SAILTHESHIPS <-> VDLOWKHVKLSV
func main() {
	fmt.Println(crypto.CeasarShiftRight("HELLOEVERYONE"))
	fmt.Println(crypto.CeasarShiftLeft("VDLOWKHVKLSV"))

}
