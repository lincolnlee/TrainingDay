package classtest

import (
	"fmt"
)

type AClass struct {
}

func deal() {
	aClass := &AClass{}
	aClass.test2()
}

func (aClass1 *AClass) test1() {
	fmt.Println("I am test1()")
}

func (aClass2 *AClass) test2() {
	fmt.Println("I am test2()")
	fmt.Println("Calling test1()")
	aClass2.test1()
	fmt.Println("test2() done")
}
