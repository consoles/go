package main

import (
	"flag"
	"fmt"
)

func style1() {
	// `./flagsample --method hello` or `./flagsample -method hello`
	// hello -1
	methodPtr := flag.String("method", "default", "method of sample") // 返回值是字符串指针
	valuePtr := flag.Int("value", -1, "value of sample")
	flag.Parse()
	fmt.Println(*methodPtr, *valuePtr)
}

func style2() {
	var method string
	var value int

	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")
	flag.Parse()
	fmt.Println(method, value)
}

func main() {
	//	style1()
	style2()
}
