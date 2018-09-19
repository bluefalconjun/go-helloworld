package basictypes

/*
const declare
var declare
fast var declare

NOTE: 
when use const/var declare $id = $value use "="
when use fast var declare $id := $value use ":="

*/


//const type with any integer
const limit = 512

//const type uint16
const top uinit16 = 1421

//var with judge type=int
start := -19

//var with define type=int64
end := int64(9876543210)

//var type=int value=0
var i int

//var with judge type=bool
var debug = false

//var with judge type=bool
checkresult := true

//var with judge type=float64
setpSize := 1.5

//var with judge type=string
acronym := "FOSS"

const (
	Cyan = 0
	Magenta = 1
	Yellow = 2
)

//iota will +1 as in continues define. !
//and reset to 0 as each const keyword appears. !
const (
	Cyan1 = iota //0
	Magenta1 	 //1
	Yellow1		//2
)

//another use of itoa
type BitFlag int

const (
	Active BitFlag = 1<<itoa  //1<<0=1
	Send					  //1<<1=2
	Receive					  //1<<2=4
)

flag := Active | Send