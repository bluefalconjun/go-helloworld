package basictype

import "fmt"

/*
key word in go:
the keyword is not allowed to used as identifier !

break default  select
case defer go
func interface map struct chan package
else goto switch return
const fallthrough if range type
continue for import
var

*/

/*
predefined identifier
DO NOT redefine predefined identifier !

append copy delete error print cap prtinln close iota len recover imag make new
int8 int16 int32 int64 ...
bool nil true false ...

*/

//sample for empty identifier as ignore result
func emptyidensample() {

	var x int
	count, err := fmt.Println(x)
	count, _ = fmt.Println(x)
	_, err = fmt.Println(x)
	fmt.Println(x)
	fmt.Println(count, err)
}
