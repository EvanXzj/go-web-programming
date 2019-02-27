// steps:
// 1. create a struct with data
// 2. create a xml file to store data
// 3. create a encoder with xml file
// 4. encodes struct into file
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	// step 1
	post := Post{
		Id: "1",
		Content: "Hello World!",
		Author: Author{
			Id: "2",
			Name: "Sau Sheong",
		},
	}

	// step 2
	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}

	// step 3
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")

	// step 4
	err = encoder.Encode(&post) // xml.Header ?
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}
