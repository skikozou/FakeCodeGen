package main

import (
	"math/rand"
	"strings"
)

func (s *src) Build() string {
	var code string
	rn := "\r\n"
	tab := "	"
	space := " "
	level := 0
	code += s.pack.title + space + s.pack.name + rn + rn
	code += s.imp.title + space + "(" + rn
	for _, p := range s.imp.name {
		code += tab + "\"" + p + "\"" + rn
	}
	code += ")" + rn + rn
	code += s.main.title + space + s.main.name + "() {" + rn
	level++
	for _, o := range s.main.obj {
		code = ObjBuild(code, level, o)
	}
	level--
	code += "}\n\n"
	for _, s := range s.stru {
		code = StrBuild(code, level, s)
	}

	return code
}

func ObjBuild(code string, level int, o interface{}) string {
	rn := "\r\n"
	tab := "	"
	space := " "
	quo := "\""
	switch obj := o.(type) {
	case runfunc:
		code += strings.Repeat(tab, level)
		for i, w := range obj.retu {
			code += w
			if i != len(obj.retu)-1 {
				code += ", "
			} else {
				code += " := "
			}
		}
		for i, w := range obj.name {
			code += w
			if i != len(obj.name)-1 {
				code += "."
			} else {
				code += "("
			}
		}
		for i, w := range obj.arg {
			code += w
			if i != len(obj.arg)-1 {
				randn := rand.Intn(2)
				if randn == 1 {
					code += "."
				} else {
					code += ", "
				}
			}
		}
		code += ")"
		code += rn
	case variable:
		if obj.Key == "" {
			code += strings.Repeat(tab, level) + obj.Name + rn
		}
		code += strings.Repeat(tab, level) + obj.Name + " := " + obj.Key + rn
	case ifObj:
		code += strings.Repeat(tab, level) + obj.title + space + obj.conditions1 + space
		randn := rand.Intn(6)
		word := []string{"==", "!=", ">=", "<=", ">", "<"}
		code += word[randn]
		code += space + obj.conditions2 + " {\n"
		level++
		for _, word := range obj.Obj {
			code += strings.Repeat(tab, level)
			for i, w := range word.retu {
				code += w
				if i != len(word.retu)-1 {
					code += ", "
				} else {
					code += " := "
				}
			}
			for i, w := range word.name {
				code += w
				if i != len(word.name)-1 {
					code += "."
				} else {
					code += "("
				}
			}
			for i, w := range word.arg {
				code += w
				if i != len(word.arg)-1 {
					randn := rand.Intn(2)
					if randn == 1 {
						code += "."
					} else {
						code += ", "
					}
				}
			}
			code += ")"
			code += rn
		}
		level--
		code += strings.Repeat(tab, level) + "}\n"
	case forObj:
		code += strings.Repeat(tab, level) + obj.title + space + obj.vari.Name + " := " + obj.vari.Key + "; " + obj.conditions.Name + " "
		randn := rand.Intn(6)
		word := []string{"==", "!=", ">=", "<=", ">", "<"}
		code += word[randn]
		code += " " + obj.conditions.Key + "; " + obj.runobj
		randn = rand.Intn(2)
		word = []string{"--", "++"}
		code += word[randn]
		code += "{\n"
		level++
		for _, word := range obj.Obj {
			code += strings.Repeat(tab, level)
			for i, w := range word.retu {
				code += w
				if i != len(word.retu)-1 {
					code += ", "
				} else {
					code += " := "
				}
			}
			for i, w := range word.name {
				code += w
				if i != len(word.name)-1 {
					code += "."
				} else {
					code += "("
				}
			}
			for i, w := range word.arg {
				code += w
				if i != len(word.arg)-1 {
					randn := rand.Intn(2)
					if randn == 1 {
						code += "."
					} else {
						code += ", "
					}
				}
			}
			code += ")"
			code += rn
		}
		level--
		code += strings.Repeat(tab, level) + "}\n"
	case structObj:
		code += strings.Repeat(tab, level) + obj.name + " := " + obj.title + "{\n"
		level++
		for _, w := range obj.Obj {
			code += strings.Repeat(tab, level) + w.Name + ": " + quo + w.Key + quo + ",\n"
		}
		level--
		code += strings.Repeat(tab, level) + "}\n"
	default:
	}

	return code
}

func StrBuild(code string, level int, s interface{}) string {
	rn := "\r\n"
	tab := "	"
	space := " "
	switch stru := s.(type) {
	case function:
		code += strings.Repeat(tab, level) + stru.title + space + stru.name + "("
		for i, w := range stru.args {
			code += stru.name + space + w.Key
			if i != len(stru.args)-1 {
				code += ", "
			}
		}
		code += ")" + space + "{" + rn
		level++
		for _, l := range stru.obj {
			code = ObjBuild(code, level, l)
		}
		level--
		code += strings.Repeat(tab, level) + "}" + rn + rn
	case structure:
		code += strings.Repeat(tab, level) + stru.typename + space + stru.name + space + stru.title + space + "{" + rn
		level++
		for _, v := range stru.Obj {
			code += strings.Repeat(tab, level) + v.Name + tab + v.Key + rn
		}
		level--
		code += "}" + rn + rn
	default:
	}
	return code
}
