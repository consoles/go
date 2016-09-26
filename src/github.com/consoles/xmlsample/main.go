package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string `xml:"name,attr"`
	Age  int    `xml:"age"`
}

func main() {
	p := person{Name: "console", Age: 18}

	// 结构体序列化为xml
	// 注意data和err的作用域
	//	if data, err := xml.Marshal(p); err != nil {
	//		fmt.Println(err)
	//		return
	//	} else {
	//		fmt.Println(string(data))
	//	}

	var data []byte
	var err error
	if data, err = xml.MarshalIndent(p, "", "  "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	p2 := new(person)
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2) // 输出是结构体
}
