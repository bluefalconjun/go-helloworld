package main

var glocal_scopea = "G"

func glocal_scope() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	glocal_scopea = "0"
	print(a)
}
