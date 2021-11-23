package main

import (
	"fmt"
	"github.com/gtools/date"
	"github.com/gtools/str"
)

func main() {
	fmt.Println("hello world", date.Now())
	fmt.Println(str.Contains("abcdefg", "bcd"))
}
