//test main
package main

import (
	"fmt"

	"github.com/bluefalconjun/gostudy/study/basictypes"
	"github.com/bluefalconjun/gostudy/study/defertracing"
	"github.com/bluefalconjun/gostudy/study/forgrammer"
	"github.com/bluefalconjun/gostudy/study/makingslices"

	"github.com/bluefalconjun/gostudy/study/arraytype"
	"github.com/bluefalconjun/gostudy/study/channel"
	"github.com/bluefalconjun/gostudy/study/function"
	"github.com/bluefalconjun/gostudy/study/interfaces"
	"github.com/bluefalconjun/gostudy/study/maps"
	"github.com/bluefalconjun/gostudy/study/runtime"
	"github.com/bluefalconjun/gostudy/study/slice"
)

var USESAMPLE = false

func main() {
	fmt.Printf("Hello,world\n")

	basictypes.Basictypes()
	defertracing.Defertracing()
	forgrammer.Forgrammer()
	makingslices.Makingslices()

	arraytype.ArrayType()
	slice.SlicesAppend()
	slice.SlicesRange()
	fmt.Printf("Go,world\n")
	maps.Maps()
	maps.MapsTest()
	functions.Funclosures()
	interfaces.Interfaces()
	interfaces.Checkelement()
	interfaces.StringerDemo()
	runtime.RuntimeDemo()
	channel.Channeldemo()
}
