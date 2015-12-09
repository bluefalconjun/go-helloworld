package factory_func

//package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	where1 := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	log.SetFlags(log.Llongfile)
	log.Print("")

	where1()
	log.Print("")
	fmt.Println(addBmp("file1"))
	where1()
	//where2
	fmt.Println(addJpeg("file2"))
	where1()
	return
}
