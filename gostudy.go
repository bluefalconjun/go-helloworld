//test main
package main

import (
	"fmt"

	"github.com/bluefalconjun/gostudy/basictypes"
	"github.com/bluefalconjun/gostudy/defertracing"
	"github.com/bluefalconjun/gostudy/forgrammer"
	"github.com/bluefalconjun/gostudy/makingslices"

	"github.com/bluefalconjun/gostudy/arraytype"
	"github.com/bluefalconjun/gostudy/function"
	"github.com/bluefalconjun/gostudy/maps"
	"github.com/bluefalconjun/gostudy/slice"
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
}
