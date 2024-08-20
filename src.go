package main

//src

type src struct {
	pack Pack
	imp  Imp
	main Main
	stru []interface{}
}

type Pack struct {
	title string
	name  string
}

type Imp struct {
	title string
	name  []string
}

type Main struct {
	title string
	name  string
	obj   []interface{}
}

//struct

//in

type runfunc struct {
	retu []string
	name []string
	arg  []string
}

type variable struct {
	Name string
	Key  string
}

type ifObj struct {
	title       string
	conditions1 string
	conditions2 string
	Obj         []runfunc
}

type forObj struct {
	title      string
	vari       variable
	conditions variable
	runobj     string
	Obj        []runfunc
}

type structObj struct {
	name  string
	title string
	Obj   []variable
}

//out

type function struct {
	title string
	name  string
	args  []variable
	obj   []interface{}
}

type structure struct {
	typename string
	name     string
	title    string
	Obj      []variable
}

func NewSrc() *src {
	return &src{}
}
