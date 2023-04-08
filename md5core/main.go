package main

import (
  "crypto/md5"
  "fmt"
  "os"
)

func main() {
    fmt.Printf("%x", md5.Sum([]byte(os.Args[1])))
}
