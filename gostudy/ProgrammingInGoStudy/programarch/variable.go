package programarch

import (
	"fmt"
)

//normal grammer
//var $name $type = $value
//one of [$type] and [= $value] can be ignored.
var varName int = 1



//Variable is
func Variable(){

	//short grammer
	//$name := $value
	//this can only be used in func. see Variable()
	s := "Variable !"
	fmt.Println(s)

	//create var with new func
	//new(T) will create a zero noname variable with type T and 
	//return the varialbe address.
	p := new(int) //p, *int address to int 0.
	
	fmt.Println(p)	//addr
	fmt.Println(*p) //val
}