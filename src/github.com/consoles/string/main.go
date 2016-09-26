package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "hello world"
	fmt.Println(strings.Contains(s, "hell"), strings.Contains(s, "??"))
	fmt.Println(strconv.ParseFloat("3.14", 32)) // 3.140000104904175
	fmt.Println(strconv.ParseFloat("3.14", 64)) // 3.14
	fmt.Println(strconv.FormatInt(123, 2))
}
