package main

import "fmt"

func main() {
	var mm map[string]string
	fmt.Println("un-init map", mm)
	//panic: assignment to entry in nil map
	//mm["test"] = "ok"

	//full init
	var mm1 map[string]string = map[string]string{}
	fmt.Println("init map", mm1)
	//or
	mm2 := map[string]string{}
	mm2["test"] = "ok"
	fmt.Println(mm2)

	//short init
	var mm3 = make(map[string]string)
	mm3["firstName"] = "Raimo"
	fmt.Println("mm3", mm3)

	//getting value
	firstName := mm3["firstName"]
	fmt.Println("firstName", firstName, len(firstName))

	//if key no exist - default value returned
	lastName := mm3["lastName"]
	fmt.Println("lastName", lastName, len(lastName))

	//exist check
	lastName, ok := mm3["lastName"]
	fmt.Println("lastName is", lastName, "exist:", ok)

	//only exist bool
	_, exist := mm3["firstName"]
	fmt.Println("first Name exist:", exist)

	//deleting value
	delete(mm3, "firstName")
	_, exist = mm3["firstName"]
	fmt.Println("firstName exist:", exist)

	//copy by value
	key := "myKey"
	mm4 := mm3
	mm4[key] = "testOne"
	fmt.Println("mm3:", mm3, "mm4:", mm4)
}
