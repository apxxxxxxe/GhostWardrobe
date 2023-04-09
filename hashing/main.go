package main

import (
	"fmt"
	"os"
	"hash/crc32"
)

func main() {
  str := os.Args[1]
	table := crc32.MakeTable(crc32.IEEE)
	hash := crc32.Checksum([]byte(str), table)
	fmt.Printf("%x", hash)
}
