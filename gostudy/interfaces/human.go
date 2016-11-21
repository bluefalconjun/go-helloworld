package interfaces

import "fmt"

//if one sub strut has method .
//ths super struct can call it.
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //sub struct
	school string
	loan   float32
}

type Employee struct {
	Human   //sub struct
	company string
	money   float32
}

//Human implement SayHi method
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human implement Sing method
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee overload Human 's SayHi method
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

//interfaces is the abstrct methods combination.
//it can not implement anything by itself.
//but it can be implement by other non-interfaces types.

// Interface Men can be implement by Human,Student and Employee.
// because all these three types has the required two method
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func SayHiSample() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//define Men
	var i Men

	//i can be Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i can be Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//define a Men slice
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//paul/sam/mike are three kind of type，but they do have the sam Men interface
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
	VoidInterfaceSample()

}

//empty interface can store any type. because it has no method.
type voidinterface interface {
}

var i int
var j voidinterface

func VoidInterfaceSample() {
	j = i
	fmt.Println(j)
}
