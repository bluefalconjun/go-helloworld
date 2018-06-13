//test main
package main

import (
	"fmt"

	"github.com/bluefalconjun/lang-study/gostudy/arraytype"
	"github.com/bluefalconjun/lang-study/gostudy/basictypes"
	"github.com/bluefalconjun/lang-study/gostudy/channel"
	"github.com/bluefalconjun/lang-study/gostudy/defertracing"
	"github.com/bluefalconjun/lang-study/gostudy/forgrammer"
	"github.com/bluefalconjun/lang-study/gostudy/interfaces"
	"github.com/bluefalconjun/lang-study/gostudy/makingslices"
	"github.com/bluefalconjun/lang-study/gostudy/maps"
	"github.com/bluefalconjun/lang-study/gostudy/runtime"
	"github.com/bluefalconjun/lang-study/gostudy/slice"
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
	//functions.Funclosures()
	interfaces.Interfaces()
	interfaces.Checkelement()
	interfaces.StringerDemo()
	runtime.RuntimeDemo()
	channel.Channeldemo()
}
