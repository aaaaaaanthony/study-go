package main

import (
	"testing"
)

func TestDemo(t *testing.T) {

	test1 := Test1()
	if test1 == true {
		t.Fail()
	}
}
