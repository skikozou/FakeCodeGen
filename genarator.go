package main

import (
	"math"
	"math/rand"
)

func srcGen(raw [][]string) *src {
	source := NewSrc()
	//package and import
	if len(raw[0]) == 2 {
		source.pack.title = raw[0][0]
		source.pack.name = raw[0][1]
	} else if len(raw[0]) == 3 {
		source.pack.title = raw[0][0]
		source.pack.name = raw[0][1] + "-" + raw[0][2]
	} else if len(raw[0]) >= 4 {
		source.pack.title = raw[0][0]
		source.pack.name = raw[0][1]
		source.imp.title = raw[0][2]
		for i := 3; i < len(raw[0]); i++ {
			source.imp.name = append(source.imp.name, raw[0][i])
		}
	}
	//main
	if len(raw[1]) == 2 {
		source.main.title = raw[1][0]
		source.main.name = raw[1][1]
	} else {
		source.main.title = raw[1][0]
		source.main.name = raw[1][1]
		elen := len(raw[1]) - 2
		if elen == 2 {
			vObj := variable{Name: raw[1][2]}
			source.main.obj = append(source.main.obj, vObj)
		} else if elen == 1 {
			fObj := runfunc{name: []string{raw[1][2]}}
			source.main.obj = append(source.main.obj, fObj)
		} else if elen == 0 {
			vObj := variable{Name: "//	(´-ω-)"}
			source.main.obj = append(source.main.obj, vObj)
		} else {
			var fObj runfunc
			retun := int(math.Floor(float64(len(raw[1])-2) / 3))
			namen := int(math.Floor(float64(len(raw[1])-2-retun) / 2))
			for i := 2; i < retun+2; i++ {
				fObj.retu = append(fObj.retu, raw[1][i])
			}
			for i := retun + 2; i < retun+namen+2; i++ {
				fObj.name = append(fObj.name, raw[1][i])
			}
			for i := retun + namen + 2; i < len(raw[1]); i++ {
				fObj.arg = append(fObj.arg, raw[1][i])
			}
			source.main.obj = append(source.main.obj, fObj)
		}
	}
	//main obj
	nraw := raw[2:]
	for i, l := range nraw {
		if i > 5 {
			if len(l) >= 5 {
				randn := rand.Intn(2)
				if randn == 0 {
					fObj := function{}
					fObj.title = l[0]
					fObj.name = l[1]
					randn = rand.Intn(2)
					if randn == 0 {
						fObj.args = append(fObj.args, variable{Name: l[2], Key: l[3]})
						fObj.obj = append(fObj.obj, rfGen(l[4:]))
					} else {
						fObj.obj = append(fObj.obj, rfGen(l[2:]))
					}
					source.stru = append(source.stru, fObj)
				} else {
					sObj := structure{}
					cObj := variable{}
					sObj.typename = l[0]
					sObj.name = l[1]
					sObj.title = l[2]
					l = l[3:]
					for j := 1; len(l) > j; j = j + 2 {
						if j >= len(l) {
							cObj.Name = "//" + l[j-1]
						}
						sObj.Obj = append(sObj.Obj, variable{Name: l[j-1], Key: l[j]})
					}
					source.stru = append(source.stru, sObj)
					if cObj.Name != "" {
						source.stru = append(source.stru, cObj)
					}
				}
			} else {
				source.stru = append(source.stru, rfGen(l))
			}
			continue
		}

		if len(l) == 1 {
			source.main.obj = append(source.main.obj, runfunc{name: []string{l[0]}})
		} else if len(l) <= 5 {
			rObj := runfunc{}
			for i, w := range l {
				if i == 0 {
					rObj.name = append(rObj.name, w)
				} else if i == 1 {
					rObj.arg = append(rObj.arg, w)
				} else if i == 2 {
					rObj.retu = append(rObj.retu, w)
				} else {
					randn := rand.Intn(3)
					if randn == 0 {
						rObj.name = append(rObj.name, w)
					} else if randn == 1 {
						rObj.arg = append(rObj.arg, w)
					} else {
						rObj.retu = append(rObj.retu, w)
					}
				}
			}
			source.main.obj = append(source.main.obj, rObj)
		} else {
			randn := rand.Intn(3)
			if randn == 0 {
				if len(l) >= 4 {
					iObj := ifObj{}
					iObj.title = l[0]
					iObj.conditions1 = l[1]
					iObj.conditions2 = l[2]
					iObj.Obj = append(iObj.Obj, rfGen(l[3:]))
					source.main.obj = append(source.main.obj, iObj)
				}
			} else if randn == 1 {
				sObj := structObj{}
				cObj := variable{}
				sObj.name = l[0]
				sObj.title = l[1]
				l = l[2:]
				for j := 1; len(l) > j; j = j + 2 {
					if j >= len(l) {
						cObj.Name = "//" + l[j-1]
					}
					sObj.Obj = append(sObj.Obj, variable{Name: l[j], Key: l[j-1]})
				}
				source.main.obj = append(source.main.obj, sObj)
				if cObj.Name != "" {
					source.main.obj = append(source.main.obj, cObj)
				}
			} else {
				if len(l) >= 7 {
					fObj := forObj{}
					fObj.title = l[0]
					fObj.vari.Key = l[1]
					fObj.vari.Name = l[2]
					fObj.conditions.Key = l[3]
					fObj.conditions.Name = l[4]
					fObj.runobj = l[5]
					fObj.Obj = append(fObj.Obj, rfGen(l[6:]))
					source.main.obj = append(source.main.obj, fObj)
				} else {
					rObj := rfGen(l)
					source.main.obj = append(source.main.obj, rObj)
				}
			}
		}
	}

	return source
}

func rfGen(line []string) runfunc {
	rObj := runfunc{}
	rObj.name = append(rObj.name, line[0])
	for _, w := range line[1:] {
		randn := rand.Intn(3)
		if randn == 0 {
			rObj.name = append(rObj.name, w)
		} else if randn == 1 {
			rObj.retu = append(rObj.retu, w)
		} else {
			rObj.arg = append(rObj.arg, w)
		}
	}
	for j, _ := range rObj.retu {
		rObj.retu[j] = line[0]
		line = line[1:]
	}
	for j, _ := range rObj.name {
		rObj.name[j] = line[0]
		line = line[1:]
	}
	for j, _ := range rObj.arg {
		rObj.arg[j] = line[0]
		line = line[1:]
	}
	return rObj
}
