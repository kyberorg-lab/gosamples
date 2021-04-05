package main

import "gitlab.com/rvasily/msu-go-11/1/60_package_func_doc/world"

func main() {
	world.PrintStartRoom()

	//var from package
	println("starting room: ", world.StartingRoom)

	//same by function
	println("starting room: ", world.GetStartingLevel())

}
