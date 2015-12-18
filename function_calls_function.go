package main

var function_calls_function_a string

func function_calls_function() {
	function_calls_function_a = "G"
	print(function_calls_function_a)
	f1()
}

func f1() {
	function_calls_function_a := "0"
	print(function_calls_function_a)
	f2()
}

func f2() {
	print(function_calls_function_a)
}
