package main

import (
	"fmt"
	piscine ".."
)

func main() {
	piscine := piscine.Split("8JvVwU93tLkrUDIDVFvcIIB24BQuGDIDj90biLjB03MUeDID0J2yFPJ2AQGnxDIDW1ESC9b6UTkvwDIDuSCNbksPp8XqWDIDJcUJ6MuflChIuDIDKbkjA6a6LzgkcDIDK9O8pI0h9UN6bDIDFfxkzD7PcvWtTDIDkukH5Xeiv2132DID2VpaJfg8AUKAzDIDwGar4ZfFryKJE", "DID")
	fmt.Println(piscine)
	fmt.Println(piscine[12], len(piscine[12]))
	fmt.Println(piscine[11], len(piscine[11]))
}
