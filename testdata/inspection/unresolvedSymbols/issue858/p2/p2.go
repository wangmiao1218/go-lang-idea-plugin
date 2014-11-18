package p2

import (
	"p1"
)

func good() {
	//p1.AA is a func defined in p1
	a:=p1.AA()
	_=a

	//p1.A is int ,defined in p1
	p1.A = 1

	//p1.V1 type is *p3.T1, defined in p3
	p1.V1.F1()

	//p1.V2 type is p3.T2, defined in p3
	p1.V2.F2()

	//p1.V3 type is p1.T3, F3() defined in file p1.go same with type T3.
	p1.V3.F3()

	//p1.V4 type is p1.T4, defined in file pf2.go not same with type T4.
	p1.V4.F4()
}

func bad(){
	a:=p1./*begin*/notExist1/*end.Unresolved symbol: 'notExist1'*/()
	_=a

	a:=p1.V1./*begin*/notExist2/*end.Unresolved symbol: 'notExist2'*/()
	_=a

	a:=p1.V2./*begin*/notExist3/*end.Unresolved symbol: 'notExist3'*/()
	_=a

	a:=p1.V3./*begin*/notExist4/*end.Unresolved symbol: 'notExist4'*/()
	_=a

	a:=p1.V4./*begin*/notExist5/*end.Unresolved symbol: 'notExist5'*/()
	_=a
}