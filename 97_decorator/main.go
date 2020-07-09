package main

import "fmt"

type iStuff interface {
	DoStuff()
}

type stuff struct {
	iStuff
	Name string
}

type realStuff string

func (r realStuff) DoStuff() {
	fmt.Println(r)
}

type fakeStuff int

func (f fakeStuff) DoStuff() {
	fmt.Println("It's fake")
}

func main() {
	r := realStuff("Hei")
	f := fakeStuff(0)

	rS := stuff{
		iStuff: r,
		Name:   "stuff",
	}

	fS := stuff{
		iStuff: f,
		Name:   "fake",
	}

	rS.DoStuff()
	fS.DoStuff()
}
