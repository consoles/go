package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 函数名 参数 返回值
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func sampleFromStdin() {
	fmt.Println("please input from stdin:")
	data, _ := ReadFrom(os.Stdin, 11) // 从stdin读取11Bytes
	fmt.Println(data)
}

func sampleFromString() {
	data, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(data)
}

func sampleFromFile() {
	file, _ := os.Open("main.go")
	defer file.Close()
	data, _ := ReadFrom(file, 9)
	fmt.Println(string(data))
}

func sampleBufIO() {
	strReader := strings.NewReader("hello world")
	bufReader := bufio.NewReader(strReader)
	data, _ := bufReader.Peek(5) // read 5 bytes
	fmt.Println(data, string(data), "已缓冲的字节：", bufReader.Buffered())

	str, _ := bufReader.ReadString(' ')
	fmt.Println(str, bufReader.Buffered())

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello ")
	fmt.Fprint(w, "world!")
	w.Flush()
}

func main() {
	//	sampleFromString()
	//	sampleFromStdin()
	//	sampleFromFile()
	sampleBufIO()
}
