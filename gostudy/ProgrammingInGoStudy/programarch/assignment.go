package programarch

import (
	"fmt"
)

var x int = 1
var p bool = false


//Assignment is
func Assignment(){

	x = 1 //assign by var name.
	*(&p) = true //assign by var addr.

	//group assignment
	y := 2
	x,y = y,x

	//implicit assignment
	medals := []string{"gold", "silver", "bronze"}

	fmt.Println(medals)
}