//test main
package main

import (
	"fmt"

	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/arraytype"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/basictype"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/channel"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/defertracing"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/forgrammer"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/function"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/interfaces"
	//"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/makingslices"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/maps"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/namedresult"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/runtime"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/slice"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/string"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/structpointer"
	"github.com/bluefalconjun/lang-study/gostudy/ProgrammingInGoStudy/switchsample"
)

var USESAMPLE = false

func main() {
	fmt.Printf("Hello,world\n")

	basictype.Basictypes()
	defertracing.Defertracing()
	forgrammer.Forgrammer()
	//makingslices.Makingslices()
	function.Funclosures()
	string.StringTest()

	arraytype.ArrayType()
	slice.SlicesAppend()
	slice.SlicesRange()
	fmt.Printf("Go,world\n")
	maps.Maps()
	namedresult.Namedresult()
	structpointer.Structpointer()

	//functions.Funclosures()
	interfaces.Interfaces()
	interfaces.Checkelement()
	interfaces.StringerDemo()
	runtime.RuntimeDemo()
	channel.Channeldemo()
	switchsample.Switchsample()
}
