package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
)

type Server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type XMLServers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Servers     []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}

func main() {
	content, err := ioutil.ReadFile("test.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	//decoder := xml.NewDecoder(bytes.NewBuffer(content))

	v := XMLServers{}
	err = xml.Unmarshal(content, &v)
	if err != nil {
		fmt.Println(err)
		return
	}

	w := bufio.NewWriter(os.Stdout)

	fmt.Fprintln(w, v)
	w.Flush()
}
