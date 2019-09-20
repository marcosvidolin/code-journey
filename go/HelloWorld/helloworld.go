package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
 reader := bufio.NewReader(os.Stdin)
 fmt.Println("Hello, World.")
 // The '_' represents a write-only value to be used as a place-holder
 // where a variable is needed but the actual value is irrelevant.
 // https://golang.org/doc/effective_go.html#blank
 text, _ := reader.ReadString('\t')
 fmt.Println(text)
}
